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
	"os"
	"os/exec"
	"strings"

	"github.com/nitroci/nitroci-core/pkg/core/terminal"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run workspace commands",
	Long:  `Run workspace commands`,
	Run: func(cmd *cobra.Command, args []string) {
		runner(args)
	},
}

func runner(args []string) {
	if !runtimeContext.HasWorkspaces() {
		workspaceRunner()
		return
	}
	workspace, _ := runtimeContext.GetCurrentWorkspace()
	workspaceModel, _ := workspace.CreateWorkspaceInstance()
	currentWorkspaceTxt := fmt.Sprintf("Your curent workspace is set to %v", workspace.WorkspaceFile)
	if len(workspaceModel.Commands) == 0 {
		terminal.Print(&terminal.TerminalOutput{
			Messages: []string{"On workspace", currentWorkspaceTxt},
			Output:   "Workspace doesn't implement commands to be executed.",
		})
		return
	}
	if len(args) != 1 || len((args)[0]) == 0 {
		commands := make([]string, len(workspaceModel.Commands))
		for i, m := range workspaceModel.Commands {
			commands[i] = strings.ToLower(m.Name) + ": " + strings.ToLower(m.Description)
		}
		tItems1 := terminal.TerminalItemsOutput{
			Messages:    []string{"Run one of the following commands:"},
			Suggestions: []string{"(use \"nitroci run <command>...\" to run a command using the default workspace)", "(use \"nitroci run <command> -w 1 ...\" to run a command using a specific workspace)"},
			ItemsType:   terminal.Info,
			Items:       commands,
		}
		terminal.Print(&terminal.TerminalOutput{
			Messages:    []string{"On workspace", currentWorkspaceTxt},
			ItemsOutput: []terminal.TerminalItemsOutput{tItems1},
		})
		return
	}
	for _, m := range workspaceModel.Commands {
		if m.Name != (args)[0] {
			continue
		}
		for _, step := range m.Steps {
			cwd := step.Cwd
			for _, script := range step.Scripts {
				cmd := exec.Command("sh", "-c", script)
				if len(cwd) > 0 {
					if _, err := os.Stat(cwd); errors.Is(err, os.ErrNotExist) {
						err := os.Mkdir(cwd, os.ModePerm)
						if err != nil {
							os.Exit(1)
						}
					}
					cmd.Dir = cwd
				}
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Run()
				if err != nil {
					fmt.Println("cmd.Run() failed with %s\n", err)
				}
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
}
