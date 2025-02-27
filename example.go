package main

import (
	"fmt"

	"github.com/CnTeng/todoist-api-go/sync/v9"
)

func main() {
	var token string
	fmt.Print("Enter your token: ")
	if _, err := fmt.Scanln(&token); err != nil {
		panic(err)
	}

	sc := sync.NewSyncClient(token)

	rp := &sync.ReadParams{
		SyncToken:     "*",
		ResourceTypes: &sync.ResourceTypes{sync.All},
	}

	r, err := sc.Read(rp)
	if err != nil {
		panic(err)
	}

	print(r.Items)

	for _, item := range r.Items {
		println(item.Content)
		if item.ParentID != nil {
			println(item.ParentID)
		}

		if item.Due.Timezone != nil {
			println(item.Due.Timezone.String())
			println(item.Due.Date.String())
			println(item.Due.Date.In(item.Due.Timezone).String())
			println(item.Due.Date.Location().String())
		}
	}
}
