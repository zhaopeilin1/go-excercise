package model

import "github.com/jinzhu/gorm"

type Category struct {
	*Model
	// 类目名称
	Name string `json:"name"`
	//类目描述
	Describe string `json:"describe"`
}

func (model Category) TableName() string {
	return "category"
}

func (t Category) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Name != "" {
		db = db.Where("name like '%?%'", t.Name)
	}

	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Category) List(db *gorm.DB, pageOffset, pageSize int) ([]*Category, error) {
	var Categorys []*Category
	var err error

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)

	}
	if t.Name != "" {
		db = db.Where("name like '%?%'", t.Name)
	}

	if err = db.Where("is_del = ?", 0).Find(&Categorys).Error; err != nil {
		return nil, err
	}
	return Categorys, nil

}

func (t Category) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Category) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ? and is_del=?", t.Id, 0).Updates(values).Error
}

func (t Category) Delete(db *gorm.DB) error {
	return db.Where("id=? and is_del = ?", t.Id, 0).Delete(&t).Error
}
