package main

import (
	"context"
	"fmt"

	"code.byted.org/gopkg/logs/v2/log"
	"github.com/DexMus/round_robin_service/biz/framework"
	"github.com/DexMus/round_robin_service/biz/router/github/DexMus/round_robin_service"
	"github.com/cloudwego/hertz/pkg/app/server"
)

type Scheduler struct {
	port           int
	averageLatency int
}

func startAndKeepSimpleAPI(ctx context.Context, count int, serviceIP string) {
	ports, err := framework.GetFreePorts(count)
	log.V2.Info().With(ctx).EmplaceKVs("Free ports", ports).Emit()
	if err != nil {
		log.V2.Fatal().With(ctx).Error(err).Emit()
		panic(err)
	}
	framework.InitRoundRobin(ports)
	addrMQ := make(chan string, len(ports))
	for _, p := range ports {
		prot := p
		addr := fmt.Sprintf("%s:%d", serviceIP, prot)
		addrMQ <- addr
	}
	for {
		select {
		case addr := <-addrMQ:
			go startSimpleAPI(ctx, addr, addrMQ)
		}
	}
}

func startSimpleAPI(ctx context.Context, addr string, addrMQ chan string) {
	defer func() { addrMQ <- addr }()
	defer framework.RoutineRecovery(ctx)
	srv := server.Default(server.WithHostPorts(addr))
	round_robin_service.SimpleRegister(srv)
	// Prevent return
	srv.Spin()
}
