package v1

//家庭，一个用户可以创建多个家庭。用户可以加入家庭，用户与家庭是多对多关系。

import (
	"home-system/global"
	"home-system/internal/service"
	"home-system/pkg/app"
	"home-system/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Home struct{}

func NewHome() Home {
	return Home{}
}

func (t Home) Get(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}
func (t Home) List(c *gin.Context) {

	param := service.HomeListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountHome(&service.HomeListRequest{Name: param.Name})
	if err != nil {
		global.Logger.Errorf("svc.CountHome err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountHomeFail)
		return
	}

	homes, err := svc.GetHomeList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetHomeList err :%v", err)
		response.ToErrorResponse(errcode.ErrorGetHomeListFail)
		return
	}

	response.ToResponseList(homes, totalRows)
	return

}
func (t Home) Create(c *gin.Context) {
	param := service.CreateHomeRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	_, err := svc.CreateHome(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateHome err:%v", err)
		response.ToErrorResponse(errcode.ErrorCreateHomeFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
func (t Home) Update(c *gin.Context) {
	param := service.UpdateHomeRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.UpdateHome(&param)

	if err != nil {
		global.Logger.Errorf("svc.UpdateHome err:%v", err)
		response.ToErrorResponse(errcode.ErrorCreateHomeFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
func (t Home) Delete(c *gin.Context) {}
