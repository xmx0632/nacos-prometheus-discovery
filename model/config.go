package model

type Config struct {
	LogLevel         string `json:logLevel`         // info
	IntervalInSecond int    `json:intervalInSecond` // 60
	TargetFilePath   string `json:targetFilePath`   // "target/gen-target.json"
	NacosHost        string `json:nacosHost`        // "http://test1:8684/nacos"
	NamespaceId      string `json:namespaceId`      // "prod"
	Group            string `json:group`            // "DEFAULT_GROUP"
	Cluster          string `json:cluster`          // "DEFAULT"
}
