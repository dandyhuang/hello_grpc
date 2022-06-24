package main

import (
	"errors"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"net"
)
const (
	consulAddress = "localhost:8500"
	serviceId     = "111"
	consulPort = "9000"
)

func getClientIp() (string ,error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}

		}
	}

	return "", errors.New("Can not find the client ip address!")

}

func ConsulCheckHeath() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = consulAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}

	// 健康检查
	ip, err :=getClientIp()
	if err != nil {
		fmt.Println("err:", err)
	}


	a, b, _ := client.Agent().AgentHealthServiceByID(ip+"-"+consulPort)
	fmt.Println("val1:", a)
	fmt.Println("val2:", b)
	fmt.Println("ConsulCheckHeath done")
}

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
	fmt.Println(getClientIp())
	ConsulCheckHeath()
}
