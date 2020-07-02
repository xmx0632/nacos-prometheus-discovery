package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"nacos-prometheus-discovery/model"
	"nacos-prometheus-discovery/service"
	"os"
	"os/signal"
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
			service.GeneratePrometheusTarget(config)
		}
	}
}
