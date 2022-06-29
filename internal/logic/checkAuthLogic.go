package logic

import (
	"context"

	"github.com/pz2147/p-rpc-1/internal/svc"
	"github.com/pz2147/p-rpc-1/prpc1"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAuthLogic {
	return &CheckAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  CheckAuth 风控检查
func (l *CheckAuthLogic) CheckAuth(in *prpc1.AuthReq) (*prpc1.AuthResp, error) {
	// todo: add your logic here and delete this line

	return &prpc1.AuthResp{}, nil
}
