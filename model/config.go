package model

type Config struct {
	LogLevel         string `json:logLevel`         // info
	NacosHost        string `json:nacosHost`        // "http://test1:8684/nacos"
	IntervalInSecond int    `json:intervalInSecond` // 60
	NamespaceId      string `json:namespaceId`      // "prod"
	TargetFilePath   string `json:targetFilePath`   // "target/gen-target.json"
	DataId           string `json:dataId`           // "application-prod.yml"
	Group            string `json:group`            // "DEFAULT_GROUP"
	Cluster          string `json:cluster`          // "DEFAULT"
	Mode             string `json:mode`             // "config|service"
}
