package consul

import (
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"strconv"
)

type consulServiceRegistry struct {
	serviceInstances     map[string]map[string]ServiceInstance
	client               api.Client
	localServiceInstance ServiceInstance
}

func (c consulServiceRegistry) GetInstances(serviceId string) ([]ServiceInstance, error) {
	catalogService, _, _ := c.client.Catalog().Service(serviceId, "", nil)
	fmt.Println(catalogService)
	if len(catalogService) > 0 {
		result := make([]ServiceInstance, len(catalogService))
		for index, sever := range catalogService {
			s := DefaultServiceInstance{
				InstanceId: sever.ServiceID,
				ServiceId:  sever.ServiceName,
				Host:       sever.Address,
				Port:       sever.ServicePort,
				Metadata:   sever.ServiceMeta,
			}
			result[index] = s
		}
		return result, nil
	}
	return nil, nil
}

func (c consulServiceRegistry) GetServices() ([]string, error) {
	services, _, _ := c.client.Catalog().Services(nil)
	fmt.Println("services:", services)
	result := make([]string, len(services))
	index := 0
	for serviceName, _ := range services {
		result[index] = serviceName
		index++
	}
	return result, nil
}

func (c consulServiceRegistry) HealthCheckServices(serviceId string) error {
	catalogService, _, _ := c.client.Catalog().Service(serviceId, "", nil)
	if len(catalogService) > 0 {
		checks, _, _ := c.client.Health().Checks(serviceId, nil)
		for _, check := range checks {
			if check.Status != "passing" {
				// check.Status: 10.194.19.151-18801 &{cpd-sorting-ctr-prd-10-194-19-151.v-bj-4.vivo.lan 10.194.19.151-18801
				//   schedule_comm_801_prd critical  dial tcp 10.194.19.151:18801: connect: connection refused 10.194.19.151-18801
				//  schedule_comm_801_prd [] tcp   { map[]    false   false 0s 0s 0s 0 0 0} 3738010626 3738010626}
				//    client_test.go:17: <nil>
				fmt.Println("check.Status:", check.ServiceID, check)
				err := c.client.Agent().ServiceDeregister(check.ServiceID)
				fmt.Println("deregister:", err)
			}
		}
	}
	return nil
}

// new a consulServiceRegistry instance
// token is optional
func NewConsulServiceRegistry(host string, port int, token string) (*consulServiceRegistry, error) {
	if len(host) < 3 {
		return nil, errors.New("check host")
	}

	if port <= 0 || port > 65535 {
		return nil, errors.New("check port, port should between 1 and 65535")
	}

	config := api.DefaultConfig()
	config.Address = host + ":" + strconv.Itoa(port)
	config.Token = token
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &consulServiceRegistry{client: *client}, nil
}
