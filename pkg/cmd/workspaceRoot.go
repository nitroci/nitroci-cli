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

	"github.com/nitroci/nitroci-core/pkg/core/io/terminal"

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
	if FlagWorkspaceRaw {
		if runtimeContext.HasWorkspaces() {
			workspace, _ := runtimeContext.GetWorkspace(0)
			fmt.Println(workspace.WorkspaceFileName)
		}
		return
	}
	if !runtimeContext.HasWorkspaces() {
		terminal.Print(&terminal.TerminalOutput{
			Messages:    []string{"Workspace is not initialized"},
			MessageType: terminal.Error,
			Output:      "use \"nitroci workspace init\" to initialize the workspace",
		})
	} else {
		files := []string{}
		workspaces, _ := runtimeContext.GetWorkspaces()
		for i, w := range workspaces {
			files = append(files, fmt.Sprintf("%v %v", i+1, w.WorkspaceFileName))
		}
		tItems := terminal.TerminalItemsOutput{
			Messages:    []string{"Intialized workspaces:"},
			Suggestions: []string{"(use \"nitroci <commamnd> -w <workspace-depth>...\" to switch workspace)"},
			ItemsType:   terminal.Info,
			Items:       files,
		}
		workspace, _ := runtimeContext.GetWorkspace(0)
		currentWorkspaceTxt := fmt.Sprintf("Your curent workspace is set to %v", workspace.WorkspaceFileName)
		terminal.Print(&terminal.TerminalOutput{
			Messages:    []string{"Workspace has been initialized", currentWorkspaceTxt},
			ItemsOutput: []terminal.TerminalItemsOutput{tItems},
		})
	}
}

func init() {
	rootWorkspaceCmd.Flags().BoolVarP(&FlagWorkspaceRaw, "raw", "r", false, "get a raw result")
}
