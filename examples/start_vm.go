package main

import (
	"os"
	"time"

	apiclient "github.com/Yuyz0112/cloudtower-go-sdk/client"
	"github.com/Yuyz0112/cloudtower-go-sdk/client/operations"
	"github.com/Yuyz0112/cloudtower-go-sdk/models"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/go-openapi/strfmt"
)

func strPtr(s string) *string {
	return &s
}

func waitTasksFinish(c *apiclient.AtTowerOperationAPI, taskIds []string) (*operations.GetTasksOK, error) {
	if len(taskIds) == 0 {
		return operations.NewGetTasksOK(), nil
	}

	tasksParams := operations.NewGetTasksParams()
	tasksParams.RequestBody = &models.GetTasksRequestBody{
		Where: &models.TaskWhereInput{
			IDIn: taskIds,
		},
	}

	for {
		tasksResp, err := c.Operations.GetTasks(tasksParams)
		if err != nil {
			return nil, err
		}

		allFinished := true
		for _, v := range tasksResp.Payload {
			if *v.Status != models.TaskStatusSUCCESSED && *v.Status != models.TaskStatusFAILED {
				allFinished = false
			}
		}

		if !allFinished {
			time.Sleep(5 * time.Second)
			continue
		}

		return tasksResp, nil
	}
}

func main() {
	transport := httptransport.New("localhost:8090", "/", []string{"http"})
	c := apiclient.New(transport, strfmt.Default)

	loginParams := operations.NewLoginParams()
	loginParams.RequestBody = &models.LoginInput{
		Username: strPtr(os.Getenv("USERNAME")),
		Password: strPtr(os.Getenv("PASSWORD")),
		Source:   models.NewUserSource(models.UserSourceLDAP),
	}
	loginResp, err := c.Operations.Login(loginParams)
	if err != nil {
		panic(err.Error())
	}

	bearerTokenAuth := httptransport.BearerToken(*loginResp.Payload.Data.Token)
	transport.DefaultAuthentication = bearerTokenAuth
	c = apiclient.New(transport, strfmt.Default)

	startVmsParams := operations.NewStartVMParams()
	startVmsParams.RequestBody = &models.VMStartParams{
		Where: &models.VMWhereInput{
			Status: models.VMStatusSTOPPED,
			Name:   strPtr("sdk-test"),
		},
	}
	startVmsResp, err := c.Operations.StartVM(startVmsParams)
	if err != nil {
		panic(err.Error())
	}

	taskIds := make([]string, 0)
	for _, v := range startVmsResp.Payload {
		taskIds = append(taskIds, *v.TaskID)
	}

	tasksResp, err := waitTasksFinish(c, taskIds)
	if err != nil {
		panic(err.Error())
	}

	for _, v := range tasksResp.Payload {
		println(*v.Status, *v.Description)
	}
}
