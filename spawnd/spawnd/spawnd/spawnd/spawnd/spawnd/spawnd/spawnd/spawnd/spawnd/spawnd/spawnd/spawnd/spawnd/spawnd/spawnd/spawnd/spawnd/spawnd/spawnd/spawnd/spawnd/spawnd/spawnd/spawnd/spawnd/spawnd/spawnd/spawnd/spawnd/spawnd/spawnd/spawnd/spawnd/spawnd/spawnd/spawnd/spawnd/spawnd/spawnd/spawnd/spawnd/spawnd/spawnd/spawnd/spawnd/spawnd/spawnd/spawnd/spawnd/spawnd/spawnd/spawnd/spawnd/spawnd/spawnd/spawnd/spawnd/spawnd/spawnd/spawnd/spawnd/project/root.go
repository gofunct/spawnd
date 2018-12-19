package project

import (
	"github.com/gofunct/grpcgen/logging"
	"github.com/gofunct/grpcgen/project/utils"
	"path"
	"path/filepath"
)

func (p *Project) CreateRootCmdFile() {
	template := `package cmd

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
	Use:   "{{.appName}}",
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
`

	data := make(map[string]interface{})
	data["viper"] = true
	data["appName"] = path.Base(p.GetName())

	rootCmdScript, err := utils.ExecTemplate(template, data)
	logging.IfErr("failed to execute template", err)
	err = utils.WriteStringToFile(filepath.Join(p.GetCmd(), "root.go"), rootCmdScript)
	logging.IfErr("failed to write file", err)

}
