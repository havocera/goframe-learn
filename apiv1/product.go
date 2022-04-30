package apiv1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ProductReq struct {
	g.Meta `path:"/product" tags:"Hello" method:"get" summary:"这是一个测试接口，返回helloworld"`
}
type ProductRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
