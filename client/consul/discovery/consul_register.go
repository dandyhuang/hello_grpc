package discovery

import (
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"strconv"
	"sync"
)

type consulServiceRegistry struct {
	serviceInstances map[string][]ServiceInstance
	client           api.Client
	rwLock           sync.RWMutex
}

func (c consulServiceRegistry) FirstInstances(serviceId string) error {
	c.rwLock.RLock()
	if len(c.serviceInstances[serviceId]) == 0 {
		c.rwLock.RUnlock()
		err := c.HealthCheckServices(serviceId)
		if err != nil {
			return err
		}
		return nil
	}
	c.rwLock.RUnlock()
	return nil
}

func (c consulServiceRegistry) GetInstances(serviceId string) ([]ServiceInstance, error) {
	c.rwLock.RLock()
	defer c.rwLock.RLock()
	var ins []ServiceInstance
	copy(ins, c.serviceInstances[serviceId])
	return ins, nil
}

func (c consulServiceRegistry) GetServices() ([]string, error) {
	services, _, _ := c.client.Catalog().Services(nil)
	result := make([]string, len(services))
	index := 0
	for serviceName, _ := range services {
		result[index] = serviceName
		index++
	}
	return result, nil
}

func (c consulServiceRegistry) HealthCheckServices(serviceId string) error {
	catalogService, _, err := c.client.Catalog().Service(serviceId, "", nil)
	if err != nil {
		return err
	}
	sSize := len(catalogService)
	if sSize > 0 {
		result := make(map[string]ServiceInstance, sSize)
		for _, sever := range catalogService {
			s := DefaultServiceInstance{
				InstanceId: sever.ServiceID,
				ServiceId:  sever.ServiceName,
				Host:       sever.Address,
				Port:       sever.ServicePort,
				Metadata:   sever.ServiceMeta,
			}
			fmt.Println("serverid:", sever.ServiceID)
			result[sever.ServiceID] = s
		}
		mSerIns := make([]ServiceInstance, sSize)

		checks, _, err := c.client.Health().Checks(serviceId, nil)
		if err != nil {
			return err
		}
		for _, check := range checks {
			if check.Status != "passing" {
				// check.Status: 10.194.19.151-18801 &{cpd-sorting-ctr-prd-10-194-19-151.v-bj-4.vivo.lan 10.194.19.151-18801
				//   schedule_comm_801_prd critical  dial tcp 10.194.19.151:18801: connect: connection refused 10.194.19.151-18801
				//  schedule_comm_801_prd [] tcp   { map[]    false   false 0s 0s 0s 0 0 0} 3738010626 3738010626}
				//    client_test.go:17: <nil>
				err := c.client.Agent().ServiceDeregister(check.ServiceID)
				fmt.Println("deregister:", err)
			}
			if check.Status == "passing" {
				fmt.Println("check.Status:", check.ServiceID, "node:", check.Node)
				mSerIns = append(mSerIns, result[check.ServiceID])
			}
		}
		c.rwLock.Lock()
		c.serviceInstances[serviceId] = mSerIns
		c.rwLock.Unlock()
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

	return &consulServiceRegistry{
		serviceInstances: make(map[string][]ServiceInstance, 0),
		client:           *client,
	}, nil
}
