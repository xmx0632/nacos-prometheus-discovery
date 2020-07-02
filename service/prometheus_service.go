package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"nacos-prometheus-discovery/httputil"
	"nacos-prometheus-discovery/model"
	"os"
	"strconv"
	"strings"
)

func GeneratePrometheusTarget(config model.Config) {
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
				validKey := ReplaceInvalidChar(key)
				lables[validKey] = value
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

func ReplaceInvalidChar(key string) string {
	validKey := key
	validKey = strings.ReplaceAll(validKey, ".", "_")
	validKey = strings.ReplaceAll(validKey, "-", "_")
	return validKey
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
