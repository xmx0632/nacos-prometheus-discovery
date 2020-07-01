package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"nacos-prometheus-discovery/httputil"
	"nacos-prometheus-discovery/model"
	"os"
	"os/signal"
	"strconv"
	"time"
)

const DefaultConfigPath = "conf/config.json"

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
}

func main() {

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill)

	log.Println("Nacos Prometheus Discovery Starting ...")
	filename := DefaultConfigPath
	if len(os.Args) > 1 {
		argsWithoutProg := os.Args[1:2]
		log.Println("config file path:", argsWithoutProg)
		filename = argsWithoutProg[0]
	}

	// read config file
	configJson, configErr := ioutil.ReadFile(filename)
	if configErr != nil {
		log.Fatal("read config file error.", configErr)
	}
	config := model.Config{}
	json.Unmarshal(configJson, &config)

	// start timer
	ticker := time.NewTicker(time.Second * time.Duration(config.IntervalInSecond))
	defer ticker.Stop()
	done := make(chan bool)
	go func() {
		// listen terminate sig and stop
		s := <-c
		fmt.Println("terminated.", s)
		done <- true
	}()
	for {
		select {
		case <-done:
			fmt.Println("Exit!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
			generatePrometheusTarget(config)
		}
	}
}

func generatePrometheusTarget(config model.Config) {
	targetFilePath := config.TargetFilePath
	nacosHost := config.NacosHost
	namespaceId := config.NamespaceId
	group := config.Group
	cluster := config.Cluster

	serviceNames := GetServiceNames(nacosHost, namespaceId, group)

	// generate target json
	promJsonTargets := []model.PromTarget{}

	for i, serviceName := range serviceNames {
		log.Println()
		log.Println("service["+strconv.Itoa(i)+"] : ", serviceName)

		instancesUrl := fmt.Sprintf("%s/v1/ns/instance/list?serviceName=%s&namespaceId=%s&cluster=%s", nacosHost, serviceName, namespaceId, cluster)
		//log.Println("=== instancesUrl:", instancesUrl)

		instanceJson, ierr := httputil.Get(instancesUrl)
		if ierr != nil {
			log.Println("get instanceJson failed", ierr)
		}
		instance := model.Instance{}
		json.Unmarshal([]byte(instanceJson), &instance)

		targets := []string{}
		lables := make(map[string]string)

		hosts := instance.Hosts

		for j, host := range hosts {
			hostAddress := host.Ip + ":" + strconv.Itoa(host.Port)
			log.Println("host["+strconv.Itoa(j)+"] :", hostAddress)
			metadata := host.Metadata
			log.Println(">> metadata :")

			targets = append(targets, hostAddress)

			for key, value := range metadata {
				log.Println("["+key+"] = ", value)
				lables[key] = value
			}
		}
		pt := model.PromTarget{Labels: &lables, Targets: &targets}
		promJsonTargets = append(promJsonTargets, pt)
	}

	targetJson, jsonErr := json.MarshalIndent(promJsonTargets, "", "  ")
	if jsonErr != nil {
		log.Println("marshal json failed", jsonErr)
	}
	jsonString := string(targetJson)
	log.Println("targetJson:", jsonString)

	wfErr := ioutil.WriteFile(targetFilePath, targetJson, os.ModePerm)
	if wfErr != nil {
		log.Println("generate target file failed", wfErr)
	}
}

func GetServiceNames(nacosHost string, namespaceId string, group string) []string {
	serviceUrl := fmt.Sprintf("%s/v1/ns/service/list?pageNo=1&pageSize=10&namespaceId=%s&groupName=%s", nacosHost, namespaceId, group)
	log.Println("=== serviceUrl:", serviceUrl)

	services, serr := httputil.Get(serviceUrl)
	if serr != nil {
		log.Println("get service failed", serr)
	}

	service := model.Service{}
	json.Unmarshal([]byte(services), &service)
	return service.Doms
}
