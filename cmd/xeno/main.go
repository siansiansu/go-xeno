package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/schollz/progressbar/v3"
	"github.com/siansiansu/go-xeno"
	"github.com/spf13/cobra"
)

var (
	search     string
	folder     string
	maxResults int
	version    = "v1.0.0"
)

var rootCmd = &cobra.Command{
	Use:     "xeno",
	Version: version,
	Short:   "go-xeno is a CLI to download audio recordings",
	Long: `
xeno

download audio recordings from xeno-canto

- download to current folder
  xeno "Taiwan blue magpie"

- download to specific folder
  xeno "Taiwan blue magpie" --folder "download"
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
			os.Exit(0)
		}

		search = args[0]

		if err := cobra.MaximumNArgs(1)(cmd, args); err != nil {
			panic(err)
		}

		var ctx = context.Background()

		client, err := xeno.NewClient()
		if err != nil {
			panic(err)
		}

		r, err := client.Get(ctx, search)
		if err != nil {
			panic(err)
		}

		totalRecords, err := strconv.ParseInt(r.NumRecordings, 10, 32)
		if err != nil {
			panic(err)
		}

		if folder != "" {
			if _, err := os.Stat(folder); errors.Is(err, os.ErrNotExist) {
				err := os.Mkdir(folder, os.ModePerm)
				if err != nil {
					log.Println(err)
				}
			}
		}

		if maxResults != 0 && int64(maxResults) < totalRecords {
			totalRecords = int64(maxResults)
		}

		bar := progressbar.Default(totalRecords)

		count := 0
		downloadList := map[string]string{}

		for i := 1; i < r.NumPages; i++ {
			if count == maxResults {
				break
			}
			s, err := client.Get(ctx, search, xeno.Page(i))
			if err != nil {
				panic(err)
			}
			for _, j := range s.Recordings {
				count++
				downloadList[j.FileName] = j.File
				if count == maxResults {
					break
				}
			}
		}

		if len(downloadList) == 0 {
			fmt.Println("No recordings to download!")
			return
		}

		var wg sync.WaitGroup

		for k, v := range downloadList {
			wg.Add(1)

			k, v := k, v
			go func() {
				defer wg.Done()
				bar.Add(1)
				xeno.DownloadFile(filepath.Join(folder, k), v)
			}()
		}
		wg.Wait()
		fmt.Println("\nDownload Completed!")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&folder, "folder", "f", "", "specify the download folder")
	rootCmd.PersistentFlags().IntVarP(&maxResults, "max-results", "m", 0, "specify the maximum downloads")
}

func main() {
	Execute()
}
