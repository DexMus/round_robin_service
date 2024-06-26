package middleware

import (
	"context"

	"code.byted.org/middleware/hertz/pkg/app"
)

type HertzContext struct {
	*app.RequestContext

	params    interface{}
	rawParams map[string]string
	rawBody   []byte
}

func QueryGameInfoMw() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		hertzCtx := &HertzContext{
			RequestContext: ctx,
		}
		hertzCtx.initRawParams()
	}
}

func (c *HertzContext) initRawParams() {
	c.rawParams = map[string]string{}
	c.RequestContext.Request.URI().QueryArgs().VisitAll(func(k, v []byte) {
		strK := string(k)
		if _, ok := c.rawParams[strK]; !ok {
			c.rawParams[strK] = string(v)
		}
	})

	c.RequestContext.Request.PostArgs().VisitAll(func(k, v []byte) {
		strK := string(k)
		if _, ok := c.rawParams[strK]; !ok {
			c.rawParams[strK] = string(v)
		}
	})

	for _, param := range c.RequestContext.Params {
		c.rawParams[param.Key] = param.Value
	}
}
