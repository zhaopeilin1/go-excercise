package model

type ArticleTag struct {
	*Model

	// 文章ID
	ArticleId int32 `json:article_id`
	// 标签ID
	TagId int32 `json:tag_id`
}

func (model ArticleTag) TableName() string {
	return "blog_article_tag"
}
