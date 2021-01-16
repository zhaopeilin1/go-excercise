package service

import (
	"home-system/internal/dao"
	"home-system/internal/model"
	"home-system/pkg/app"
)

type HomeRequest struct {
	ID    uint32 `json:"id" binding:"required,gte=1"`
	State uint8  `json:"state,default=1" binding:"oneof=0 1"`
}

type HomeListRequest struct {
	Name string `json:"name"`
}

type CreateHomeRequest struct {
	Name string `json:"name" binding:"required,min=2,max=255"`
}

type UpdateHomeRequest struct {
	ID   uint32 `json:"id" binding:"required,gte=1"`
	Name string `json:"Name" binding:"min=2,max=100"`
}

type DeleteHomeRequest struct {
	ID uint32 `json:"id" binding:"required,gte=1"`
}

func (svc *Service) CreateHome(param *CreateHomeRequest) (*model.Home, error) {
	home, err := svc.dao.CreateHome(&dao.Home{
		Name: param.Name,
	})
	if err != nil {
		return nil, err
	}

	return home, nil
}

func (svc *Service) UpdateHome(param *UpdateHomeRequest) error {
	err := svc.dao.UpdateHome(&dao.Home{
		ID:   param.ID,
		Name: param.Name,
	})
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) CountHome(param *HomeListRequest) (int, error) {
	return svc.dao.CountHome(param.Name)
}

func (svc *Service) GetHomeList(param *HomeListRequest, pager *app.Pager) ([]*model.Home, error) {
	return svc.dao.GetHomeList(param.Name, pager.Page, pager.PageSize)
}
