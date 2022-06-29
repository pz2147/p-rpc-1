package es

type (
	EsSetting struct {
		// 索引名称/前缀guild.json
		Index string `json:"index"`
		//// 索引类型 YEAR 按年创建 MONTH 按月创建 DATE 按日创建 ONE 只有一份
		//TimeFormat string `json:"timeFormat"`
		//// 几个shard
		//Shards int `json:"shards"`
		//// 几份复制
		//Replics int `json:"replics"`
		//// 刷新时间，数据入库后最多多久能查询到，单位为 ms/s/m，对应毫秒/秒/分钟
		//RefreshInterval string `json:"refreshInterval"`
		// mapping设置
		Mapping map[string]interface{} `json:"mapping"`
	}
)
