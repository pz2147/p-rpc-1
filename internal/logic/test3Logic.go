package logic

import (
	"context"

	"github.com/pz2147/p-rpc-1/internal/svc"
	"github.com/pz2147/p-rpc-1/prpc1"

	"github.com/tal-tech/go-zero/core/logx"
)

type Test3Logic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTest3Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Test3Logic {
	return &Test3Logic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  Test3 测试3
func (l *Test3Logic) Test3(in *prpc1.Test3Req) (*prpc1.Test3Resp, error) {
	// todo: add your logic here and delete this line

	return &prpc1.Test3Resp{}, nil
}
