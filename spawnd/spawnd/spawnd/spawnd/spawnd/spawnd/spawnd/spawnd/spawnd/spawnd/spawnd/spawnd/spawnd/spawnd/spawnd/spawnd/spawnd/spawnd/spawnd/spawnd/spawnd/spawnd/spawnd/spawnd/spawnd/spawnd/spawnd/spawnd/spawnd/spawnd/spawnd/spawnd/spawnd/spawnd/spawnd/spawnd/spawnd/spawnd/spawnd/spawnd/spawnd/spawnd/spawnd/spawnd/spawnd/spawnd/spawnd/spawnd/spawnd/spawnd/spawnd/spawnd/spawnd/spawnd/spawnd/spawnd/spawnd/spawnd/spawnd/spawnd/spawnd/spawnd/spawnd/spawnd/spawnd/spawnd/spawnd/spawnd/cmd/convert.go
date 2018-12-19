package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var(
	directory string
	pkg string
)
func init() {
	rootCmd.AddCommand(convertCmd)
	convertCmd.Flags().StringVarP(&directory, "dir", "d", ".", "directory to create swagger.pb.go")
	convertCmd.Flags().StringVarP(&pkg, "package", "p", "", "package of the directory that swagger.pb.go will reside in")
}
// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fs, _ := ioutil.ReadDir(directory)
		out, _ := os.Create("swagger.pb.go")
		out.Write([]byte("package "+pkg+"\n\nconst (\n"))
		for _, f := range fs {
			if strings.HasSuffix(f.Name(), ".json") {
				name := strings.TrimPrefix(f.Name(), "service.")
				out.Write([]byte(strings.TrimSuffix(name, ".json") + " = `"))
				f, _ := os.Open(f.Name())
				io.Copy(out, f)
				out.Write([]byte("`\n"))
			}
		}
		out.Write([]byte(")\n"))
	},
}