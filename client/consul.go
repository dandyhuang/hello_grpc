package main

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
)
const (
	consulAddress = "localhost:8500"
	serviceId     = "111"
)
func ConsulFindServer() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = consulAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	// 获取所有service
	services, _ := client.Agent().Services()
	for _, value := range services {
		fmt.Println("address:", value.Address)
		fmt.Println("port:", value.Port)
	}

	fmt.Println("=================================")
	// 获取指定service
	service, _, err := client.Agent().Service(serviceId, nil)
	if err == nil {
		fmt.Println("address:", service.Address)
		fmt.Println("port:", service.Port)
	}
	if err == nil {
		fmt.Println("ConsulFindServer done")
	}
}

func main() {
	ConsulFindServer()
}
