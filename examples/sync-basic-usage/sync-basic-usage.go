package main

import (
	"context"
	"net/http"
	"os"

	"github.com/CnTeng/todoist-api-go/sync"
	"github.com/CnTeng/todoist-api-go/todoist"
)

func main() {
	token := os.Getenv("API_TOKEN")
	cli := todoist.NewClient(http.DefaultClient, token, todoist.DefaultHandler)
	taskSvc := todoist.NewTaskService(cli)

	_, err := cli.SyncWithAutoToken(context.Background(), true)
	if err != nil {
		panic(err)
	}

	args := &sync.TaskAddArgs{Content: "test"}
	_, err = taskSvc.AddTask(context.Background(), args)
	if err != nil {
		panic(err)
	}
}
