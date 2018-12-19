package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:     "init [name]",
	Aliases: []string{"initialize", "initialise", "create"},
	Short:   "Initialize a GrpcGen Application",
	Long: `Initialize (grpcgen init) will create a new application
with the appropriate structure for a grpcgen-based CLI application.`,

	Run: func(cmd *cobra.Command, args []string)  {
		//project.InitializeProject(newProject)
		fmt.Fprintln(cmd.OutOrStdout(), `Your grpcgen application is ready at
`+newProject.GetAbsPath()+`

Give it a try by going there and running `+"`go run main.go`."+`
Add commands to it by running `+"`cobra add [cmdname]`.")

	},
}
