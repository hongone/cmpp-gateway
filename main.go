package main

import (
	"github.com/JoeCao/cmpp-gateway/gateway"
	"flag"
	"log"
	"encoding/json"
	"os"
	"io/ioutil"
)

func main() {
	//
	var configPath string
	var config = &gateway.Config{}

	//
	flag.StringVar(&configPath, "c", "", "配置文件路径")
	flag.Parse()

	//
	err := LoadJsonFile(configPath, config)
	if err == nil {
		log.Println("加载成功 => ", config)
	} else {
		log.Fatal("加载失败 ", configPath, " => ", err)
	}
	go gateway.StartInput(config)
	gateway.Serve(config)

}

func LoadJsonFile(filePath string, obj interface{}) error {
	data, err := ReadBytes(filePath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, obj)
	if err != nil {
		return err
	}
	return nil
}

func ReadBytes(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}