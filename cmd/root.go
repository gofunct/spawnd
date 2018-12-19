package cmd

import (
	"github.com/spf13/cobra"
	"github.com/gofunct/spawnd/scaffold"
)

var (
	service, packageName, parentName, gopath, templatePath string
	newProject = scaffold.NewScaffoldFromCurrentPath()
)

func init() {
	rootCmd.Flags().StringVar(&service, "service", "", "The protobuf message used for this configuration")
}

var (
	rootCmd = &cobra.Command{
		Use:   "spawnd",
		Short: "spawnd is a utility for easily creating highly configurable golang microservices",
	}
)

// Execute executes the root command.
func Execute() {
	rootCmd.Execute()
}

