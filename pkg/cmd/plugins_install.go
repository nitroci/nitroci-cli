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
	"path/filepath"

	"github.com/nitroci/nitroci-core/pkg/core/contexts"
	"github.com/nitroci/nitroci-core/pkg/core/registries"
	"github.com/nitroci/nitroci-core/pkg/core/terminal"
	"github.com/spf13/cobra"
)

var installWorkspaceCmd = &cobra.Command{
	Use:   "install",
	Short: "Install plugins",
	Long:  `Install plugins`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !runtimeContext.HasWorkspaces() {
			return errors.New("workspace is not initialized")
		}
		return pluginsInstallRunner()
	},
}

func pluginsInstallRunner() error {
	workspace, err := runtimeContext.GetCurrentWorkspace()
	if err != nil {
		return err
	}
	wksModel, err := workspace.CreateWorkspaceInstance()
	if err != nil {
		return err
	}
	currentWorkspaceTxt := fmt.Sprintf("Your curent workspace is set to %v", terminal.ConvertToCyanColor(workspace.WorkspacePath))
	if len(wksModel.Workspace.Plugins) == 0 {
		terminal.Print(&terminal.TerminalOutput{
			Messages:    []string{"Workspace doesn't require any plugin", currentWorkspaceTxt},
			ItemsOutput: []terminal.TerminalItemsOutput{},
		})
		return nil
	}
	cachePluginsPath := runtimeContext.Cli.Settings[contexts.CFG_NAME_CACHE_PLUGINS_PATH]
	wksCachePluginsPath := filepath.Join(filepath.Join(workspace.WorkspaceFileFolder, "cache"), "plugins")
	registryMap := registries.CreateRegistryMap(cachePluginsPath, wksCachePluginsPath, runtimeContext.Cli.Goos, runtimeContext.Cli.Goarch)
	pluginKeys := []string{}
	for _, plugin := range wksModel.Workspace.Plugins {
		registryKey := plugin.Registry
		if len(registryKey) == 0 {
			registryKey = runtimeContext.Cli.Settings[contexts.CFG_NAME_PLUGINS_REGISTRY]
		}
		pluginKeys = append(pluginKeys, registries.GetPackageName(plugin.Name, plugin.Version))
		err = registryMap.AddDependency(registryKey, plugin.Name, plugin.Version)
		if err != nil {
			return err
		}
	}
	tItems := terminal.TerminalItemsOutput{
		Messages:    []string{"Configured plugins:"},
		Suggestions: []string{},
		ItemsType:   terminal.Info,
		Items:       pluginKeys,
	}
	terminal.Print(&terminal.TerminalOutput{
		Messages:    []string{"Plugins are going to be installed", currentWorkspaceTxt},
		ItemsOutput: []terminal.TerminalItemsOutput{tItems},
	})

	tAction := &terminal.TerminalActionOutput{
		Step:    "Downloading plugins",
		Outputs: []string{},
	}
	terminal.PrintActions(tAction)
	printOkFunc := func(text string) {
		tAction.Outputs = append(tAction.Outputs, fmt.Sprintf("❯ %v", text))
		terminal.PrintActions(tAction)
	}
	printKoFunc := func(text string) {
		tAction.Outputs = append(tAction.Outputs, fmt.Sprintf("❯ %v", terminal.ConvertToRedColor(text)))
		terminal.PrintActions(tAction)
	}
	return registryMap.Download(printOkFunc, printKoFunc)
}

func init() {
	pluginsCmd.AddCommand(installWorkspaceCmd)
}
