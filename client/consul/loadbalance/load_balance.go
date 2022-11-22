package loadbalance

import (
	"errors"
	"hello_grpc/client/consul/discovery"
	"math/rand"
	"net"
)

type LoadBalance interface {
	choose(serviceId string, instances []discovery.ServiceInstance) (discovery.ServiceInstance, error)
}

type FirstLoadBalance struct {
}

func (f FirstLoadBalance) choose(serviceId string, instances []discovery.ServiceInstance) (discovery.ServiceInstance, error) {
	if instances != nil && len(instances) > 0 {
		return instances[0], nil
	}
	return nil, errors.New("no available instance")
}

func (f FirstLoadBalance) Pick(serviceId string, instances []discovery.ServiceInstance) (discovery.ServiceInstance, error) {
	if instances != nil && len(instances) > 0 {
		return instances[0], nil
	}
	for {
		if len(instances) == 0 {
			break
		}

		i := rand.Int() % len(instances)
		endpoint := instances[i]

		// Try to connect
		_, err := net.Dial("tcp", endpoint.GetEndPoint())
		if err != nil {
			// reg.Failure(serviceName, serviceVersion, endpoint, err)
			// Failure: remove the endpoint from the current list and try again.
			instances = append(instances[:i], instances[i+1:]...)
			continue
		}
		// Success: return the connection.
		return endpoint, nil
	}

	return nil, errors.New("no available instance")
}
