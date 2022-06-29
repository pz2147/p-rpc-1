package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
	"os"
	"time"
)

type (
	AggregationTypes string

	EsService struct {
		indexMap    map[string]indexSetting // 索引的配置
		client      *elastic.Client         // es客户端
	}

	indexSetting struct {
		indexNames    map[string]bool
		esSetting     EsSetting
		createSetting map[string]interface{}
	}
)

func NewService(ctx context.Context, addr []string, indexNames []string, path string) *EsService {
	return NewServiceByUser(ctx, addr, indexNames, path, "", "")
}

func NewServiceByUser(ctx context.Context, addr []string, indexNames []string, path string, username string, password string) *EsService {

	coFunc := []elastic.ClientOptionFunc{
		elastic.SetSniff(false),
		// 设置监控检查时间间隔
		elastic.SetHealthcheckInterval(10 * time.Second),
		elastic.SetURL(addr...),

		// 启用gzip压缩
		elastic.SetGzip(true),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	}
	if len(username) > 0 {
		coFunc = append(coFunc, elastic.SetBasicAuth(username, password))
	}

	client, nErr := elastic.NewClient(coFunc...)
	if nErr != nil {
		logx.WithContext(ctx).Errorf("[NewServiceByUser] 链接es报错 err:%s", nErr.Error())
		return nil
	}

	service := EsService{
		client:   client,
		indexMap: map[string]indexSetting{},
	}
	//// 按照索引列表读取索引的配置文件
	//for _, indexName := range indexNames {
	//	indexJson := utils.ReadFile(path + "/" + indexName + ".json")
	//	if indexJson == nil || len(indexJson) == 0 {
	//		continue
	//	}
	//	indexJsonCfg := ""
	//	for _, str := range indexJson {
	//		indexJsonCfg += str + "\r\n"
	//	}
	//	var esSetting EsSetting
	//	err := json.Unmarshal([]byte(indexJsonCfg), &esSetting)
	//	if err != nil {
	//		continue
	//	}
	//	//if esSetting.Shards == 0 {
	//	//	esSetting.Shards = 3
	//	//}
	//	//if esSetting.RefreshInterval == "" {
	//	//	esSetting.RefreshInterval = "10s"
	//	//}
	//	createSetting := make(map[string]interface{})
	//	settings := make(map[string]interface{})
	//	//settings["number_of_shards"] = esSetting.Shards
	//	//settings["number_of_replicas"] = esSetting.Replics
	//	//settings["refresh_interval"] = esSetting.RefreshInterval
	//	createSetting["settings"] = settings
	//	createSetting["mappings"] = esSetting.Mapping
	//	indexSetting := indexSetting{
	//		indexNames:    map[string]bool{},
	//		esSetting:     esSetting,
	//		createSetting: createSetting,
	//	}
	//	service.indexMap[indexName] = indexSetting
	//	//if indexSetting.esSetting.TimeFormat == "ONE" {
	//	//	service.CheckOrCreateIndex(indexSetting.esSetting.Index, 0)
	//	//}
	//}
	//go service.bulk()
	return &service
}
