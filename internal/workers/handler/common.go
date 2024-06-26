package handler

import (
	"context"
	"encoding/json"
	"swclabs/swipecore/internal/core/service/common"

	"swclabs/swipecore/pkg/lib/worker"

	"github.com/hibiken/asynq"
)

type Common struct {
	common.Task                       // embedded delay function here
	handler     common.ICommonService // create handler for services
}

func NewCommonConsume(_common common.ICommonService) *Common {
	return &Common{
		handler: _common,
	}
}

func (common *Common) HandleHealthCheck() (taskName string, fn worker.HandleFunc) {
	// get task name from delay function
	taskName = worker.GetTaskName(common.WorkerCheckResult)
	// implement handler function base on delay function
	return taskName, func(_ context.Context, task *asynq.Task) error {
		var num int64
		if err := json.Unmarshal(task.Payload(), &num); err != nil {
			return err
		}
		if err := common.handler.WorkerCheck(context.Background(), num); err != nil {
			return err
		}
		_, err := task.ResultWriter().Write([]byte("HandleHealthCheck: success"))
		return err
	}
}
