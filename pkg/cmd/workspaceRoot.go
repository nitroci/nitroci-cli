/*
Copyright 2021 The NitroCI Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"nitroci/pkg/internal/config"
	"nitroci/pkg/internal/io/terminal"

	"github.com/spf13/cobra"
)

var rootWorkspaceCmd = &cobra.Command{
	Use:   "root",
	Short: "List workspace files",
	Long:  `List workspace files`,
	Run: func(cmd *cobra.Command, args []string) {
		workspaceRootRunner()
	},
}

func workspaceRootRunner() {
	projectFiles := config.FindProjectFiles()
	if FlagWorkspaceRaw == true {
		if len(*projectFiles) > 0 {
			fmt.Println((*projectFiles)[0])
		}
		return
	}
	if len(*projectFiles) == 0 {
		terminal.Print(&terminal.TerminalOutput{
			Messages:    []string{"Workspace is not initialized"},
			MessageType: terminal.Error,
			Output:      "use \"nitroci workspace init\" to initialize the workspace",
		})
	} else {
		files := []string{}
		for i, m := range *projectFiles {
			files = append(files, fmt.Sprintf("%v %v", i+1, m))
		}
		tItems := terminal.TerminalItemsOutput{
			Messages:    []string{"Intialized workspaces:"},
			Suggestions: []string{"(use \"nitroci <commamnd> -w <workspace-depth>...\" to switch workspace)"},
			ItemsType:   terminal.Info,
			Items:       files,
		}
		currentWorkspaceTxt := fmt.Sprintf("Your curent workspace is set to %v", (*projectFiles)[0])
		terminal.Print(&terminal.TerminalOutput{
			Messages:    []string{"Workspace has been initialized", currentWorkspaceTxt},
			ItemsOutput: []terminal.TerminalItemsOutput{tItems},
		})
	}
}

func init() {
	rootWorkspaceCmd.Flags().BoolVarP(&FlagWorkspaceRaw, "raw", "r", false, "get a raw result")
}
