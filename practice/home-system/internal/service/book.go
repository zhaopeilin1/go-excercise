package service

import (
	"fmt"
	"github.com/blevesearch/bleve"
	"home-system/internal/model"
	"home-system/pkg/app"
	"home-system/pkg/convert"

	"home-system/global"

	"home-system/internal/dao"
	"regexp"
	"strings"
)

//章节标题的正则表达式
var chapTitleRegEx = "^.{0,8}[第]? *[一二三四五六七八九十百千万两1234567890 ]+[章]"

//type CountBookRequest struct {
//	Title  string `form:"title"`
//	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
//}
type SearchBookRequest struct {
	//BookName  string `form:"book_name" binding:"max=100"`
	Keyword string  `form:"keyword"`
}

type BookListRequest struct {
	Title  string `form:"title"`
	//Name  string `form:"name" binding:"max=100"`
	//State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateBookRequest struct {
	Name     string `form:"name" binding:"required,min=3,max=100"`
	CreateBy string `form:"create_by" binding:"required,min=2,max=100"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateBookRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteBookRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountBook(param *BookListRequest) (int, error) {
	return svc.dao.CountBook(param.Title)
}
func splitBookContent(content string) []*dao.Chapter {
	lines := strings.Split(content, "\n")
	re := regexp.MustCompile(chapTitleRegEx)

	titleIndexs := []int{}
	chapters := []*dao.Chapter{&dao.Chapter{Title: "前言"}}

	for index, line := range lines {
		if re.MatchString(line) {
			titleIndexs = append(titleIndexs, index)
			chapters = append(chapters, &dao.Chapter{Title: line})
		}
	}

	start := 0
	for i, index := range titleIndexs {
		chapters[i].Content = strings.Join(lines[start:index], "\n")
		chapters[i].Length = len(chapters[i].Content)
		start = index
	}

	chapters[len(titleIndexs)].Content = strings.Join(lines[start:], "\n")
	chapters[len(titleIndexs)].Length = len(chapters[len(titleIndexs)].Content)
	return chapters

}

func (svc *Service) IndexBook(content, title string) error {
	// global.Logger.Info("title:", title, "content:", content[:240])
	//1 create book 2 split content, 3 save chapter and index
	book := &dao.Book{
		Title:  title,
		Length: len(content),
	}
	id, err := svc.dao.CreateBook(book)
	if err != nil {
		return err
	}

	chapters := splitBookContent(content)
	for i, chapter := range chapters {
		chapter.BookId = id
		chapter.Offset = i
		chapterId, err := svc.dao.CreateChapter(chapter)
		if err != nil {
			return err
		}
		chapter.ID = chapterId
	}
	//章节索引
	batch := (*global.Index).NewBatch()
	for _, chapter := range chapters {
		docId := model.Chapter{}.TableName()+ "_" + fmt.Sprint(chapter.ID)
		batch.Index(docId, chapter.Content)
		//(*global.Index).Index(docId, chapter.Content)
	}
	fmt.Println("batch_size:",batch.Size())
	(*global.Index).Batch(batch)
	return nil
}



func (svc *Service) Search(keyword string, pager *app.Pager) (int,[]model.Chapter) {
	//内容查询。
	fmt.Println("keyword,",keyword)
	query := bleve.NewMatchQuery(keyword)
	request := bleve.NewSearchRequest(query)
	request.Size = pager.PageSize;
	request.From = app.GetPageOffset(pager.Page,pager.PageSize)
	searchResult, err := (*global.Index).Search(request)
	if err != nil {
		global.Logger.Errorf("search error %v",err)
		return 0, make([]model.Chapter,0)
	}
	fmt.Println("searchresult:",searchResult)
	if searchResult.Hits.Len() == 0 {
		return 0, make([]model.Chapter,0)
	}
	fmt.Println("hists size:",len(searchResult.Hits),"len:",searchResult.Hits.Len())
	ids := make([]uint32,pager.PageSize)
	for i,doc := range  searchResult.Hits {
		docID := doc.ID
		split := strings.Split(docID, "_")
		//chapterID := strings.Replace(docID,model.Chapter{}.TableName()+ "_" ,"",1)
		chapterID:=split[len(split)-1]
		ids[i] =  uint32(convert.StrTo(chapterID).MustInt())
	}
	chapters, err := svc.dao.FindChapterByIds(ids)
	return int(searchResult.Total),chapters
}

// func (svc *Service) CountBook(param *CountBookRequest) (int, error) {
// 	return svc.dao.CountBook(param.Name, param.State)
// }

func (svc *Service) GetBookList(param *BookListRequest, pager *app.Pager) ([]*model.Book, error) {
	return svc.dao.GetBookList(param.Title, pager.Page, pager.PageSize)
}

// func (svc *Service) CreateBook(param *CreateBookRequest) error {
// 	return svc.dao.CreateBook(param.Name, param.State, param.CreateBy)
// }

// func (svc *Service) UpdateBook(param *UpdateBookRequest) error {
// 	return svc.dao.UpdateBook(param.ID, param.Name, param.State, param.ModifyBy)
// }

// func (svc *Service) DeleteBook(param *DeleteBookRequest) error {
// 	return svc.dao.DeleteBook(param.ID)
// }
