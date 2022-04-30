package service

import (
	"context"
	"helloworld/internal/model"
	"helloworld/internal/model/entity"
	"helloworld/internal/service/internal/dao"
)

type (
	// sUser is service struct of module User.
	sUser struct{}
)

var (
	// insUser is the instance of service User.
	insUser = sUser{}
)

// User returns the interface of User service.
func User() *sUser {
	return &insUser
}
func (s *sUser) GetUser(ctx context.Context, in model.UserList) (user []*entity.User, err error) {
	err = dao.User.Ctx(ctx).Page(in.Page, in.Limit).Scan(&user)

	return user, err

}
