package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetUserReq struct {
	g.Meta   `path:"/getuser" tags:"User" method:"get" summary:"获取用户表"`
	PageNum  int
	PageSize int
}
type GetUserRes struct {
	g.Meta `mime:"json" example:"string"`
}
