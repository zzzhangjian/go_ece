package response

import (
	"github.com/gogf/gf/net/ghttp"
)

type PageRequest struct {
	Current int `json:"current"`
	Size    int `json:"size"`
}

type PageResult struct {
	PageRequest
	Total   int         `json:"total"`
	Records interface{} `json:"records"`
}

// 数据返回通用JSON数据结构
type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// 标准返回结果数据结构封装。
func Json(r *ghttp.Request, code int, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	_ = r.Response.WriteJson(JsonResponse{
		Code:    code,
		Message: message,
		Data:    responseData,
	})
}

// 返回JSON数据并退出当前HTTP执行函数。
func JsonExit(r *ghttp.Request, err int, msg string, data ...interface{}) {
	Json(r, err, msg, data...)
	r.Exit()
}

// 返回JSON数据并退出当前HTTP执行函数。
func Ok(r *ghttp.Request, data ...interface{}) {
	OkMsg(r, "success", data...)
}

// 返回JSON数据并退出当前HTTP执行函数。
func OkMsg(r *ghttp.Request, msg string, data ...interface{}) {
	JsonExit(r, 200, msg, data...)
}

// 返回JSON数据并退出当前HTTP执行函数。
func Error(r *ghttp.Request, data ...interface{}) {
	ErrorMsg(r, "fail", data...)
}

// 返回JSON数据并退出当前HTTP执行函数。
func ErrorMsg(r *ghttp.Request, msg string, data ...interface{}) {
	JsonExit(r, -1, msg, data...)
}
