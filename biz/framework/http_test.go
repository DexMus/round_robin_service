package framework

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPost(t *testing.T) {
	ctx := context.Background()
	body := []byte(`{"key": "value"}`)

	// Test normal case
	resp, err := PostWithRetry(ctx, body, 3)
	assert.Nil(t, err)
	assert.NotNil(t, resp)

	// Test error case
	ctx = context.Background()
	body = []byte(`{"key": "value"}`)

	// Test normal case
	resp, err = PostWithRetry(ctx, body, 1)
	assert.NotNil(t, err)
	assert.Nil(t, resp)
}

func TestGetFreePorts(t *testing.T) {
	port, err := GetFreePorts(3)
	if err != nil {
		log.Println(err)
	}
	log.Println(port)
}
