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

	pkgCRegistries "github.com/nitroci/nitroci-core/pkg/core/registries"
	pkgCTerminal "github.com/nitroci/nitroci-core/pkg/core/terminal"

	"github.com/spf13/cobra"
)

var (
	pluginsShow, pluginsRaw bool
)

var pluginsCmd = &cobra.Command{
	Use:   "plugins",
	Short: "Plugins managament",
	Long:  `Plugins management`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !runtimeContext.HasWorkspaces() {
			return errors.New("workspace is not initialized")
		}
		return pluginsRunner()
	},
}

func pluginsRunner() error {
	if !pluginsShow {
		return nil
	}
	workspace, err := runtimeContext.GetCurrentWorkspace()
	if err != nil {
		return err
	}
	workspaceModel, _ := workspace.CreateWorkspaceInstance()
	currentWorkspaceTxt := fmt.Sprintf("Your curent workspace is set to %v", pkgCTerminal.ConvertToCyanColor(workspace.WorkspacePath))
	if len(workspaceModel.Workspace.Plugins) == 0 {
		if !pluginsRaw {
			pkgCTerminal.Print(&pkgCTerminal.TerminalOutput{
				Messages: []string{"On workspace", currentWorkspaceTxt},
				Output:   "Workspace doesn't include any plugin.",
			})
		}
	} else {
		if pluginsRaw {
			for _, m := range workspaceModel.Workspace.Plugins {
				pkgCTerminal.Println(pkgCRegistries.GetPackageName(m.Name, m.Version))
			}
		} else {
			commands := make([]string, len(workspaceModel.Commands))
			for i, m := range workspaceModel.Workspace.Plugins {
				commands[i] = pkgCRegistries.GetPackageName(m.Name, m.Version)
			}
			tItems1 := pkgCTerminal.TerminalItemsOutput{
				Messages: []string{"Run one of the following commands:"},
				Suggestions: []string{
					"(use \"nitroci install \" to install plugins using the default workspace)",
					"(use \"nitroci install -w 1 ...\" to install plugins using a specific workspace)"},
				ItemsType: pkgCTerminal.Info,
				Items:     commands,
			}
			pkgCTerminal.Print(&pkgCTerminal.TerminalOutput{
				Messages:    []string{"On workspace", currentWorkspaceTxt},
				ItemsOutput: []pkgCTerminal.TerminalItemsOutput{tItems1},
			})
		}
	}
	return nil
}

func init() {
	rootCmd.AddCommand(pluginsCmd)
	pluginsCmd.Flags().BoolVarP(&pluginsShow, "show", "s", false, "show configurations")
	pluginsCmd.Flags().BoolVarP(&pluginsRaw, "raw", "r", false, "output raw configurations")
}
