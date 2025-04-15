package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/CnTeng/todoist-api-go/rest"
	"github.com/CnTeng/todoist-api-go/todoist"
)

func main() {
	token := os.Getenv("API_TOKEN")
	c := todoist.NewClient(http.DefaultClient, token, todoist.DefaultHandler)

	r, err := c.GetTasksCompletedByDueDate(context.Background(), &rest.TaskGetCompletedByDueDateParams{
		Since: time.Now().AddDate(-1, 0, 0),
		Until: time.Now(),
	})
	if err != nil {
		panic(err)
	}
	data, _ := json.MarshalIndent(r, "", "  ")
	fmt.Println(string(data))
}
