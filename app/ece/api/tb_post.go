package post

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/zzzhangjian/go_ece/app/ece/service"
	"github.com/zzzhangjian/go_ece/app/model/tb_posts"
	"github.com/zzzhangjian/go_ece/library/response"
)

// 用户API管理对象
type Controller struct{}

// 首页新闻接口
func (c *Controller) Index(r *ghttp.Request) {
	var data *tb_posts.PostPageRequest
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	result := post.QueryPage(data)
	if result != nil {
		response.JsonExit(r, 1, "fail", result)
	} else {
		response.JsonExit(r, 0, "ok", result)
	}
}
