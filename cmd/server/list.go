package server

import (
	"github.com/mozzet/cli/aws"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use: "list",
	Run: runListCmd,
}

func runListCmd(cmd *cobra.Command, args []string) {
	aws.EC2List()
}
