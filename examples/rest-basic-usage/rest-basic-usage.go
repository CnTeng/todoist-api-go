package main

import (
	"context"
	"net/http"
	"os"

	"github.com/CnTeng/todoist-api-go/rest"
	"github.com/CnTeng/todoist-api-go/todoist"
)

func main() {
	token := os.Getenv("API_TOKEN")
	cli := todoist.NewClient(http.DefaultClient, token, todoist.DefaultHandler)

	req := &rest.TaskQuickAddRequest{Text: "test"}
	_, err := cli.AddTaskQuick(context.Background(), req)
	if err != nil {
		panic(err)
	}
}
