package main

import (
	_ "goframe-bluebell/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"goframe-bluebell/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
