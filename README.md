# goframe框架说明文档
## 根目录的sql文件为对应学习数据库
## goframe创建

```shell
gf init 项目名称
```

## goframe 安装包

```shell
go mod tidy	
```

## 运行项目

```shell
gf run ./main.go
```

## 项目关键目录指引

1. cmd.go  -------加载所有的controller
2. Apiv1 -----------定义接口的 路径、参数、接口分组、请求方式（GET, POST, PUT , DELETE）;定义返回值的类型：text/html，json
3. controller------定义接口返回函数，给前端返回数据
4. model ---------定义数据模型，既定义传入参数   in参数格式，又定义返回参数格式，
5. service --------定义数据库操作，
6. config.yaml---------项目数据库配置，端口配置

## 基础代码复制

### service

```go
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

```

### model

```go
package model

type UserList struct {
	Page  int
	Limit int
}

```

### controller

```go
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

```

### apiv1

```go
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

```

