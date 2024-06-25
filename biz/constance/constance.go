package constance

import "time"

const (
	SimpleApiTimeOut = 1 * time.Second
	RetryTimes       = 5

	LagencyMin = 100
	LagencyMax = 1100
)

var GameInfoURL string = "http://127.0.0.1:%d/api/game_center/v1/meta/query/"
