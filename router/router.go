package router

import (
	"git.irss.cn/zhang/smit_cloud_ece/app/ece/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/news", post.Controller{}.Index)
	})
}
