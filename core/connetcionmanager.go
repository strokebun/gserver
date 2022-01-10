package core

import (
	"errors"
	"github.com/strokebun/gserver/iface"
	"github.com/strokebun/gserver/log"
	"sync"
)

// @Description: 连接管理模块实现
// @Author: StrokeBun
// @Date: 2022/1/8 17:09
type ConnectionManager struct {
	// 当前所有连接
	connections map[uint32]iface.IConnection
	// 保护连接的读写锁
	lock sync.RWMutex
}

func NewConnectionManager() *ConnectionManager{
	return &ConnectionManager{
		connections: make(map[uint32]iface.IConnection),
	}
}

func (cm *ConnectionManager) Add(connection iface.IConnection)  {
	cm.lock.Lock()
	cm.connections[connection.GetConnectionId()] = connection
	cm.lock.Unlock()
	log.GlobalLogger.Println("connection add to ConnManager successfully: conn num =", cm.ConnNum())
}

func (cm *ConnectionManager) Remove(connection iface.IConnection) {
	cm.lock.Lock()
	connectionId := connection.GetConnectionId()
	delete(cm.connections, connectionId)
	cm.lock.Unlock()

	log.GlobalLogger.Println("connection Remove ConnID=", connectionId, "successfully: conn num =", cm.ConnNum())
}

func (cm *ConnectionManager)  Get(connID uint32) (iface.IConnection, error) {
	cm.lock.RLock()
	defer cm.lock.RUnlock()

	if conn, ok := cm.connections[connID]; ok {
		return conn, nil
	}
	return nil, errors.New("connection not found")
}

func (cm *ConnectionManager) ConnNum() int {
	cm.lock.RLock()
	defer cm.lock.RUnlock()
	return len(cm.connections)
}

func (cm *ConnectionManager) Clear() {
	cm.lock.Lock()
	defer cm.lock.Unlock()

	//停止并删除全部的连接信息
	for connID, conn := range cm.connections {
		conn.Stop()
		delete(cm.connections, connID)
	}
	log.GlobalLogger.Println("Clear All Connections successfully")
}