package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/zzzhangjian/go_ece/app/api/ece/post"
)

func init() {
	s := g.Server()
	ctlPost := new(post.Controller)
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.ALL("/", ctlPost)
	})
}
