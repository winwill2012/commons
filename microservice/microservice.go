package microservice

import (
	"errors"
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"net"
	"strings"
)

type ConsulRegistry struct {
	ConsulAddrs               []string
	ServiceName               string
	Port                      int
	Tags                      []string
	HealthCheckUrl            string
	DeregisterTimeoutInSecond int
}

func GetServiceInstances(consulAddress string, consulPort int, serviceName string) (instances []string, err error) {
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", consulAddress, consulPort)
	client, err := api.NewClient(config)
	if err != nil {
		return
	}
	services, _, err := client.Catalog().Service(serviceName, "", nil)
	for _, service := range services {
		instances = append(instances, fmt.Sprintf("%s:%d", service.ServiceAddress, service.ServicePort))
	}
	return
}

func RegisterConsul(consulRegistry *ConsulRegistry) (err error) {
	config := api.DefaultConfig()
	for _, consulAddr := range consulRegistry.ConsulAddrs {
		if !strings.Contains(consulAddr, ":") {
			err = errors.New("Consul地址不合法，必须IP:Port格式")
			return
		}
		parts := strings.Split(consulAddr, ":")
		config.Address = fmt.Sprintf("%s:%s", parts[0], parts[1])
		client, e := api.NewClient(config)
		if e != nil {
			err = e
			return
		}
		localIp := GetLocalIp()
		ipAddress := strings.ReplaceAll(localIp, ".", "-")
		serviceID := fmt.Sprintf("%s-%s-%d", consulRegistry.ServiceName, ipAddress, consulRegistry.Port)
		registration := new(api.AgentServiceRegistration)
		registration.ID = serviceID
		registration.Name = consulRegistry.ServiceName
		registration.Address = localIp
		registration.Port = consulRegistry.Port
		registration.Tags = consulRegistry.Tags
		registration.Check = &api.AgentServiceCheck{ // 健康检查
			HTTP:                           fmt.Sprintf("http://%s:%d%s", localIp, consulRegistry.Port, consulRegistry.HealthCheckUrl),
			Timeout:                        "3s",
			Interval:                       "5s",
			DeregisterCriticalServiceAfter: fmt.Sprintf("%ds", consulRegistry.DeregisterTimeoutInSecond),
		}
		e = client.Agent().ServiceRegister(registration)
		if e != nil {
			err = e
			return
		}
		log.Printf("服务[%s]注册到[%s:%s]成功", consulRegistry.ServiceName, parts[0], parts[1])
	}
	return
}

func GetLocalIp() string {
	conn, _ := net.Dial("udp", "m.9ji.com:8080")
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}
