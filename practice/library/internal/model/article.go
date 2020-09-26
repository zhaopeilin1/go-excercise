package model

import (
	"library/pkg/app"

	"github.com/jinzhu/gorm"
)

type Article struct {
	*Model
	//书ID
	BookId string `json:book_id`
	//章节标题
	Title string `json:title`
	// 文章内容
	Content string `json:content`
	//在书中的位置，第几章
	Offset uint32 `json:offset`
	//字数
	Length uint32 `json:length`

	// 文章简述
	Desc string `json:desc`
	// 封面图片地址
	CoverImageUrl string `json:cover_image_url`
	// 状态 0为禁用、1为启用
	State uint8 `json:state`
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
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

type ArticleRow struct {
	ArticleID     uint32
	TagID         uint32
	TagName       string
	ArticleTitle  string
	ArticleDesc   string
	CoverImageUrl string
	Content       string
}

func (a Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {
	fields := []string{"ar.id AS article_id", "ar.title AS article_title", "ar.desc AS article_desc", "ar.cover_image_url", "ar.content"}
	fields = append(fields, []string{"t.id AS tag_id", "t.name AS tag_name"}...)

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	rows, err := db.Select(fields).Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id = ar.id").
		Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*ArticleRow
	for rows.Next() {
		r := &ArticleRow{}
		if err := rows.Scan(&r.ArticleID, &r.ArticleTitle, &r.ArticleDesc, &r.CoverImageUrl, &r.Content, &r.TagID, &r.TagName); err != nil {
			return nil, err
		}

		articles = append(articles, r)
	}

	return articles, nil
}

func (a Article) CountByTagID(db *gorm.DB, tagID uint32) (int, error) {
	var count int
	err := db.Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id = ar.id").
		Where("at.`tag_id` = ? AND ar.state = ? AND ar.is_del = ?", tagID, a.State, 0).
		Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}
