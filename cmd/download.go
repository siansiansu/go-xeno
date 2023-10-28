package cmd

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
	"github.com/siansiansu/go-xeno/xeno"
	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download audio recordings from xeno-canto",
	Long: `
xeno download

download audio recordings from xeno-canto

- download to current folder
  xeno download --search "Taiwan blue magpie"

- download to specific folder
    xeno download --search "Taiwan blue magpie" --folder "download"
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

		if maxResult != 0 && int64(maxResult) < totalRecords {
			totalRecords = int64(maxResult)
		}

		bar := progressbar.Default(totalRecords)

		var wg sync.WaitGroup

		for i := 1; i < r.NumPages; i++ {
			s, err := client.Get(ctx, search, xeno.Page(i))
			if err != nil {
				panic(err)
			}

			for _, j := range s.Recordings {
				wg.Add(1)

				j := j

				go func() {
					defer wg.Done()
					bar.Add(1)
					xeno.DownloadFile(filepath.Join(folder, j.FileName), j.File)
				}()
			}
		}
		wg.Wait()
		fmt.Println("\nDownload Completed!")
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
