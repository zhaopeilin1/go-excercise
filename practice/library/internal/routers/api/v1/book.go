package v1

import (
	"library/internal/model"
	"library/pkg/upload"
	"os"
	"path"

	"io/ioutil"
	"library/global"
	"library/internal/service"
	"library/pkg/app"

	"library/pkg/convert"
	"library/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Book struct{}

func NewBook() Book {
	return Book{}
}

func (t Book) UploadAndIndex(c *gin.Context) {
	_, fileHeader, err := c.Request.FormFile("file")
	fileName := fileHeader.Filename

	uniqueFileName := upload.GetUniqueFilePath(fileName)
	fileHeader.Filename = uniqueFileName
	upload1 := NewUpload()

	upload1.UploadFile(c)

	response := app.NewResponse(c)

	uploadSavePath := upload.GetSavePath()
	filePath := path.Join(uploadSavePath, uniqueFileName)

	f, err := os.Open(filePath)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorReadFileFail)
		return
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorReadFileFail)
		return
	}
	text := string(bytes)
	svc := service.New(c.Request.Context())

	ext := upload.GetFileExt(fileName)

	bookName := fileName[:len(fileName)-len(ext)]

	svc.IndexBook(text, bookName)

	response.ToResponse(gin.H{"index": uniqueFileName})
}

func (t Book) Search(c *gin.Context) {
	param := service.SearchBookRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	total, chapters := svc.Search(param.Keyword,&pager)
	response.ToResponseList(chapters,total)
}


//根据章节id,获取此章节的前n个，后m个章节。
func (t Book) GetChapters(c *gin.Context) {
	response := app.NewResponse(c)
	chapterId := convert.StrTo(c.Param("id")).MustUInt32()
	svc := service.New(c.Request.Context())
	chapter, err := svc.GetChapter(chapterId)
	var chapters []model.Chapter
	if err != nil {
		global.Logger.Errorf("GetChapter err:%v", err)
		response.ToErrorResponse(errcode.ServerError.WithDetails(err.Error()))
		return
	}
	preSize:=convert.StrTo(c.Query("pre")).MustInt()
	if preSize>0{
		pre,_ := svc.GetPreNChapter(chapter.BookId,chapter.Offset,preSize)
		chapters= append(chapters,pre...)
	}
	chapters= append(chapters,chapter)
	nextSize:=convert.StrTo(c.Query("next")).MustInt()
	if nextSize>0{
		next,_:=svc.GetNextNChapter(chapter.BookId,chapter.Offset,nextSize)
		chapters= append(chapters,next...)
	}



	response.ToResponseList(chapters,1+preSize+nextSize)
}

//根据书籍id分页查询章节
func (t Book) GetBookChapters(c *gin.Context) {
	response := app.NewResponse(c)
	bookId := convert.StrTo(c.Param("id")).MustUInt32()
	svc := service.New(c.Request.Context())
	c1 := &service.CountChapterRequest{BookId: bookId}
	total, err := svc.CountChapter(c1)
	if err != nil {
		global.Logger.Errorf("GetBookChapters err:%v", err)
		response.ToErrorResponse(errcode.ServerError.WithDetails(err.Error()))
		return
	}

	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	c2 := &service.ChapterListRequest{BookId: bookId}
	chapters, err := svc.GetBookChapterList(c2,&pager)
	if err != nil {
		global.Logger.Errorf("GetBookChapters err:%v", err)
		response.ToErrorResponse(errcode.ServerError.WithDetails(err.Error()))
		return
	}
	response.ToResponseList(chapters,total)
}

// @Summary 获取多个标签
// @Produce  json
// @Param name query string false "标签名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.Book "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [get]
func (t Book) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	request := &service.BookListRequest{Title: c.GetString("title")}
	totalRows, err := svc.CountBook(request)
	if err != nil {
		global.Logger.Errorf("svc.CountBook err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.GetBookList(request, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetBookList err :%v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
	return

}

// @Summary 新增标签
// @Produce  json
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string false "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Book "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags [post]
func (t Book) Create(c *gin.Context) {

	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)

	if err != nil {
		global.Logger.Errorf("svc.CreateTag err :%v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return

}

// @Summary 更新标签
// @Produce  json
// @Param id path int true "标签ID"
// @Param name body string false "标签名称" minlength(3) maxlength(100)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Book "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [put]
func (t Book) Update(c *gin.Context) {

	param := service.UpdateTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)

	if err != nil {
		global.Logger.Errorf("svc.UpdateTag err :%v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// @Summary 删除标签
// @Produce  json
// @Param id path int true "标签ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags/{id} [delete]
func (t Book) Delete(c *gin.Context) {

	param := service.DeleteTagRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)

	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)

	if err != nil {
		global.Logger.Errorf("svc.DeleteTag err :%v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
