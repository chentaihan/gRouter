package gRouter

import "time"

type Option struct {
	ServerReadTimeout  time.Duration //读超时
	ServerWriteTimeout time.Duration //写超时
	IsDebug            bool
	HttpPort           int64 //http服务端口
	MaxMultipartMemory int64 //post参数占用最大内存
}

var Opt = &Option{
	ServerReadTimeout:  120 * time.Second,
	ServerWriteTimeout: 120 * time.Second,
	IsDebug:            true,
	HttpPort:           8080,
	MaxMultipartMemory: 32 << 20, // 32 MB
}
