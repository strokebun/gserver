package server

import "github.com/strokebun/gserver/iface"

// @Description: 默认路由模块
// @Author: StrokeBun
// @Date: 2022/1/6 19:23
type BaseRouter struct {}

func (r *BaseRouter) PreHandle(request iface.IRequest) {

}

func (r *BaseRouter) Handle(request iface.IRequest) {

}

func (r *BaseRouter) AfterHandle(request iface.IRequest) {

}