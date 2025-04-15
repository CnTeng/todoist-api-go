package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/CnTeng/todoist-api-go/sync"
)

func main() {
	token := os.Getenv("API_TOKEN")
	c := sync.NewClient(http.DefaultClient, token, sync.DefaultHandler)

	r, err := c.GetTasksCompleted(context.Background(), &sync.TaskGetCompletedByDueDateParams{
		Since: time.Now().AddDate(-1, 0, 0),
		Until: time.Now(),
	})
	if err != nil {
		panic(err)
	}
	data, _ := json.MarshalIndent(r, "", "  ")
	fmt.Println(string(data))
}
