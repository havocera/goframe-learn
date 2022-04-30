package main

import (
	_ "helloworld/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"helloworld/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
