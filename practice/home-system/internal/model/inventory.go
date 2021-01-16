package model

import (
	"github.com/jinzhu/gorm"
)

type Inventory struct {
	*Model
	//数量，在某日date的数量
	Amount float64 `json:"amount"`
	//日期
	Date uint32 `json:"date"`
	//物品ID
	ItemId uint32 `json:"item_id"`
	//库位id
	LocationId uint32 `json:"location_id"`
}

func (model Inventory) TableName() string {
	return "inventory"
}

func (t Inventory) Count(db *gorm.DB) (int, error) {
	var count int

	if t.ItemId > 0 {
		db = db.Where("item_id=?", t.ItemId)
	}
	if t.LocationId > 0 {
		db = db.Where("location_id=?", t.LocationId)
	}
	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Inventory) List(db *gorm.DB, pageOffset, pageSize int) ([]*Inventory, error) {
	var tags []*Inventory
	var err error

	if t.ItemId > 0 {
		db = db.Where("item_id=?", t.ItemId)
	}
	if t.LocationId > 0 {
		db = db.Where("location_id=?", t.LocationId)
	}

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)

	}

	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil

}

func (t Inventory) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Inventory) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ? and is_del=?", t.Id, 0).Updates(values).Error
}

func (t Inventory) Delete(db *gorm.DB) error {
	return db.Where("id=? and is_del = ?", t.Id, 0).Delete(&t).Error
}
