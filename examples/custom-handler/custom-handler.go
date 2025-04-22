package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/CnTeng/todoist-api-go/rest"
	"github.com/CnTeng/todoist-api-go/sync"
	"github.com/CnTeng/todoist-api-go/todoist"
)

type customHandler struct{}

func (h *customHandler) SyncToken(ctx context.Context) (*string, error) {
	// Full sync every time.
	st := sync.DefaultSyncToken
	return &st, nil
}

func (h *customHandler) ResourceTypes(ctx context.Context) (*sync.ResourceTypes, error) {
	// Only sync tasks.
	return &sync.ResourceTypes{sync.Tasks}, nil
}

func (h *customHandler) HandleResponse(ctx context.Context, resp any) error {
	// Print the tasks.
	switch r := resp.(type) {
	case *sync.SyncResponse:
		for _, task := range r.Tasks {
			fmt.Printf("Task ID: %s, Content: %s\n", task.ID, task.Content)
		}
	case *rest.TaskGetCompletedResponse:
		for _, task := range r.Tasks {
			fmt.Printf("Completed Task ID: %s, Content: %s\n", task.ID, task.Content)
		}
	}

	return nil
}

func main() {
	token := os.Getenv("API_TOKEN")

	c := todoist.NewClient(http.DefaultClient, token, &customHandler{})

	_, err := c.SyncWithAutoToken(context.Background(), true)
	if err != nil {
		panic(err)
	}
}
