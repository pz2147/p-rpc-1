package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pz2147/p-rpc-1/internal/lib/mock"
	"time"

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

// Login 风控检查
func (l *LoginLogic) Login(in *prpc1.AuthReq) (*prpc1.AuthResp, error) {

	l.Logger.Infof("%v", in)

	mUser := mock.UserMock()

	uMap := make(map[string]interface{}, 0)
	uMap["uid"] = mUser.Uid

	st, et, rt, err := l.GenJWTToken(uMap)
	if err != nil {
		return nil, err
	}

	return &prpc1.AuthResp{
		Uid: "123456789",
		Nickname: in.Phone,
		Pic: "pic/pic/" + in.Phone,
		AccessToken: st,
		AccessExpire: et,
		RefreshAfter: rt,
	}, nil
}

// GenJWTToken 获取Jwt token
//	payloads: 加密内容
//	return
//		st: jwt token jwt token
//		et: expire time 过期时间
//		rt: refresh time 刷新时间
//		err: 错误提示
func (l *LoginLogic) GenJWTToken(payloads map[string]interface{}) (st string, et int64, rt int64, err error) {
	jwtConf := l.svcCtx.Config.JwtAuth
	accessSecret := jwtConf.AccessSecret
	accessExpire := jwtConf.AccessExpire
	now := time.Now().Unix()

	claims := make(jwt.MapClaims)
	expireTime := now + accessExpire
	claims["exp"] = expireTime
	claims["iat"] = now
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	signToken, signErr := token.SignedString([]byte(accessSecret))
	return signToken, expireTime, now + accessExpire/2, signErr
}
