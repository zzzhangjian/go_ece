package main

import (
	"github.com/gogf/gf/frame/g"
	_ "github.com/zzzhangjian/go_ece/boot"
	_ "github.com/zzzhangjian/go_ece/router"
)

func main() {
	g.Server().Run()
}
