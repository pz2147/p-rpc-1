package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	EsConfig EsConfig
	JwtAuth JwtAuth
}

type JwtAuth struct {
	AccessSecret string
	AccessExpire int64
}

// EsConfig ES配置
type EsConfig struct {
	Addresses  []string
	ConfigPath string
	Indexes    []string
	Username   string
	Password   string
}
