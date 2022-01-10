package log

import (
	"log"
	"os"
)

// @Description: 全局日志配置
// @Author: StrokeBun
// @Date: 2022/1/10 15:35
var GlobalLogger = log.New(os.Stderr, "", log.LstdFlags)

