package controller

import (
	"context"

	"helloworld/apiv1"
	"helloworld/internal/model"
	"helloworld/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	User = cUser{}
)

type cUser struct{}

func (h *cUser) GetUser(ctx context.Context, req *apiv1.GetUserReq) (res *apiv1.GetUserRes, err error) {
	// g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	userdata, _ := service.User().GetUser(ctx, model.UserList{req.PageNum, req.PageSize})
	err = g.RequestFromCtx(ctx).Response.WriteJson(g.Map{"data": userdata})
	if err != nil {
		return
	}
	return
}
