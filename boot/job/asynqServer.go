package job

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/hibiken/asynq"
)

var AsynqServer *asynq.Server

func NewAsynqServer(ctx context.Context) *asynq.Server {

	return asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     g.Cfg().MustGet(ctx, "redis.default.address").String(),
			Password: g.Cfg().MustGet(ctx, "redis.default.pass").String(),
			DB:       g.Cfg().MustGet(ctx, "redis.default.jobDb").Int(),
		},
		asynq.Config{
			IsFailure: func(err error) bool {
				fmt.Printf("asynq server exec task IsFailure ======== >>>>>>>>>>>  err : %+v \n", err)
				return true
			},
			Concurrency: 20, // max concurrent process jobsvc task num
			// 关键队列中的任务将被处理 60% 的时间
			// 默认队列中的任务将被处理 30% 的时间
			// 低队列中的任务将被处理 10% 的时间
			Queues: map[string]int{
				"critical": 6,
				"default":  3,
				"low":      1,
			},
			StrictPriority: true, // strict mode! if true , will not exec low queue task
		},
	)
}
