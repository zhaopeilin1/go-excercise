package dao

import (
	"home-system/internal/model"
	"home-system/pkg/app"
)

func (d *Dao) CountPosition(name string, state uint8) (int, error) {
	tag := model.Position{Name: name, State: state}
	return tag.Count(d.engine)
}

func (d *Dao) GetPositionList(name string, state uint8, page, pageSize int) ([]*model.Position, error) {
	tag := model.Position{Name: name, State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return tag.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreatePosition(name string, state uint8, createBy string) error {
	tag := model.Position{
		Name:  name,
		State: state,
		Model: &model.Model{CreateBy: createBy},
	}

	return tag.Create(d.engine)
}

func (d *Dao) UpdatePosition(id uint32, name string, state uint8, modifiedBy string) error {
	tag := model.Position{
		Model: &model.Model{
			Id: id,
		},
	}
	values := map[string]interface{}{
		"state":       state,
		"modified_by": modifiedBy,
	}
	if name != "" {
		values["name"] = name
	}

	return tag.Update(d.engine, values)
}

func (d *Dao) DeletePosition(id uint32) error {
	tag := model.Position{Model: &model.Model{Id: id}}
	return tag.Delete(d.engine)
}
