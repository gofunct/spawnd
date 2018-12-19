package cmd

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"

	"github.com/gofunct/spawnd/cobra/generator"
)

func init() {
	rootCmd.AddCommand(cobraCmd)
}

var cobraCmd = &cobra.Command{
	Use: "cobra",
	Short: 	"generate a cobra client command",
	Run: func(cmd *cobra.Command, args []string) {
		// Begin by allocating a generator. The request and response structures are stored there
		// so we can do error handling easily - the response structure contains the field to
		// report failure.
		g := generator.New()

		data, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			g.Error(err, "reading input")
		}

		if err := proto.Unmarshal(data, g.Request); err != nil {
			g.Error(err, "parsing input proto")
		}

		if len(g.Request.FileToGenerate) == 0 {
			g.Fail("no files to generate")
		}

		g.CommandLineParameters(g.Request.GetParameter())

		// Create a wrapped version of the Descriptors and EnumDescriptors that
		// point to the file that defines them.
		g.WrapTypes()

		g.SetPackageNames()

		g.GenerateAllFiles()

		// Send back the results.
		data, err = proto.Marshal(g.Response)
		if err != nil {
			g.Error(err, "failed to marshal output proto")
		}
		_, err = os.Stdout.Write(data)
		if err != nil {
			g.Error(err, "failed to write output proto")
		}
	},
}