package service

import (
	"fmt"
	// "io/ioutil"
	"library/internal/dao"
	// "library/internal/model"
	"regexp"
	"strings"
)

//章节标题的正则表达式
var chapTitleRegEx = "^.{0,8}[第]? *[一二三四五六七八九十百千万两1234567890 ]+[章]"

type CountBookRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type BookListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
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

func splitBookContent(content string) []*dao.Chapter {
	lines := strings.Split(content, "\n")
	re := regexp.MustCompile(chapTitleRegEx)

	titleIndexs := []int{}
	chapters := []*dao.Chapter{&dao.Chapter{Title: "前言"}}

	for index, line := range lines {
		if re.MatchString(line) {
			fmt.Println(index, line[:])
			titleIndexs = append(titleIndexs, index)
			chapters = append(chapters, &dao.Chapter{Title: line})
		}
	}

	start := 0
	for i, index := range titleIndexs {
		chapters[i].Content = strings.Join(lines[start:index], "\n")
		chapters[i].Length = len(chapters[i].Content)
		fmt.Println("length:", chapters[i].Length)
		start = index
	}

	chapters[len(titleIndexs)].Content = strings.Join(lines[start:], "\n")
	chapters[len(titleIndexs)].Length = len(chapters[len(titleIndexs)].Content)
	return chapters

}

func (svc *Service) IndexBook(content, title string) error {
	// global.Logger.Info("title:", title, "content:", content[:240])
	// fmt.Println(content[:240])
	//1 create book 2 split content, 3 save chapter and index
	book := &dao.Book{
		Title:  title,
		Length: len(content),
	}
	id, err := svc.dao.CreateBook(book)
	if err != nil {
		return err
	}

	// fmt.Println("bookId:", id)

	chapters := splitBookContent(content)
	for _, chapter := range chapters {
		chapter.BookId = id
		_, err = svc.dao.CreateChapter(chapter)
		if err != nil {
			return err
		}

	}

	// fmt.Println()

	return nil
}

// func (svc *Service) CountBook(param *CountBookRequest) (int, error) {
// 	return svc.dao.CountBook(param.Name, param.State)
// }

// func (svc *Service) GetBookList(param *BookListRequest, pager *app.Pager) ([]*model.Book, error) {
// 	return svc.dao.GetBookList(param.Name, param.State, pager.Page, pager.PageSize)
// }

// func (svc *Service) CreateBook(param *CreateBookRequest) error {
// 	return svc.dao.CreateBook(param.Name, param.State, param.CreateBy)
// }

// func (svc *Service) UpdateBook(param *UpdateBookRequest) error {
// 	return svc.dao.UpdateBook(param.ID, param.Name, param.State, param.ModifiedBy)
// }

// func (svc *Service) DeleteBook(param *DeleteBookRequest) error {
// 	return svc.dao.DeleteBook(param.ID)
// }
