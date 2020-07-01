package model

//{"doms":["ms-auth","ms-monitor","ms-gateway"],"count":3}
type Service struct {
	Doms  []string `json:"doms"`
	Count string   `json:"count"`
}

//{"hosts":
//[
//{"ip":"10.170.0.11","port":9999,"valid":true,"healthy":true,"marked":false,"instanceId":"10.170.0.11#9999#DEFAULT#DEFAULT_GROUP@@ms-gateway",
//"metadata":
//{"preserved.register.source":"SPRING_CLOUD","version":"V1","target-version":"V1"},
//"enabled":true,"weight":1.0,"clusterName":"DEFAULT","serviceName":"ms-gateway","ephemeral":true}
//],
//"dom":"ms-gateway",
//"name":"DEFAULT_GROUP@@ms-gateway","cacheMillis":3000,"lastRefTime":1593590661828,
//"checksum":"8d0392053b7bffea89a7224b8cc782b4","useSpecifiedURL":false,
//"clusters":"","env":"","metadata":{}}
type Instance struct {
	Hosts []Host `json:"hosts"`
	//Dom             string                 `json:"dom"`
	//Name            string                 `json:"name"`
	//CacheMillis     int                    `json:"cacheMillis"`
	//LastRefTime     string                 `json:"lastRefTime"`
	//Checksum        string                 `json:"checksum"`
	//UseSpecifiedURL bool                   `json:"useSpecifiedURL"`
	//Env             string                 `json:"env"`
	//Clusters        string                 `json:"clusters"`
	//Metadata        map[string]interface{} `json:"metadata"`
}

type Host struct {
	Ip          string            `json:"ip"`
	Port        int               `json:"port"`
	Metadata    map[string]string `json:"metadata"`
	ServiceName string            `json:"serviceName"`
	//Valid       bool                   `json:"valid"`
	//Healthy     bool                   `json:"healthy"`
	//Marked      bool                   `json:"marked"`
	//InstanceId  string                 `json:"instanceId"`
	//Enabled     bool                   `json:"enabled"`
	//Weight      float32                `json:"weight"`
	//ClusterName string                 `json:"clusterName"`
	//Ephemeral   bool                   `json:"ephemeral"`
}
