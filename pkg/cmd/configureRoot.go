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

var rootConfigureCmd = &cobra.Command{
	Use:   "root",
	Short: "Show configuration file",
	Long:  `Show configuration file`,
	Run: func(cmd *cobra.Command, args []string) {
		configureRootRunner()
	},
}

func configureRootRunner() {
	config.EnsureConfiguration()
	globalConfig := config.FindGlobalconfig()
	if FlagConfigureRaw == true {
		if len(globalConfig) > 0 {
			fmt.Println(globalConfig)
		}
		return
	}
	if len(globalConfig) == 0 {
		terminal.Print(&terminal.TerminalOutput{
			Messages:    []string{"Global config file is not initialized"},
			MessageType: terminal.Error,
			Output:      "(use \"nitroci configure <tool> --profile <profile>\" to initialize a specific tool)",
		})
	} else {
		tItems := terminal.TerminalItemsOutput{
			Messages:    []string{"Configure the required tool"},
			Suggestions: []string{"(use \"nitroci configure <tool> --profile <profile>\" to initialize a specific tool)"},
			ItemsType:   terminal.Info,
		}
		currentConfigureTxt := fmt.Sprintf("Your curent configure is set to %v", globalConfig)
		terminal.Print(&terminal.TerminalOutput{
			Messages:    []string{"Global configuration has been initialized", currentConfigureTxt},
			ItemsOutput: []terminal.TerminalItemsOutput{tItems},
		})
	}
}

func init() {
	rootConfigureCmd.Flags().BoolVarP(&FlagConfigureRaw, "raw", "r", false, "get a raw result")
}
