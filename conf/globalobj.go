package conf

import (
	"encoding/json"
	"github.com/strokebun/gserver/iface"
	"github.com/strokebun/gserver/log"
	"io/ioutil"
	"runtime"
)

// @Description: 服务器全局配置
// @Author: StrokeBun
// @Date: 2022/1/7 9:39

type GlobalObj struct {
	// 全局Server对象
	TcpServer iface.IServer
	// 服务器名称
	Name string
	// 服务器ip
	Host string
	// 服务器监听的端口
	Port int

	// gserver版本
	Version string
	// 最大连接数
	MaxConn int
	// 数据包最大字节数
	MaxPackageSize uint32
	// 工作池大小
	WorkPoolSize uint32
	// 任务队列最大长度
	MaxWorkTaskLen uint32

}

var GlobalObject * GlobalObj

func (g *GlobalObj) Load() {
	data, err := ioutil.ReadFile("gserver.json")
	if err != nil {
		log.GlobalLogger.Println("[Warning] there is no configuration file, use default configuration")
		return
	}
	json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

// 读取配置文件形成配置类
func init() {
	GlobalObject = &GlobalObj{
		Name:           "gServerApp",
		Host:           "0.0.0.0",
		Port:           6023,
		Version:        "V1.0",
		MaxConn:        1024,
		MaxPackageSize: 4096,
		WorkPoolSize: uint32(runtime.GOMAXPROCS(runtime.NumCPU())),
		MaxWorkTaskLen: 1024,
	}
	GlobalObject.Load()
}