/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	search    string
	folder    string
	maxResult int
)

var rootCmd = &cobra.Command{
	Use:   "go-xeno",
	Short: "go-xeno is a CLI to download audio recordings",
	Long:  `go-xeno is a CLI to download audio recordings`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&folder, "folder", "", "specify the download folder")
}
