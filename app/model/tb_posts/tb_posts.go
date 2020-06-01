// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package tb_posts

import (
	"git.irss.cn/zhang/smit_cloud_ece/app/ece/service"
	"github.com/gogf/gf/frame/g"
)

// Fill with you ideas below.
// 分页查询
func Page(request post.PostPageRequest) (*post.PageResult, error) {
	condition := &g.Map{
		"post_type": request.Type,
	}
	total, err := Model.Where(condition).Count()
	if err != nil {
		return nil, err
	}
	list, err := Model.Where(condition).Limit(request.PageSize).Offset((request.PageNum - 1) * request.PageSize).FindAll()
	if err != nil {
		return nil, err
	}
	result := &post.PageResult{
		PageRequest: request.PageRequest,
		Total:       total,
		List:        list,
	}
	return result, err
}
