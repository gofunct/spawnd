package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"fmt"
)

var (
	service string
)

func init() {
	rootCmd.Flags().StringVar(&service, "service", "", "The protobuf message used for this configuration")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "example",
	Short: "A brief description of your application",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
