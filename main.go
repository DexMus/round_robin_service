// Code generated by hertztool.

package main

import (
	"context"
	"fmt"

	"code.byted.org/middleware/hertz/byted"
	"github.com/DexMus/round_robin_service/biz/config"
	"github.com/DexMus/round_robin_service/biz/router/github/DexMus/round_robin_service"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	byted.Init()
	ctx := context.Background()
	ipPorts := config.GetConfInstance().IPPorts
	serviceIP := ipPorts.ServiceIP
	go startMainAPI(ctx, serviceIP, ipPorts.ServiceMainPort)
	startAndKeepSimpleAPI(ctx, ipPorts.InstanceCount, serviceIP)
}

func startMainAPI(ctx context.Context, ip string, port int) {
	p := fmt.Sprintf("%s:%d", ip, port)
	r := server.Default(server.WithHostPorts(p))
	round_robin_service.RoundRegister(r)
	// Prevent return, catching os.Signal or error returned by h.Run()
	r.Spin()
}
