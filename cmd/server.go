package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "Create a new command to start a grpc and http based application",

	Run: func(cmd *cobra.Command, args []string) {
		//scaffold.NewGokitServerCmd(newProject)
		fmt.Fprintln(cmd.OutOrStdout(), `A gokit grpc/http server command has been added to your cli. Try it by running your app followed by "server"`)},
}

