// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/gofunct/grpcgen/project"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "Create a new command to start a grpc and http based application",

	Run: func(cmd *cobra.Command, args []string) {
		project.NewGokitServerCmd(newProject)
		fmt.Fprintln(cmd.OutOrStdout(), `A gokit grpc/http server command has been added to your cli. Try it by running your app followed by "server"`)},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
