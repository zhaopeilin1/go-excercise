package dao

import (
	"library/internal/model"
	"library/pkg/app"
)

type Chapter struct {
	ID uint32 `json:id`

	BookId uint32 `json:book_id`
	//标题
	Title   string `json:title`
	Content string `json:content`

	//字数
	Length int `json:length`
	Offset int `json:offset`
}

func (d *Dao) CountChapter(bookId uint32) (int, error) {
	chapter := model.Chapter{BookId: bookId}
	return chapter.Count(d.engine)
}

func (d *Dao) GetChapterList(bookId uint32, page, pageSize int) ([]*model.Chapter, error) {
	chapter := model.Chapter{BookId: bookId}
	pageOffset := app.GetPageOffset(page, pageSize)
	return chapter.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateChapter(param *Chapter) (uint32, error) {
	chapter := &model.Chapter{
		BookId:  param.BookId,
		Title:   param.Title,
		Content: param.Content,
		Length:  param.Length,
		Offset:  param.Offset,
	}
	err := chapter.Create(d.engine)
	if err != nil {
		return 0, err
	} else {
		return chapter.Id, nil
	}
}

// func (d *Dao) UpdateChapter(param *Chapter) error {
// 	chapter := model.Chapter{
// 		Model: &model.Model{
// 			Id: param.ID,
// 		},
// 	}
// 	values := map[string]interface{}{}

// 	return chapter.Update(d.engine, values)
// }

func (d *Dao) DeleteChapter(id uint32) error {
	chapter := model.Chapter{Model: &model.Model{Id: id}}
	return chapter.Delete(d.engine)
}
