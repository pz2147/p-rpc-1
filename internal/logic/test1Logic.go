package logic

import (
	"context"

	"github.com/pz2147/p-rpc-1/internal/svc"
	"github.com/pz2147/p-rpc-1/prpc1"

	"github.com/tal-tech/go-zero/core/logx"
)

type Test1Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTest1Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Test1Logic {
	return &Test1Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Test1 测试1
func (l *Test1Logic) Test1(in *prpc1.Test1Req) (*prpc1.Test1Resp, error) {
	// todo: add your logic here and delete this line

	return &prpc1.Test1Resp{}, nil
}
