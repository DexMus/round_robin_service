// Code generated by hertz generator.

package round_robin_service

import (
	"code.byted.org/middleware/hertz/pkg/app"
	"github.com/DexMus/round_robin_service/middleware"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _apiMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _game_centerMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _v1Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _metaMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _queryMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _querygameinfoMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _round_queryMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _roundquerygameinfoMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.QueryGameInfoMw(),
	}
}
