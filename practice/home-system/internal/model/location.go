package model

import (
	"github.com/jinzhu/gorm"
)

type Location struct {
	*Model
	//库位名称
	Name string `json:"name"`
}

func (model Location) TableName() string {
	return "location"
}

func (t Location) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name=?", t.Name)
	}

	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Location) List(db *gorm.DB, pageOffset, pageSize int) ([]*Location, error) {
	var tags []*Location
	var err error

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)

	}
	if t.Name != "" {
		db.Where("name = ?", t.Name)
	}

	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil

}

func (t Location) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Location) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ? and is_del=?", t.Id, 0).Updates(values).Error
}

func (t Location) Delete(db *gorm.DB) error {
	return db.Where("id=? and is_del = ?", t.Id, 0).Delete(&t).Error
}
