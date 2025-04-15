package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/CnTeng/todoist-api-go/todoist"
)

func main() {
	token := os.Getenv("API_TOKEN")
	c := todoist.NewClient(http.DefaultClient, token, todoist.DefaultHandler)

	r, err := c.SyncWithAutoToken(context.Background(), true)
	if err != nil {
		panic(err)
	}
	data, _ := json.MarshalIndent(r, "", "  ")
	fmt.Println(string(data))
}
