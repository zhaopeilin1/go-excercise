package model

import (
	"github.com/jinzhu/gorm"
)

type Chapter struct {
	*Model
	BookId  uint32 `json:book_id`
	Title   string `json:title`
	Content string `json:content`
	Offset  int    `json:offset`
	Length  int    `json:length`
}

func (model Chapter) TableName() string {
	return "chapter"
}

func (t Chapter) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Title != "" {
		db = db.Where("title like '?%'", t.Title)
	}

	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Chapter) List(db *gorm.DB, pageOffset, pageSize int) ([]*Chapter, error) {
	var tags []*Chapter
	var err error

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Title != "" {
		db.Where("title like '?%'", t.Title)
	}

	if err = db.Where("is_del = ?", 0).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil

}

func (a Chapter) Get(db *gorm.DB) (Chapter, error) {
	var chapter Chapter
	db = db.Where("id = ? AND  is_del = ?", a.Id, 0)
	err := db.First(&chapter).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return chapter, err
	}
	return chapter, nil

}

func (t *Chapter) Create(db *gorm.DB) error {

	return db.Create(t).Error
}

func (t Chapter) BatchInsert(db *gorm.DB) error {
	return nil

}

func (t Chapter) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ? and is_del=?", t.Id, 0).Updates(values).Error
}

func (t Chapter) Delete(db *gorm.DB) error {
	return db.Where("id=? and is_del = ?", t.Id, 0).Delete(&t).Error
}
