package test

import (
	"flag"
	"github.com/pz2147/p-rpc-1/internal/config"
	"github.com/pz2147/p-rpc-1/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"testing"
)

var configFile = flag.String("f", "../etc/pRpc1.yaml", "the config file")
var testCtx *svc.ServiceContext

func setup() {
	var c config.Config

	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	//logx.MustSetup(c.LogConf)

	testCtx = ctx
}

func teardown() {
	print("After all tests\n")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	print(code)
}
