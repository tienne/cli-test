package cmd

import (
	"github.com/spf13/cobra"
)

var sshCmd = &cobra.Command{
	Use: "ssh",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
}
