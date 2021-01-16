package dao

import (
	"home-system/internal/model"
	"home-system/pkg/app"
	//	"home-system/pkg/app"
)

type Home struct {
	ID            uint32 `json:"id"`
	Name          string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         uint8  `json:"state"`
}

func (d *Dao) CountHome(name string) (int, error) {
	home := model.Home{Name: name}
	return home.Count(d.engine)
}

func (d *Dao) GetHomeList(name string, page, pageSize int) ([]*model.Home, error) {
	home := model.Home{Name: name}
	pageOffset := app.GetPageOffset(page, pageSize)
	return home.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateHome(param *Home) (*model.Home, error) {
	home := model.Home{
		Name:  param.Name,
		Model: &model.Model{CreateBy: param.CreatedBy},
	}

	return home.Create(d.engine)
}

func (d *Dao) UpdateHome(param *Home) error {
	home := model.Home{Model: &model.Model{Id: param.ID}}
	values := map[string]interface{}{
		"modified_by": param.ModifiedBy,
		"state":       param.State,
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

	return home.Update(d.engine, values)
}

func (d *Dao) GetHome(id uint32, state uint8) (model.Home, error) {
	home := model.Home{Model: &model.Model{Id: id}}
	return home.Get(d.engine)
}

func (d *Dao) CountHomeListByTagID(id uint32, state uint8) (int, error) {
	// home := model.Home{State: state}
	// return home.CountByTagID(d.engine, id)
	return 0, nil
}

// func (d *Dao) GetHomeListByTagID(id uint32, state uint8, page, pageSize int) ([]*model.HomeRow, error) {
// 	home := model.Home{State: state}
// 	return home.ListByTagID(d.engine, id, app.GetPageOffset(page, pageSize), pageSize)
// }

// func (d *Dao) UpdateHome(id uint32, title, desc, coverImageUrl, content string, state uint8, modifiedBy string) error {
// 	tag := model.Home{
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

func (d *Dao) DeleteHome(id uint32) error {

	tag := model.Home{Model: &model.Model{Id: id}}
	return tag.Delete(d.engine)
}
