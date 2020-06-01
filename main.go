package main

import (
	_ "git.irss.cn/zhang/smit_cloud_ece/boot"
	_ "git.irss.cn/zhang/smit_cloud_ece/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
