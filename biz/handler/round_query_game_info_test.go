package handler

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"github.com/stretchr/testify/mock"
)

type MockFramework struct {
	mock.Mock
}

func (m *MockFramework) PostWithRetry(ctx context.Context, data []byte, retryTimes int) ([]byte, error) {
	args := m.Called(ctx, data, retryTimes)
	return args.Get(0).([]byte), args.Error(1)
}

func TestRoundQueryGameInfo(t *testing.T) {
	mockFramework := new(MockFramework)

	h := server.Default(server.WithHostPorts("127.0.0.1:8080"))

	h.POST("/api/game_center/v1/meta/round_query/", RoundQueryGameInfo)

	c, _ := ut.CreateTestContext(http.MethodPost, "/api/game_center/v1/meta/round_query/", nil)
	c.Request.SetBodyString("invalid request")
	RoundQueryGameInfo(context.Background(), c)
	assert.Equal(t, http.StatusBadRequest, c.Response.StatusCode())

	mockFramework.On("PostWithRetry", mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("post error"))
	c, _ = ut.CreateTestContext(http.MethodPost, "/api/game_center/v1/meta/round_query/", nil)
	c.Request.SetBodyString(`{"key": "value"}`)
	RoundQueryGameInfo(context.Background(), c)
	assert.Equal(t, http.StatusInternalServerError, c.Response.StatusCode())

	mockFramework.On("PostWithRetry", mock.Anything, mock.Anything, mock.Anything).Return([]byte(`{"invalid": "json"}`), nil)
	c, _ = ut.CreateTestContext(http.MethodPost, "/api/game_center/v1/meta/round_query/", nil)
	c.Request.SetBodyString(`{"key": "value"}`)
	RoundQueryGameInfo(context.Background(), c)
	assert.Equal(t, http.StatusInternalServerError, c.Response.StatusCode())

	mockFramework.On("PostWithRetry", mock.Anything, mock.Anything, mock.Anything).Return([]byte(`{"result": "success"}`), nil)
	c, _ = ut.CreateTestContext(http.MethodPost, "/api/game_center/v1/meta/round_query/", nil)
	c.Request.SetBodyString(`{"key": "value"}`)
	RoundQueryGameInfo(context.Background(), c)
	assert.Equal(t, http.StatusOK, c.Response.StatusCode())
}
