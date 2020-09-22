package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model

	// 文章标题
	Title string `json:title`
	// 文章简述
	Desc string `json:desc`
	// 封面图片地址
	CoverImageUrl string `json:cover_image_url`
	// 文章内容
	Content string `json:content`
	// 状态 0为禁用、1为启用
	State uint8 `json:state`
}

func (t Article) Count(db *gorm.DB) (int, error) {
	var count int
	if t.Title != "" {
		db = db.Where("title like ''", t.Title)
	}
	db = db.Where("state = ?", t.State)

	if err := db.Model(&t).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (model Article) TableName() string {
	return "blog_article"
}

func (t Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&t).Error; err != nil {
		return nil, err
	}
	return &t, nil
}

func (t Article) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ? and is_del=?", t.Id, 0).Update(values).Error

}

func (t Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? and is_del=?", t.Id, 0).Delete(&t).Error
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	db = db.Where("id = ? AND state = ? AND is_del = ?", a.Id, a.State, 0)
	err := db.First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}
	return article, nil
	// return db.Where("id = ? and is_del=?", t.Id, 0).Error
}
