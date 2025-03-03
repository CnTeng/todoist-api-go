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

	param := &sync.CompletedGetParams{
		AnnotateItems: new(bool),
	}
	*param.AnnotateItems = true
	cr, err := sc.GetCompletedInfo(param)
	if err != nil {
		panic(err)
	}

	for _, item := range cr.Items {
		fmt.Printf("%+v\n", item)
		if item.ItemObject != nil {
			fmt.Println(item.ItemObject.Content)
			fmt.Println(item.MetaData)
		}
	}

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
	}
}
