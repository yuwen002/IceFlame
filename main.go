package main

import (
	_ "ice_flame/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"ice_flame/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
