package handler

import (
	"context"
	"encoding/json"

	"swclabs/swipe-api/internal/core/domain"
	"swclabs/swipe-api/internal/core/service"
	"swclabs/swipe-api/internal/helper/tasks"
	"swclabs/swipe-api/pkg/tools/worker"

	"github.com/hibiken/asynq"
)

type AccountManagement struct {
	tasks.AccountManagement
	handler domain.IAccountManagementService
}

func NewAccountManagement() *AccountManagement {
	return &AccountManagement{
		handler: service.NewAccountManagement(),
	}
}

func (account *AccountManagement) HandleSignUp() (string, worker.HandleFunc) {
	return worker.GetTaskName(account.DelaySignUp),
		func(_ context.Context, task *asynq.Task) error {
			var data domain.SignUpRequest
			if err := json.Unmarshal(task.Payload(), &data); err != nil {
				return err
			}
			return account.handler.SignUp(context.Background(), &data)
		}
}

func (account *AccountManagement) HandleOAuth2SaveUser() (string, worker.HandleFunc) {
	return worker.GetTaskName(account.DelayOAuth2SaveUser),
		func(_ context.Context, task *asynq.Task) error {
			var data domain.OAuth2SaveUser
			if err := json.Unmarshal(task.Payload(), &data); err != nil {
				return err
			}
			return account.handler.OAuth2SaveUser(context.Background(), &data)
		}
}

func (account *AccountManagement) HandleUpdateUserInfo() (string, worker.HandleFunc) {
	return worker.GetTaskName(account.DelayUpdateUserInfo),
		func(_ context.Context, task *asynq.Task) error {
			var data domain.UserUpdate
			if err := json.Unmarshal(task.Payload(), &data); err != nil {
				return err
			}
			return account.handler.UpdateUserInfo(context.Background(), &data)
		}
}