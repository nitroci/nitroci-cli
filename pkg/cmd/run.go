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
	"bufio"
	"errors"
	"fmt"
	"nitroci/pkg/internal/config"
	"nitroci/pkg/internal/io/terminal"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var FlagRunWorkspaceDepth int

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a workspace command",
	Long:  `Run a workspace command`,
	Run: func(cmd *cobra.Command, args []string) {
		runner(&args)
	},
}

func runner(args *[]string) {
	projectFile := config.FindCurrentProjectFile(FlagRunWorkspaceDepth)
	if len(*projectFile) == 0 {
		workspaceRootRunner()
		return
	}
	var config config.ProjectConfiguration
	config.LoadProject(projectFile)
	currentWorkspaceTxt := fmt.Sprintf("Your curent workspace is set to %v", *projectFile)
	if len(config.Commands) == 0 {
		terminal.Print(&terminal.TerminalOutput{
			Messages: []string{"On workspace", currentWorkspaceTxt},
			Output:   "Workspace doesn't implement commands to be executed.",
		})
		return
	}
	if len(*args) != 1 || len((*args)[0]) == 0 {
		commands := make([]string, len(config.Commands))
		for i, m := range config.Commands {
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
	for _, m := range config.Commands {
		if m.Name != (*args)[0] {
			continue
		}
		for i, step := range m.Steps {
			cwd := step.Cwd
			for j, script := range step.Scripts {
				step := ""
				if FlagVerbose {
					step = fmt.Sprintf("[Step %v of %v][%v]", i+1, len(m.Steps), j+1)
				}
				tAction := &terminal.TerminalActionOutput{
					StepId: step,
					Step:   script,
				}
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
				stdout, err := cmd.StdoutPipe()
				cmd.Stderr = cmd.Stdout
				if err != nil {
					fmt.Println(err)
				}
				err = cmd.Start()
				if err != nil {
					fmt.Println(err)
				}
				scanner := bufio.NewScanner(stdout)
				for scanner.Scan() {
					m := scanner.Text()
					tAction.Outputs = append(tAction.Outputs, m)
					terminal.PrintActions(tAction)
				}
				cmd.Wait()
			}
		}
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.PersistentFlags().IntVarP(&FlagRunWorkspaceDepth, "workspace", "w", 0, "set current workspace")
}
