package post

import (
	"git.irss.cn/zhang/smit_cloud_ece/app/model/tb_posts"
)

type PageRequest struct {
	PageNum  int
	PageSize int
}

type PageResult struct {
	PageRequest
	Total int
	List  interface{}
}

type Type string

const (
	AD     Type = "AD"
	Banner Type = "BANNER"
	Notice Type = "NOTICE"
)

type PostPageRequest struct {
	PageRequest
	Type Type
}

// 分页请求
func QueryPage(data *PostPageRequest) *PageResult {
	result, err := tb_posts.Page(*data)
	if err != nil {
		return nil
	} else {
		return result
	}
}
