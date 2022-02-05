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
	"errors"
	"fmt"

	pkgCTerminal "github.com/nitroci/nitroci-core/pkg/core/terminal"

	"github.com/spf13/cobra"
)

var (
	workspaceShow, workspaceRaw bool
)

var workspacesCmd = &cobra.Command{
	Use:   "workspaces",
	Short: "List and interact with configured workspaces",
	Long:  `List and interact with configured workspaces`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !runtimeContext.HasWorkspaces() {
			return errors.New("workspace is not initialized")
		}
		return workspaceRunner()
	},
}

func workspaceRunner() error {
	if !workspaceShow {
		return nil
	}
	if workspaceRaw {
		workspace, err := runtimeContext.GetCurrentWorkspace()
		if err != nil {
			return err
		}
		pkgCTerminal.Println(workspace.WorkspacePath)
		return nil
	}
	files := []string{}
	workspaces, err := runtimeContext.GetWorkspaces()
	if err != nil {
		return err
	}
	if len(workspaces) == 0 {
		return errors.New("please initialize a workspace.")
	}
	for i, w := range workspaces {
		files = append(files, fmt.Sprintf("%v %v", i+1, w.WorkspacePath))
	}
	tItems := pkgCTerminal.TerminalItemsOutput{
		Messages:    []string{"Intialized workspaces:"},
		Suggestions: []string{"(use \"nitroci <commamnd> -w <workspace-depth>...\" to switch workspace)"},
		ItemsType:   pkgCTerminal.Info,
		Items:       files,
	}
	workspace, err := runtimeContext.GetCurrentWorkspace()
	if err != nil {
		return err
	}
	currentWorkspaceTxt := fmt.Sprintf("Your curent workspace is set to %v", pkgCTerminal.ConvertToCyanColor(workspace.WorkspacePath))
	pkgCTerminal.Print(&pkgCTerminal.TerminalOutput{
		Messages:    []string{"Workspace has been initialized", currentWorkspaceTxt},
		ItemsOutput: []pkgCTerminal.TerminalItemsOutput{tItems},
	})
	return nil
}

func init() {
	rootCmd.AddCommand(workspacesCmd)
	workspacesCmd.Flags().BoolVarP(&workspaceShow, "show", "s", false, "show configurations")
	workspacesCmd.Flags().BoolVarP(&workspaceRaw, "raw", "r", false, "output raw configurations")
}
