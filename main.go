package main

import (
	_ "smit_cloud_ece/boot"
	_ "smit_cloud_ece/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
