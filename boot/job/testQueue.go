package job

import (
	"context"
	"encoding/json"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/hibiken/asynq"
)

// TestQueueHandler test queue jobsvc
type TestQueueHandler struct{}

func NewTestQueueHandler() *TestQueueHandler {
	return &TestQueueHandler{}
}

// ProcessTask every one minute exec : if return err != nil , asynq will retry
func (l *TestQueueHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	// 有效载荷
	var p TestQueuePayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return gerror.New("unmarshal payload")
	}
	g.Log().Infof(ctx, "-----------------> handle test queue jobsvc , payload: %+v", p)
	g.Dump(p)
	g.Dump(gtime.Now().String())
	return nil
}
