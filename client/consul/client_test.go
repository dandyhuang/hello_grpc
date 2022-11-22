package consul

import (
	"hello_grpc/client/consul/discovery"
	"hello_grpc/client/consul/loadbalance"
	"testing"
)

func TestConsulServiceDiscovery(t *testing.T) {
	host := "127.0.0.1"
	port := 8500
	token := ""
	registryDiscoveryClient, err := discovery.NewConsulServiceRegistry(host, port, token)
	if err != nil {
		panic(err)
	}

	t.Log(registryDiscoveryClient.GetServices())

	t.Log(registryDiscoveryClient.HealthCheckServices("schedule_comm_801_prd"))
	ins, err := registryDiscoveryClient.GetInstances("schedule_comm_801_prd")
	var f loadbalance.FirstLoadBalance
	t.Log(f.Pick("schedule_comm_801_prd", ins))
	t.Log(f.Pick("schedule_comm_801_prd", ins))
	t.Log(f.Pick("schedule_comm_801_prd", ins))
	t.Log(f.Pick("schedule_comm_801_prd", ins))
	t.Log(f.Pick("schedule_comm_801_prd", ins))
	t.Log(f.Pick("schedule_comm_801_prd", ins))
}
