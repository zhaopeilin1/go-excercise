package service

import (
	"home-system/internal/model"
	"home-system/pkg/app"
)

type CountCategoryRequest struct {
	Name  string `form:"name" binding:"max=100"`
	//State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CategoryListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	//State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateCategoryRequest struct {
	Name     string `form:"name" binding:"required,min=3,max=100"`
	CreateBy string `form:"create_by" binding:"required,min=2,max=100"`
	//State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateCategoryRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	//State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteCategoryRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}
type Category struct {
	//BookName  string `form:"book_name" binding:"max=100"`
	Id int
	Name string
	Describe string

}

func (svc *Service) CountCategory(param *CountCategoryRequest) (int, error) {
	return 1,nil;
	return svc.dao.CountCategory(param.Name)
}

func (svc *Service) GetCategoryList(param *CategoryListRequest, pager *app.Pager) ([]*model.Category, error) {
	return svc.dao.GetCategoryList(param.Name, pager.Page, pager.PageSize)
}

func (svc *Service) CreateCategory(param *CreateCategoryRequest) error {
	return svc.dao.CreateCategory(param.Name, param.CreateBy)
}

func (svc *Service) UpdateCategory(param *UpdateCategoryRequest) error {
	return svc.dao.UpdateCategory(param.ID, param.Name,  param.ModifiedBy)
}

func (svc *Service) DeleteCategory(param *DeleteCategoryRequest) error {
	return svc.dao.DeleteCategory(param.ID)
}