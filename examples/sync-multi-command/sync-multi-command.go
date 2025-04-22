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

	projCmd := sync.NewCommand(&sync.ProjectAddArgs{Name: "test"})

	tempID := projCmd.TempID.String()
	taskCmd := sync.NewCommand(&sync.TaskAddArgs{Content: "test", ProjectID: &tempID})

	_, err := cli.ExecuteCommands(context.Background(), sync.Commands{projCmd, taskCmd})
	if err != nil {
		panic(err)
	}
}
