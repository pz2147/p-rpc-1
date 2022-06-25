package logic

import (
	"context"

	"github.com/pz2147/p-rpc-1/internal/svc"
	"github.com/pz2147/p-rpc-1/prpc1"

	"github.com/tal-tech/go-zero/core/logx"
)

type Test2Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTest2Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Test2Logic {
	return &Test2Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Test2 测试2
func (l *Test2Logic) Test2(in *prpc1.Test2Req) (*prpc1.Test2Resp, error) {
	// todo: add your logic here and delete this line

	return &prpc1.Test2Resp{}, nil
}
