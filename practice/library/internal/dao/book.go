package dao

import (
	"library/internal/model"
	"library/pkg/app"
)

type Book struct {
	ID uint32 `json:id`
	//标题
	Title string `json:title`
	//副标题
	Subtitle string `json:subtitle`
	//作者，多个作者用逗号分隔
	Author string `json:author`
	//字数
	Length int `json:length`
	//分类
	Category string `json:category`
	//描述
	Description string `json:description`
	//出版社
	Publisher string `json:publisher`
	//出版时间
	PublishTime string `json:publish_time`
	//isbn
	ISBN string `json:isbn`

	//CategoryId uint32 `json:`
}

func (d *Dao) CountBook(title string) (int, error) {
	book := model.Book{Title: title}
	return book.Count(d.engine)
}

func (d *Dao) GetBookList(title string, page, pageSize int) ([]*model.Book, error) {
	book := model.Book{Title: title}
	pageOffset := app.GetPageOffset(page, pageSize)
	return book.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateBook(param *Book) (uint32, error) {
	book := &model.Book{
		Title:       param.Title,
		Subtitle:    param.Subtitle,
		Author:      param.Author,
		Length:      param.Length,
		Category:    param.Category,
		Description: param.Description,
		Publisher:   param.Publisher,
		PublishTime: param.PublishTime,
		ISBN:        param.ISBN,
	}
	err := book.Create(d.engine)
	if err != nil {
		return 0, err
	} else {

		return book.Id, nil
	}
}

func (d *Dao) UpdateBook(param *Book) error {
	book := model.Book{
		Model: &model.Model{
			Id: param.ID,
		},
	}
	values := map[string]interface{}{}
	app.SetNotEmptyStr(values, "subtitle", param.Subtitle)
	app.SetNotEmptyStr(values, "author", param.Author)
	app.SetNotEmptyStr(values, "category", param.Category)
	app.SetNotEmptyStr(values, "description", param.Description)
	app.SetNotEmptyStr(values, "publisher", param.Publisher)
	app.SetNotEmptyStr(values, "publishTime", param.PublishTime)
	app.SetNotEmptyStr(values, "isbn", param.ISBN)

	return book.Update(d.engine, values)
}

func (d *Dao) DeleteBook(id uint32) error {
	book := model.Book{Model: &model.Model{Id: id}}
	return book.Delete(d.engine)
}
