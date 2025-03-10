package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CnTeng/todoist-api-go/sync/v9"
)

func main() {
	var token string
	fmt.Print("Enter your token: ")
	if _, err := fmt.Scanln(&token); err != nil {
		panic(err)
	}

	c := sync.NewClient(http.DefaultClient, token)

	r, err := c.AddItem(context.Background(), &sync.ItemAddArgs{Content: "Test Item"})
	if err != nil {
		panic(err)
	}

	rj, _ := json.MarshalIndent(r, "", "  ")

	fmt.Printf("%v\n", string(rj))

	rr, err := c.Sync(context.Background())
	if err != nil {
		panic(err)
	}

	rrj, _ := json.MarshalIndent(rr, "", "  ")
	fmt.Printf("%v\n", string(rrj))
}
