package framework

import (
	"context"
	"net"
	"net/http"
	"time"

	"code.byted.org/gopkg/logs/v2/log"
	"code.byted.org/middleware/hertz/byted"
	"github.com/DexMus/round_robin_service/biz/constance"
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// @Function: Send Http Post request
// @Param: common context, Post body, Post address
// @Return: Post response body, error
func PostWithTimeout(ctx context.Context, body []byte, addr string) ([]byte, error) {
	cli, err := byted.NewClient(byted.WithAppClientOptions(client.WithDialTimeout(1 * time.Second)))
	queryReq := &protocol.Request{}
	res := &protocol.Response{}
	queryReq.SetMethod(consts.MethodPost)
	queryReq.Header.SetContentTypeBytes([]byte("application/json"))
	queryReq.SetRequestURI(addr)
	queryReq.SetBody(body)
	err = cli.DoTimeout(ctx, queryReq, res, constance.SimpleApiTimeOut)
	if err != nil {
		log.V2.Error().With(ctx).Error(err).Emit()
		return nil, err
	}
	if res.StatusCode() != http.StatusOK {
		log.V2.Error().With(ctx).EmplaceKVs("status", res.StatusCode()).Emit()
		return nil, err
	}
	log.V2.Info().With(ctx).EmplaceKVs("status", res.StatusCode(), "body", string(res.Body())).Emit()
	return res.Body(), nil
}

func PostWithRetry(ctx context.Context, body []byte, retriesTimes int) ([]byte, error) {
	if retriesTimes == 0 {
		retriesTimes = 3
	}
	for {
		addr := GetNextURL()
		body, err := PostWithTimeout(ctx, body, addr)
		log.V2.Debug().With(ctx).EmplaceKVs("addr", addr).Emit()
		if err == nil {
			return body, err
		}
		log.V2.Error().With(ctx).Error(err).Emit()
		retriesTimes--
		if retriesTimes <= 0 {
			return nil, err
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func GetFreePorts(count int) ([]int, error) {
	var ports []int
	for i := 0; i < count; i++ {
		addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
		if err != nil {
			return nil, err
		}

		l, err := net.ListenTCP("tcp", addr)
		if err != nil {
			return nil, err
		}
		defer l.Close()
		ports = append(ports, l.Addr().(*net.TCPAddr).Port)
	}
	return ports, nil
}
