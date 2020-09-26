package service

// "library/internal/model"
// "library/pkg/app"
// "library/pkg/upload"
// "mime/multipart"
// "path"

type CountBookRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type BookListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateBookRequest struct {
	Name     string `form:"name" binding:"required,min=3,max=100"`
	CreateBy string `form:"create_by" binding:"required,min=2,max=100"`
	State    uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateBookRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"max=100"`
	State      uint8  `form:"state,default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteBookRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

// func (svc *Service) CountBook(param *CountBookRequest) (int, error) {
// 	return svc.dao.CountBook(param.Name, param.State)
// }

// func (svc *Service) GetBookList(param *BookListRequest, pager *app.Pager) ([]*model.Book, error) {
// 	return svc.dao.GetBookList(param.Name, param.State, pager.Page, pager.PageSize)
// }

// func (svc *Service) CreateBook(param *CreateBookRequest) error {
// 	return svc.dao.CreateBook(param.Name, param.State, param.CreateBy)
// }

// func (svc *Service) UpdateBook(param *UpdateBookRequest) error {
// 	return svc.dao.UpdateBook(param.ID, param.Name, param.State, param.ModifiedBy)
// }

// func (svc *Service) DeleteBook(param *DeleteBookRequest) error {
// 	return svc.dao.DeleteBook(param.ID)
// }
