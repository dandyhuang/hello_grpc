package consul

import "testing"

func TestConsulServiceDiscovery(t *testing.T) {
	host := "127.0.0.1"
	port := 8500
	token := ""
	registryDiscoveryClient, err := NewConsulServiceRegistry(host, port, token)
	if err != nil {
		panic(err)
	}

	t.Log(registryDiscoveryClient.GetServices())

	t.Log(registryDiscoveryClient.GetInstances("schedule_comm_801_prd"))
	t.Log(registryDiscoveryClient.HealthCheckServices("schedule_comm_801_prd"))
}
