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

	pkgCContexts "github.com/nitroci/nitroci-core/pkg/core/contexts"
	pkgCTerminal "github.com/nitroci/nitroci-core/pkg/core/terminal"

	"github.com/spf13/cobra"
)

var (
	configureShow, configureRaw bool
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Show or modify configurations",
	Long:  `Show or modify configurations`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return configurationRunner()
	},
}

func configurationRunner() error {
	if !configureShow {
		return nil
	}
	globalConfig := runtimeContext.Cli.Settings[pkgCContexts.CFG_NAME_CONFIG_PATH]
	if configureRaw {
		if len(globalConfig) > 0 {
			pkgCTerminal.Println(globalConfig)
		}
		return nil
	}
	if len(globalConfig) == 0 {
		pkgCTerminal.Print(&pkgCTerminal.TerminalOutput{
			Messages:    []string{"Global config file is not initialized"},
			MessageType: pkgCTerminal.Error,
			Output:      "(use \"nitroci configure <tool> --profile <profile>\" to initialize a specific tool)",
		})
	} else {
		tItems := pkgCTerminal.TerminalItemsOutput{
			Messages:    []string{"Configure the required tool"},
			Suggestions: []string{"(use \"nitroci configure <tool> --profile <profile>\" to initialize a specific tool)"},
			ItemsType:   pkgCTerminal.Info,
		}
		currentConfigureTxt := fmt.Sprintf("Your curent configure is set to %v", globalConfig)
		pkgCTerminal.Print(&pkgCTerminal.TerminalOutput{
			Messages:    []string{"Global configuration has been initialized", currentConfigureTxt},
			ItemsOutput: []pkgCTerminal.TerminalItemsOutput{tItems},
		})
	}
	return nil
}

func init() {
	rootCmd.AddCommand(configureCmd)
	configureCmd.Flags().BoolVarP(&configureShow, "show", "s", false, "show configurations")
	configureCmd.Flags().BoolVarP(&configureRaw, "raw", "r", false, "output raw configurations")
}
