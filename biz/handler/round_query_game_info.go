// Code generated by hertztool.

package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"code.byted.org/gopkg/logs/v2/log"
	"code.byted.org/middleware/hertz/pkg/app"
	"code.byted.org/middleware/hertz_ext/v2/binding"
	"github.com/DexMus/round_robin_service/biz/constance"
	"github.com/DexMus/round_robin_service/biz/framework"
	"github.com/DexMus/round_robin_service/biz/model/github/DexMus/round_robin_service"
	"github.com/bytedance/sonic"
)

// RoundQueryGameInfo .
// @router /api/game_center/v1/meta/round_query/ [POST]
func RoundQueryGameInfo(ctx context.Context, c *app.RequestContext) {
	// Step 1: bind request and validation check
	var (
		err error
		req round_robin_service.RoundQueryGameInfoReq
	)
	err = binding.BindAndValidate(c, &req)
	if err != nil {
		log.V2.Error().With(ctx).Error(err).Emit()
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	// Step 2: Get round robin index and Post address
	resp := new(round_robin_service.RoundQueryGameInfoResp)
	data, err := json.Marshal(req)
	if err != nil {
		log.V2.Error().With(ctx).Error(err).Emit()
		return
	}

	// Step 3: Post simple API with retry
	body, err := framework.PostWithRetry(ctx, data, constance.RetryTimes)
	if err != nil {
		log.V2.Error().With(ctx).Error(err).Emit()
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if err = sonic.Unmarshal(body, &resp); err != nil {
		log.V2.Error().With(ctx).Error(err).Emit()
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	err = binding.WriteHeader(c, resp)
	if err != nil {
		log.V2.Error().With(ctx).Error(err).Emit()
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(200, resp)
}
