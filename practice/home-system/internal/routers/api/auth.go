package api

import (
	"github.com/gin-gonic/gin"
	"home-system/global"
	"home-system/internal/service"
	"home-system/pkg/app"
	"home-system/pkg/errcode"
)

func GetAuth(c *gin.Context){
	param:=service.AuthRequest{}
	response:=app.NewResponse(c);
	valid,errs:= app.BindAndValid(c,&param)
	global.Logger.Info("param:",param.AppSecret,param.AppKey)
	if valid != true {
		global.Logger.Errorf("app.BindAndValid errs: %v",errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}
	svc:=service.New(c)
	err:=svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("svc.CheckAuth err: %v",err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}
	token,err := app.GenerateToken(param.AppKey,param.AppSecret)
	if err!= nil {
		global.Logger.Errorf("app.GenerateToken err: %v",err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}
	response.ToResponse(gin.H{"token":token})
}
