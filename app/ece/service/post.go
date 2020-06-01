package post

import (
	"github.com/zzzhangjian/go_ece/app/model/tb_posts"
	"github.com/zzzhangjian/go_ece/library/response"
)

// 分页请求
func QueryPage(data *tb_posts.PostPageRequest) *response.PageResult {
	result, err := tb_posts.Page(*data)
	if err != nil {
		return nil
	} else {
		return result
	}
}
