package logic

import (
	"context"

	"github.com/pz2147/p-rpc-1/internal/svc"
	"github.com/pz2147/p-rpc-1/prpc1"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//  Login 风控检查
func (l *LoginLogic) Login(in *prpc1.AuthReq) (*prpc1.AuthResp, error) {
	// todo: add your logic here and delete this line
	l.Logger.Infof("%v", in)

	return &prpc1.AuthResp{
		Uid: "123456789",
		Nickname: in.Phone,
		Pic: "pic/pic/" + in.Phone,
	}, nil
}
