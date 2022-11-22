package consul

import (
	"errors"
	"math/rand"
	"net"
	"strconv"
	"time"
)

type DiscoveryClient interface {

	/**
	 * Gets all ServiceInstances associated with a particular serviceId.
	 * @param serviceId The serviceId to query.
	 * @return A List of ServiceInstance.
	 */
	GetInstances(serviceId string) ([]ServiceInstance, error)

	/**
	 * @return All known service IDs.
	 */
	GetServices() ([]string, error)
}

type ServiceInstance interface {

	// return The unique instance ID as registered.
	GetInstanceId() string

	// return The service ID as registered.
	GetServiceId() string

	// return The hostname of the registered service instance.
	GetHost() string

	// return The port of the registered service instance.
	GetPort() int

	// return Whether the port of the registered service instance uses HTTPS.
	IsSecure() bool

	// return The key / value pair metadata associated with the service instance.
	GetMetadata() map[string]string
}

type DefaultServiceInstance struct {
	InstanceId string
	ServiceId  string
	Host       string
	Port       int
	Secure     bool
	Metadata   map[string]string
}

func getClientIp() (string, error) {
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

func NewDefaultServiceInstance(serviceId string, host string, port int, secure bool,
	metadata map[string]string, instanceId string) (*DefaultServiceInstance, error) {

	// 如果没有传入 IP 则获取一下，这个方法在多网卡的情况下，并不好用
	if len(host) == 0 {
		localIP, err := getClientIp()
		if err != nil {
			return nil, err
		}
		host = localIP
	}

	if len(instanceId) == 0 {
		instanceId = serviceId + "-" + strconv.FormatInt(time.Now().Unix(), 10) + "-" + strconv.Itoa(rand.Intn(9000)+1000)
	}

	return &DefaultServiceInstance{InstanceId: instanceId, ServiceId: serviceId, Host: host, Port: port, Secure: secure, Metadata: metadata}, nil
}

func (serviceInstance DefaultServiceInstance) GetInstanceId() string {
	return serviceInstance.InstanceId
}

func (serviceInstance DefaultServiceInstance) GetServiceId() string {
	return serviceInstance.ServiceId
}

func (serviceInstance DefaultServiceInstance) GetHost() string {
	return serviceInstance.Host
}

func (serviceInstance DefaultServiceInstance) GetPort() int {
	return serviceInstance.Port
}

func (serviceInstance DefaultServiceInstance) IsSecure() bool {
	return serviceInstance.Secure
}

func (serviceInstance DefaultServiceInstance) GetMetadata() map[string]string {
	return serviceInstance.Metadata
}
