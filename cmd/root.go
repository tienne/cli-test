package cmd

import (
	"fmt"
	"github.com/mozzet/cli/config"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "mozzet",
	Short: "mozzet cli",
	Long:  "mozzet command line interface",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)
}
