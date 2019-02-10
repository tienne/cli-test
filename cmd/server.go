package cmd

import (
	"github.com/mozzet/cli/cmd/server"
	"github.com/spf13/cobra"
)

var (
	subCommands []string = []string{"list"}
	serverCmd            = &cobra.Command{
		Use:       "server",
		ValidArgs: subCommands,
	}
)

func init() {
	// sub command
	serverCmd.AddCommand(server.ListCmd)

	rootCmd.AddCommand(serverCmd)
}
