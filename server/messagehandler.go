package server

import (
	"fmt"
	"github.com/strokebun/gserver/iface"
)

// @Description: 消息处理模块实现
// @Author: StrokeBun
// @Date: 2022/1/7 15:27
type MessageHandler struct {
	Apis map[uint32]iface.IRouter
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		Apis: make(map[uint32]iface.IRouter),
	}
}

func (mh *MessageHandler) DoMessageHandler(request iface.IRequest) {
	router, ok := mh.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("[WARNING] api msgId =", request.GetMsgId(), "miss")
		return
	}
	router.PreHandle(request)
	router.Handle(request)
	router.AfterHandle(request)
}

func (mh *MessageHandler) AddRouter(msgId uint32, router iface.IRouter) {
	if _, ok := mh.Apis[msgId]; ok {
		fmt.Println("[WARNING] msgId has existed...")
	}
	mh.Apis[msgId] = router
	fmt.Println("add api msgId = ", msgId, " success")
}
