package dao

import (
	"library/internal/model"
	"library/pkg/app"
	//	"library/pkg/app"
)

type Article struct {
	ID            uint32 `json:"id"`
	TagID         uint32 `json:"tag_id"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

// func (d *Dao) CountArticle(title, desc, coverImageUrl, content string, state uint8) (int, error) {
// 	tag := model.Article{Title: title, State: state}
// 	return tag.Count(d.engine)
// }

// func (d *Dao) GetArticleList(name string, state uint8, page, pageSize int) ([]*model.Article, error) {
// 	tag := model.Article{Title: name, State: state}
// 	pageOffset := app.GetPageOffset(page, pageSize)
// 	return tag.List(d.engine, pageOffset, pageSize)
// }

func (d *Dao) CreateArticle(param *Article) (*model.Article, error) {
	article := model.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		CoverImageUrl: param.CoverImageUrl,
		Content:       param.Content,
		State:         param.State,
		Model:         &model.Model{CreatedBy: param.CreatedBy},
	}

	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(param *Article) error {
	article := model.Article{Model: &model.Model{Id: param.ID}}
	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
	}
	if param.Title != "" {
		values["title"] = param.Title
	}

	if param.CoverImageUrl != "" {
		values["cover_image_url"] = param.CoverImageUrl
	}
	if param.Desc != "" {
		values["desc"] = param.Desc
	}
	if param.Content != "" {
		values["content"] = param.Content
	}

	return article.Update(d.engine, values)
}

func (d *Dao) GetArticle(id uint32, state uint8) (model.Article, error) {
	article := model.Article{Model: &model.Model{Id: id}, State: state}
	return article.Get(d.engine)
}

func (d *Dao) CountArticleListByTagID(id uint32, state uint8) (int, error) {
	article := model.Article{State: state}
	return article.CountByTagID(d.engine, id)
}

func (d *Dao) GetArticleListByTagID(id uint32, state uint8, page, pageSize int) ([]*model.ArticleRow, error) {
	article := model.Article{State: state}
	return article.ListByTagID(d.engine, id, app.GetPageOffset(page, pageSize), pageSize)
}

func set_not_empty(m map[string]interface{}, name, value string) {
	if value != "" {
		m[name] = value
	}
}

// func (d *Dao) UpdateArticle(id uint32, title, desc, coverImageUrl, content string, state uint8, modifiedBy string) error {
// 	tag := model.Article{
// 		Model: &model.Model{
// 			Id: id,
// 		},
// 	}
// 	values := map[string]interface{}{
// 		"state":       state,
// 		"modified_by": modifiedBy,
// 	}
// 	set_not_empty(values, "title", title)
// 	set_not_empty(values, "desc", desc)
// 	set_not_empty(values, "cover_image_url", coverImageUrl)
// 	set_not_empty(values, "content", content)

// 	return tag.Update(d.engine, values)
// }

func (d *Dao) DeleteArticle(id uint32) error {

	tag := model.Article{Model: &model.Model{Id: id}}
	return tag.Delete(d.engine)
}
