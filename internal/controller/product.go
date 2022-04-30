package controller

import (
	"context"

	"helloworld/apiv1"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Product = cProduct{}
)

type cProduct struct{}

func (h *cProduct) Product(ctx context.Context, req *apiv1.ProductReq) (res *apiv1.ProductRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}
