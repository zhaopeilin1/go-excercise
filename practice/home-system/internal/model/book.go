package model

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	*Model
	//标题
	Title string `json:title`
	//副标题
	Subtitle string `json:subtitle`
	//作者，多个作者用逗号分隔
	Author string `json:author`
	//字数
	Length int `json:length`
	//分类
	Category string `json:category`
	//描述
	Description string `json:description`
	//出版社
	Publisher string `json:publisher`
	//出版时间
	PublishTime string `json:publish_time`
	//isbn
	ISBN string `json:isbn`

	//CategoryId uint32 `json:`
}

func (model Book) TableName() string {
	return "book"
}

func (t Book) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Title != "" {
		db = db.Where("title like '%?%'", t.Title)
	}

	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Book) List(db *gorm.DB, pageOffset, pageSize int) ([]*Book, error) {
	var tags []*Book
	var err error

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Title != "" {
		db.Where("title like '%?%'", t.Title)
	}

	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil

}

func (t *Book) Create(db *gorm.DB) error {
	err := db.Create(t).Error
	return err
}

func (t Book) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ? and is_del=?", t.Id, 0).Updates(values).Error
}

func (t Book) Delete(db *gorm.DB) error {
	return db.Where("id=? and is_del = ?", t.Id, 0).Delete(&t).Error
}
