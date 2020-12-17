package service

import (
	//"fmt"
	"home-system/internal/model"
	"home-system/pkg/app"

	// "io/ioutil"
	"home-system/internal/dao"
	// "home-system/internal/model"
	"regexp"
	"strings"
)

//章节标题的正则表达式
//var chapTitleRegEx = "^.{0,8}[第]? *[一二三四五六七八九十百千万两1234567890 ]+[章]"

type CountChapterRequest struct {
	BookId  uint32  //`form:"book_id"`
	Keyword string  //`form:"keyword"`
}

type ChapterListRequest struct {
	BookId  uint32 //`form:"book_id"`
	Keyword string //`form:"name" binding:"max=100"`
}

type CreateChapterRequest struct {
	Name     string `form:"name" binding:"required,min=3,max=100"`
	CreateBy string `form:"create_by" binding:"required,min=2,max=100"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateChapterRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteChapterRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func splitChapterContent(content string) []*dao.Chapter {
	lines := strings.Split(content, "\n")
	re := regexp.MustCompile(chapTitleRegEx)

	titleIndexs := []int{}
	chapters := []*dao.Chapter{&dao.Chapter{Title: "前言"}}

	for index, line := range lines {
		if re.MatchString(line) {
			// fmt.Println(index, line[:])
			titleIndexs = append(titleIndexs, index)
			chapters = append(chapters, &dao.Chapter{Title: line})
		}
	}

	start := 0
	for i, index := range titleIndexs {
		chapters[i].Content = strings.Join(lines[start:index], "\n")
		chapters[i].Length = len(chapters[i].Content)
		// fmt.Println("length:", chapters[i].Length)
		start = index
	}

	chapters[len(titleIndexs)].Content = strings.Join(lines[start:], "\n")
	chapters[len(titleIndexs)].Length = len(chapters[len(titleIndexs)].Content)
	return chapters

}

func (svc *Service) CountChapter(param *CountChapterRequest) (int, error) {
	return svc.dao.CountChapter(uint32(param.BookId), param.Keyword)
}

func (svc *Service) GetBookChapterList(param *ChapterListRequest, pager *app.Pager) ([]*model.Chapter, error) {
	return svc.dao.GetChapterList(param.BookId, param.Keyword, pager.Page, pager.PageSize)
}

func (svc *Service) GetChapter(id uint32) (model.Chapter, error) {
	return svc.dao.GetChapter(id)
}

func (svc *Service) GetPreChapter(id int) (model.Chapter, error) {
	return svc.dao.GetPreChapter(id)
}

func (svc *Service) GetNextChapter(id int) (model.Chapter, error) {
	return svc.dao.GetNextChapter(id)
}

func (svc *Service) GetPreNChapter(bookId uint32,offset, pre int) ([]model.Chapter, error) {
	return svc.dao.GetPreNChapter(bookId,offset, pre)
}

func (svc *Service) GetNextNChapter(bookId uint32,offset, next int) ([]model.Chapter, error) {
	return svc.dao.GetNextNChapter(bookId,offset, next)
}
