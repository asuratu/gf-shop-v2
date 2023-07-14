package job

import (
	"context"

	"shop/internal/consts"

	"github.com/hibiken/asynq"
)

type CronJob struct {
	ctx context.Context
}

func NewCronJob(ctx context.Context) *CronJob {
	return &CronJob{
		ctx: ctx,
	}
}

// Register register jobsvc
func (l *CronJob) Register() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	// test queue jobsvc
	mux.Handle(consts.JobTestQueue, NewTestQueueHandler())

	// defer jobsvc
	// mux.Handle(consts.RegisterUserJob, NewRegisterHandler(l.svcCtx))

	return mux
}
