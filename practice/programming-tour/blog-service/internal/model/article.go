package model

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
	State int8 `json:state`
}

func (model Article) TableName() string {
	return "blog_article"
}
