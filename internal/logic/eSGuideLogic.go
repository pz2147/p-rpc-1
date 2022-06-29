package logic

import (
	"context"
	"errors"
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"time"

	"github.com/pz2147/p-rpc-1/internal/svc"
	"github.com/pz2147/p-rpc-1/prpc1"

	"github.com/zeromicro/go-zero/core/logx"
)

type ESGuideLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewESGuideLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ESGuideLogic {
	return &ESGuideLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ESGuide elastic教程
func (l *ESGuideLogic) ESGuide(in *prpc1.EmptyReq) (*prpc1.EmptyResp, error) {

	esConfig := l.svcCtx.Config.EsConfig
	client, cErr := elastic.NewClient(
		elastic.SetURL(esConfig.Addresses...),
		elastic.SetSniff(false),

		// 基于http base auth验证机制的账号和密码
		//elastic.SetBasicAuth(esConfig.Username, esConfig.Password)
		// 启用gzip压缩
		elastic.SetGzip(true),
		// 设置监控检查时间间隔
		elastic.SetHealthcheckInterval(10*time.Second),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		)
	if cErr != nil {
		l.Logger.Errorf("[ESGuide] 创建ecClietn报错 %s", cErr.Error())
		return nil, errors.New("报错")
	} else {
		l.Logger.Infof("[ESGuide] 链接成功")
	}




	_, iErr := client.IndexExists("www").Do(l.ctx)
	if iErr != nil {
		l.Logger.Errorf("[ESGuide] 创建ecClietn报错 %s", iErr.Error())
	}

	return &prpc1.EmptyResp{}, nil
}
