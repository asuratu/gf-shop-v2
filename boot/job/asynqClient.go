package job

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/hibiken/asynq"
)

var AsynqClient *asynq.Client

func NewAsynqClient(ctx context.Context) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr:     g.Cfg().MustGet(ctx, "redis.default.address").String(),
		Password: g.Cfg().MustGet(ctx, "redis.default.pass").String(),
		DB:       g.Cfg().MustGet(ctx, "redis.default.jobDb").Int(),
	})
}
