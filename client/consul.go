package main

import (
	"errors"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc/resolver"
	"net"
	"strconv"
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

func ConsulFindServer1() {
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
func HealthCheck(name string) {
	config := consulapi.DefaultConfig()
	config.Address = consulAddress
	client, _ := consulapi.NewClient(config)
	var lastIndex uint64
	for {
		fmt.Println("health start")
		services, metainfo, err := client.Health().Service(name, "", true, &consulapi.QueryOptions{WaitIndex: lastIndex})
		if err != nil {
			fmt.Printf("error retrieving instances from Consul: %v", err)
		}

		lastIndex = metainfo.LastIndex
		var newAddrs []resolver.Address
		for _, service := range services {
			addr := fmt.Sprintf("%v:%v", service.Service.Address, service.Service.Port)
			newAddrs = append(newAddrs, resolver.Address{Addr: addr})
		}
		fmt.Printf("adding service addrs\n")
		fmt.Printf("newAddrs: %v lastIndex: %d\n", newAddrs, lastIndex)
		// cr.cc.NewAddress(newAddrs)
		// cr.cc.UpdateState(resolver.State{Addresses: newAddrs})
		// cr.cc.NewServiceConfig(cr.name)
	}
}
func ConsulFindServer(scheduleSchema string) {

	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = consulAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		fmt.Println("consul client error : ", err)
	}
	var lastIndex uint64
	services, metainfo, err := client.Health().Service(scheduleSchema, "", true, &consulapi.QueryOptions{
		WaitIndex: lastIndex, // 同步点，这个调用将一直阻塞，直到有新的更新
	})
	if err != nil {
		fmt.Println("error retrieving instances from Consul: %v", err)
	}
	lastIndex = metainfo.LastIndex

	addrs := map[string]struct{}{}
	for _, service := range services {
		fmt.Println("service.Service.Address:", service.Service.Address, "service.Service.Port:", service.Service.Port)
		addrs[net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port))] = struct{}{}
	}
}

func main() {
	ConsulFindServer("schedule_comm_824_prd")
	HealthCheck("schedule_comm_824_prd")
	fmt.Println(getClientIp())
	//conn, err := grpc.Dial(
	//	// consul服务
	//	"consul://192.168.1.11:8500/hello?wait=14s",
	//	// 指定round_robin策略
	//	grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//)
	ConsulCheckHeath()
}
