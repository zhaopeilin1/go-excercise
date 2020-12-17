package service

import "errors"

type AuthRequest struct {
	AppKey string `json:"app_key" form:"app_key"`
	AppSecret string `json:"app_secret" form:"app_secret"`
}
func (svc *Service) CheckAuth(param *AuthRequest) error {
	auth,err := svc.dao.GetAuth(
		param.AppKey,
		param.AppSecret,
		)
	if err!= nil {
		return err
	}
	if auth.Id>0{
		return nil
	}
	return errors.New("auth info does not exists.")
}