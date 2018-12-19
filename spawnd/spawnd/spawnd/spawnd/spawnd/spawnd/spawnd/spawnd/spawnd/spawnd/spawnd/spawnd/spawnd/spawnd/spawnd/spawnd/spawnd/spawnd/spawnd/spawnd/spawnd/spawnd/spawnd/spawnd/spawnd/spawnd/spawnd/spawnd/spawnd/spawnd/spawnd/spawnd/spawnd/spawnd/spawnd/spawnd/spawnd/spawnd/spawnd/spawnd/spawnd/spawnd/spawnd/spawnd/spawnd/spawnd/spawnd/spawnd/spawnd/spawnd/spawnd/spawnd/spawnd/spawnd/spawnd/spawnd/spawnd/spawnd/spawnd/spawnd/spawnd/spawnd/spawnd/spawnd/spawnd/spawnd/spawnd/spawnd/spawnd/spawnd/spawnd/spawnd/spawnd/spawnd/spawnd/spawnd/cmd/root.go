package cmd

import (
	"github.com/gofunct/grpcgen/project"
	"github.com/spf13/cobra"
)

var (
	service, packageName, parentName, gopath, templatePath string
	newProject = project.NewProjectFromCurrentPath()
)

func init() {
	rootCmd.Flags().StringVar(&service, "service", "", "The protobuf message used for this configuration")
	rootCmd.AddCommand(initCmd)
}

var (
	rootCmd = &cobra.Command{
		Use:   "grpcgen",
		Short: "grpcgen is a utility for easily creating highly configurable golang microservices",
	}
)

// Execute executes the root command.
func Execute() {
	rootCmd.Execute()
}

