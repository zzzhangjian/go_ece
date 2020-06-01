package post

import (
	"git.irss.cn/zhang/smit_cloud_ece/app/ece/service"
	"git.irss.cn/zhang/smit_cloud_ece/library/response"
	"github.com/gogf/gf/net/ghttp"
)

// 用户API管理对象
type Controller struct{}

// 首页新闻接口
func (c *Controller) Index(r *ghttp.Request) {
	var data *post.PostPageRequest
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
