package model

//[
//{ "targets": [ "10.170.0.11:9999" ], "labels": { "Usage": "gateway", "Env": "prod", "Name": "gateway-1", "Category": "micro-service" } },
//{ "targets": [ "10.170.0.11:3000" ], "labels": { "Usage": "auth", "Env": "prod", "Name": "auth-1", "Category": "micro-service" } }
//]
type PromTarget struct {
	Targets *[]string          `json:"targets"`
	Labels  *map[string]string `json:"labels"`
}
