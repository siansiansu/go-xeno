package main

import (
	"context"
	"fmt"

	"github.com/siansiansu/go-xeno/xeno"
)

var (
	count          = 0
	downloadFolder = "./download"
	species        = "Eurasian Tree Sparrow"
)

func main() {
	var ctx = context.Background()

	client, err := xeno.NewClient()
	if err != nil {
		panic(err)
	}

	r, err := client.Get(ctx, species)
	if err != nil {
		panic(err)
	}

	for i := 1; i < r.NumPages; i++ {
		s, err := client.Get(ctx, species, xeno.Page(i))
		if err != nil {
			panic(err)
		}

		for _, j := range s.Recordings {
			fmt.Printf("Current Recording: %d Total Recordings %s, Current Page: %d, Total Pages: %d\n", count, s.NumRecordings, s.Page, s.NumPages)
			count++
			xeno.DownloadFile(downloadFolder+"/"+j.FileName, j.File)
		}
	}
}
