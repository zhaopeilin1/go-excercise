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

func (d *Dao) CountChapter(bookId uint32, keyword string) (int, error) {
	chapter := model.Chapter{BookId: bookId, Content: keyword}
	return chapter.Count(d.engine)
}

func (d *Dao) GetChapterList(bookId uint32, keyword string, page, pageSize int) ([]*model.Chapter, error) {
	chapter := model.Chapter{BookId: bookId, Content: keyword}
	pageOffset := app.GetPageOffset(page, pageSize)
	return chapter.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) FindChapterByIds(ids []uint32) ([]model.Chapter, error) {
	chapter := model.Chapter{}
	return chapter.FindByIds(d.engine, ids)
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

func (d *Dao) GetChapter(id uint32) (model.Chapter, error) {
	chapter := model.Chapter{
		Model: &model.Model{
			Id: uint32(id),
		},
	}
	return chapter.Get(d.engine)
}

func (d *Dao) GetPreChapter(id int) (model.Chapter, error) {
	chapter := model.Chapter{
		Model: &model.Model{
			Id: uint32(id),
		},
	}
	return chapter.GetPre(d.engine)
}

func (d *Dao) GetNextChapter(id int) (model.Chapter, error) {
	chapter := model.Chapter{
		Model: &model.Model{
			Id: uint32(id),
		},
	}
	return chapter.GetNext(d.engine)
}

func (d *Dao) GetPreNChapter(bookId uint32,offset, pre int) ([]model.Chapter, error) {
	chapter := model.Chapter{
		BookId: bookId,
		Offset: offset,
	}
	return chapter.GetPreN(d.engine, pre)
}

func (d *Dao) GetNextNChapter(bookId uint32,offset, next int) ([]model.Chapter, error) {
	chapter := model.Chapter{
		BookId: bookId,
		Offset: offset,
	}
	return chapter.GetNextN(d.engine, next)
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
