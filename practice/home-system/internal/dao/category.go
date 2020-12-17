package dao

import (
	"home-system/internal/model"
	"home-system/pkg/app"
)

func (d *Dao) CountCategory(name string) (int, error) {
	Category := model.Category{Name: name}
	return Category.Count(d.engine)
}

func (d *Dao) GetCategoryList(name string, page, pageSize int) ([]*model.Category, error) {
	Category := model.Category{Name: name}
	pageOffset := app.GetPageOffset(page, pageSize)
	return Category.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateCategory(name string, createBy string) error {
	Category := model.Category{
		Name:  name,
		Model: &model.Model{CreateBy: createBy},
	}

	return Category.Create(d.engine)
}

func (d *Dao) UpdateCategory(id uint32, name string, modifiedBy string) error {
	Category := model.Category{
		Model: &model.Model{
			Id: id,
		},
	}
	values := map[string]interface{}{
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	return Category.Update(d.engine, values)
}

func (d *Dao) DeleteCategory(id uint32) error {
	Category := model.Category{Model: &model.Model{Id: id}}
	return Category.Delete(d.engine)
}