package logic

import (
	"context"
	"github.com/pz2147/p-rpc-1/common/ec"
	"github.com/pz2147/p-rpc-1/internal/lib/mock"

	"github.com/pz2147/p-rpc-1/internal/svc"
	"github.com/pz2147/p-rpc-1/prpc1"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoGetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoGetLogic {
	return &UserInfoGetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// UserInfoGet 获取用户数据
func (l *UserInfoGetLogic) UserInfoGet(in *prpc1.UserInfoReq) (*prpc1.UserInfoResp, error) {

	uid := in.GetUid()
	if len(uid) == 0 {
		return nil, ec.ParamInvalid
	}

	um := mock.UserMock()
	return &prpc1.UserInfoResp{
		Uid: um.Uid,
		Nickname: um.Nickname,
		Pic: um.Pic,
	}, nil
}
