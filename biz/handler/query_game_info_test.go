package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/DexMus/round_robin_service/biz/model/github/DexMus/round_robin_service"
	"github.com/stretchr/testify/mock"
)

// 模拟 binding 包的行为
type MockBinding struct {
	mock.Mock
}

func (m *MockBinding) BindAndValidate(c interface{}, req interface{}) error {
	args := m.Called(c, req)
	return args.Error(0)
}

func (m *MockBinding) WriteHeader(c interface{}, resp interface{}) error {
	args := m.Called(c, resp)
	return args.Error(0)
}

func TestQueryGameInfo(t *testing.T) {
	mockBinding := new(MockBinding)
	mockCtx := new(MockRequestContext)

	mockBinding.On("BindAndValidate", mockCtx, &round_robin_service.QueryGameInfoReq{}).Return(errors.New("binding error"))
	QueryGameInfo(context.Background(), mockCtx)
	mockCtx.AssertCalled(t, "String", 400, "binding error")

	mockBinding.ExpectedCalls = nil
	mockCtx.ExpectedCalls = nil

	mockBinding.On("BindAndValidate", mockCtx, &round_robin_service.QueryGameInfoReq{}).Return(nil)
	mockBinding.On("WriteHeader", mockCtx, &round_robin_service.QueryGameInfoResp{}).Return(errors.New("write header error"))
	QueryGameInfo(context.Background(), mockCtx)
	mockCtx.AssertCalled(t, "String", 500, "write header error")

	mockBinding.ExpectedCalls = nil
	mockCtx.ExpectedCalls = nil

	mockBinding.On("BindAndValidate", mockCtx, &round_robin_service.QueryGameInfoReq{}).Return(nil)
	mockBinding.On("WriteHeader", mockCtx, &round_robin_service.QueryGameInfoResp{}).Return(nil)
	QueryGameInfo(context.Background(), mockCtx)
	mockCtx.AssertCalled(t, "JSON", 200, &round_robin_service.QueryGameInfoResp{})
}
