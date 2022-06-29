package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	EsConfig EsConfig
}

// EsConfig ES配置
type EsConfig struct {
	Addresses  []string
	ConfigPath string
	Indexes    []string
	Username   string
	Password   string
}
