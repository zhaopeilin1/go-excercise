package model

import (
	"github.com/jinzhu/gorm"
)

type InventoryEvent struct {
	*Model
	//库存事件，增减库存，盘点等。
	//物品ID
	ItemId uint32 `json:"item_id"`

	//库位id
	LocationId uint32 `json:"location_id"`
	//数量变化
	AmountChange float64 `json:"amount_change"`
	//数量变化
	Amount float64 `json:"amount"`

	//日期
	Date uint32 `json:"date"`
	//备注
	Remark string `json:"remark"`
}

func (model InventoryEvent) TableName() string {
	return "inventory_event"
}

func (t InventoryEvent) Count(db *gorm.DB) (int, error) {
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

func (t InventoryEvent) List(db *gorm.DB, pageOffset, pageSize int) ([]*InventoryEvent, error) {
	var tags []*InventoryEvent
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

func (t InventoryEvent) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t InventoryEvent) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ? and is_del=?", t.Id, 0).Updates(values).Error
}

func (t InventoryEvent) Delete(db *gorm.DB) error {
	return db.Where("id=? and is_del = ?", t.Id, 0).Delete(&t).Error
}
