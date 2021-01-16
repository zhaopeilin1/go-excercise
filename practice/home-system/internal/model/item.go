package model

import (
	"github.com/jinzhu/gorm"
)

type Item struct {
	*Model
	// 物品名称
	Name string `json:"name"`

	// 状态
	State    string `json:"state"`
	Category string `json:"category"`
	//类目
	CategoryId uint32 `json:"category_id"`

	HomeId uint32 `json:"home_id"`
	//计量单位
	Unit string `json:"unit"`
	//消耗速度
	ConsumptionSpeed float64 `json:"consumption_speed"`
}

func (model Item) TableName() string {
	return "item"
}

func (t Item) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name like '%?%'", t.Name)
	}
	if t.HomeId != 0 {
		db = db.Where("home_id=?", t.HomeId)
	}
	if t.CategoryId != 0 {
		db = db.Where("category_id=?", t.CategoryId)
	}

	if t.State != "" {
		db = db.Where("state = ?", t.State)
	}

	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Item) List(db *gorm.DB, pageOffset, pageSize int) ([]*Item, error) {
	var tags []*Item
	var err error

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)

	}
	if t.Name != "" {
		db = db.Where("name like '%?%'", t.Name)
	}
	if t.HomeId != 0 {
		db = db.Where("home_id=?", t.HomeId)
	}
	if t.CategoryId != 0 {
		db = db.Where("category_id=?", t.CategoryId)
	}

	if t.State != "" {
		db = db.Where("state = ?", t.State)
	}

	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil

}

func (t Item) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Item) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ? and is_del=?", t.Id, 0).Updates(values).Error
}

func (t Item) Delete(db *gorm.DB) error {
	return db.Where("id=? and is_del = ?", t.Id, 0).Delete(&t).Error
}
