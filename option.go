package gRouter

import "time"

type Option struct {
	ReadTimeout        time.Duration //读超时
	WriteTimeout       time.Duration //写超时
	IsDebug            bool
	HttpPort           int64 //http服务端口
	MaxMultipartMemory int64 //post参数占用最大内存
}

var option = &Option{
	ReadTimeout:        120 * time.Second,
	WriteTimeout:       120 * time.Second,
	IsDebug:            true,
	HttpPort:           8080,
	MaxMultipartMemory: 32 << 20, // 32 MB
}

func setOption(opt *Option) *Option {
	if opt.ReadTimeout > 0 {
		option.ReadTimeout = opt.ReadTimeout
	}
	if opt.WriteTimeout > 0 {
		option.ReadTimeout = opt.WriteTimeout
	}
	option.IsDebug = opt.IsDebug
	if opt.HttpPort > 0 {
		option.HttpPort = opt.HttpPort
	}
	if opt.MaxMultipartMemory > 0 {
		option.MaxMultipartMemory = opt.MaxMultipartMemory
	}
	return option
}
