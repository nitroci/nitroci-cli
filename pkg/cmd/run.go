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

	pkgCTerminal "github.com/nitroci/nitroci-core/pkg/core/terminal"

	"github.com/spf13/cobra"
)

var (
	runShow, runRaw bool
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run workspace commands",
	Long:  `Run workspace commands`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !runtimeContext.HasWorkspaces() {
			return errors.New("workspace is not initialized")
		}
		return runner(args)
	},
}

func runner(args []string) error {
	workspace, err := runtimeContext.GetCurrentWorkspace()
	if err != nil {
		return err
	}
	workspaceModel, _ := workspace.CreateWorkspaceInstance()
	if runShow {
		currentWorkspaceTxt := fmt.Sprintf("Your curent workspace is set to %v", pkgCTerminal.ConvertToCyanColor(workspace.WorkspacePath))
		if len(workspaceModel.Commands) == 0 {
			if !runRaw {
				pkgCTerminal.Print(&pkgCTerminal.TerminalOutput{
					Messages: []string{"On workspace", currentWorkspaceTxt},
					Output:   "Workspace doesn't implement commands to be executed.",
				})
			}
			return nil
		}
		if len(args) != 1 || len(args[0]) == 0 {
			if runRaw {
				for _, m := range workspaceModel.Commands {
					pkgCTerminal.Println(strings.ToLower(m.Name) + ": " + strings.ToLower(m.Description))
				}
			} else {
				commands := make([]string, len(workspaceModel.Commands))
				for i, m := range workspaceModel.Commands {
					commands[i] = strings.ToLower(m.Name) + ": " + strings.ToLower(m.Description)
				}
				tItems1 := pkgCTerminal.TerminalItemsOutput{
					Messages:    []string{"Run one of the following commands:"},
					Suggestions: []string{"(use \"nitroci run <command>...\" to run a command using the default workspace)", "(use \"nitroci run <command> -w 1 ...\" to run a command using a specific workspace)"},
					ItemsType:   pkgCTerminal.Info,
					Items:       commands,
				}
				pkgCTerminal.Print(&pkgCTerminal.TerminalOutput{
					Messages:    []string{"On workspace", currentWorkspaceTxt},
					ItemsOutput: []pkgCTerminal.TerminalItemsOutput{tItems1},
				})
			}
		}
		return nil
	} else if len(args) == 1 && len(args[0]) != 0 {
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
								return err
							}
						}
						cmd.Dir = cwd
					}
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					err := cmd.Run()
					if err != nil {
						return err
					}
				}
			}
		}
		return nil
	}
	return errors.New("invalid command")
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().BoolVarP(&runShow, "show", "s", false, "show configurations")
	runCmd.Flags().BoolVarP(&runRaw, "raw", "r", false, "output raw configurations")
}
