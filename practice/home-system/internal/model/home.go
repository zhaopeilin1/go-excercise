package model

import (
	"github.com/jinzhu/gorm"
)

type Home struct {
	*Model
	// 家庭名称
	Name string `json:"name"`

	// 状态 0为禁用、1为启用
	//State uint8 `json:state`
}

func (model Home) TableName() string {
	return "home"
}

func (t Home) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db.Where("name like '%?%'", t.Name)
	}

	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Home) List(db *gorm.DB, pageOffset, pageSize int) ([]*Home, error) {
	var homes []*Home
	var err error

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)

	}
	if t.Name != "" {
		db.Where("name like '%?%'", t.Name)
	}

	if err = db.Where("is_del = ?", 0).Find(&homes).Error; err != nil {
		return nil, err
	}
	return homes, nil

}

func (t Home) Create(db *gorm.DB) (*Home, error) {
	err := db.Create(&t).Error
	return &t, err
}

func (t Home) Get(db *gorm.DB) (Home, error) {

	var home Home
	db = db.Where("id = ?", t.Id)
	err := db.First(&home).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return home, err
	}
	return home, nil

}

func (t Home) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ? and is_del=?", t.Id, 0).Updates(values).Error
}

func (t Home) Delete(db *gorm.DB) error {
	return db.Where("id=? and is_del = ?", t.Id, 0).Delete(&t).Error
}
