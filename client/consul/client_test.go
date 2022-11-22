package consul

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"hello_grpc/client/consul/discovery"
	"hello_grpc/client/consul/loadbalance"
	"testing"
)

func SendPost(client *resty.Client, i discovery.ServiceInstance) {
	body := `{"imei":"86428105997713","gaid":"86428105997713","req_id":"113113","scene_id":10001,"scene_subalg":"vre001","channel_id":0,"personal_switch":"1","extra_info":{"parent_vid":"qutTvtwxiMg"},"mobil_user_model":{"expose_list":[{"item":"item1","dur":10000000, "ts":100000000000000000}]}}`
	resp, err := client.R().EnableTrace().SetHeader("Content-Type", "application/json").
		SetBody(body).Post(i.GetEndPoint() + "/feed/predict?sceneid=10001")
	if err != nil {
		fmt.Println("post errr:", err)
		return
	}
	ti := resp.Request.TraceInfo()
	fmt.Println("Request Trace Info:")
	fmt.Println("DNSLookup:", ti.DNSLookup)
	fmt.Println("ConnTime:", ti.ConnTime)
	fmt.Println("TCPConnTime:", ti.TCPConnTime)
	fmt.Println("TLSHandshake:", ti.TLSHandshake)
	fmt.Println("ServerTime:", ti.ServerTime)
	fmt.Println("ResponseTime:", ti.ResponseTime)
	fmt.Println("TotalTime:", ti.TotalTime)
	fmt.Println("IsConnReused:", ti.IsConnReused)
	fmt.Println("IsConnWasIdle:", ti.IsConnWasIdle)
	fmt.Println("ConnIdleTime:", ti.ConnIdleTime)
	fmt.Println("RequestAttempt:", ti.RequestAttempt)
	fmt.Println("RemoteAddr:", ti.RemoteAddr.String())
}

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
	i, err := f.Pick("schedule_comm_801_prd", ins)
	t.Log()
	t.Log(f.Pick("schedule_comm_801_prd", ins))
	t.Log(f.Pick("schedule_comm_801_prd", ins))
	t.Log(f.Pick("schedule_comm_801_prd", ins))
	t.Log(f.Pick("schedule_comm_801_prd", ins))
	t.Log(f.Pick("schedule_comm_801_prd", ins))
	client := resty.New()
	SendPost(client, i)
	SendPost(client, i)
	SendPost(client, i)
	SendPost(client, i)
}
