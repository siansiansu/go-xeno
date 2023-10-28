package main

import (
	"context"

	"github.com/siansiansu/go-xeno"
)

func main() {
	var ctx = context.Background()

	client, err := xeno.NewClient()
	if err != nil {
		panic(err)
	}

	r, err := client.Get(ctx, "Taiwan blue magpie", xeno.Page(1), xeno.NumPages(1))
	if err != nil {
		panic(err)
	}

	for _, e := range r.Recordings {
		xeno.DownloadFile(e.FileName, e.File)
	}
}
