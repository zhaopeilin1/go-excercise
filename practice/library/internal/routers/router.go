package routers

import (
	_ "library/docs"
	"library/internal/middleware"
	"library/internal/routers/api/v1"

	// "library/pkg/upload"

	"net/http"

	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		//fmt.Println(method)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, developerId")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(Cors())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.Translations())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//tag := v1.NewTag()
	book := v1.NewBook()
	upload := v1.NewUpload()
	//article := v1.NewArticle()

	apiv1 := r.Group("/api/v1")
	r.POST("/upload/file", upload.UploadFile)
	{
		//apiv1.POST("/tags", tag.Create)
		//apiv1.DELETE("/tags/:id", tag.Delete)
		//apiv1.PUT("/tags/:id", tag.Update)
		//apiv1.PATCH("/tags/:id/state", tag.Update)
		//apiv1.GET("/tags", tag.List)

		apiv1.POST("/books/upload", book.UploadAndIndex)
		apiv1.GET("/books/search", book.Search)
		//根据章节id，查询前后章节
		apiv1.GET("/chapters/:id", book.GetChapters)
		//根据书籍id分页查询章节
		apiv1.GET("/books/chapters/:id", book.GetBookChapters)

		apiv1.POST("/books", book.Create)
		apiv1.DELETE("/books/:id", book.Delete)
		apiv1.PUT("/books/:id", book.Update)
		//apiv1.PATCH("/books/:id/state", book.Update)
		apiv1.GET("/books", book.List)

		//apiv1.POST("/articles", article.Create)
		//apiv1.DELETE("/articles/:id", article.Delete)
		//apiv1.PUT("/articles/:id", article.Update)
		//apiv1.PATCH("/articles/:id/state", article.Update)
		//apiv1.GET("/articles/:id", article.Get)
		//apiv1.GET("/articles", article.List)
	}
	return r

}
