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
	"path/filepath"

	pkgCCore "github.com/nitroci/nitroci-core/pkg/core"
	pkgCContexts "github.com/nitroci/nitroci-core/pkg/core/contexts"
	pkgCRegistries "github.com/nitroci/nitroci-core/pkg/core/registries"
	pkgCTerminal "github.com/nitroci/nitroci-core/pkg/core/terminal"

	"github.com/spf13/cobra"
)

var installWorkspaceCmd = &cobra.Command{
	Use:   "install",
	Short: "Install plugins",
	Long:  `Install plugins`,
	RunE: func(cmd *cobra.Command, args []string) error {
		runtimeCtx, _ := pkgCCore.CreateAndInitalizeContext(pkgCContexts.CORE_BUILDER_WORKSPACE_TYPE)
		return pluginsInstallRunner(runtimeCtx)
	},
}

func pluginsInstallRunner(runtimeCtx pkgCContexts.RuntimeContexter) error {
	workspace, err := runtimeCtx.GetCurrentWorkspace()
	if err != nil {
		return err
	}
	wksModel, err := workspace.CreateWorkspaceInstance()
	if err != nil {
		return err
	}
	currentWorkspaceTxt := fmt.Sprintf("Your curent workspace is set to %v", pkgCTerminal.ConvertToCyanColor(workspace.GetWorkspacePath()))
	if len(wksModel.Workspace.Plugins) == 0 {
		pkgCTerminal.Print(&pkgCTerminal.TerminalOutput{
			Messages:    []string{"Workspace doesn't require any plugin", currentWorkspaceTxt},
			ItemsOutput: []pkgCTerminal.TerminalItemsOutput{},
		})
		return nil
	}
	cachePluginsPath, _ := runtimeCtx.GetSettings(pkgCContexts.CFG_NAME_CACHE_PLUGINS_PATH)
	wksCachePluginsPath := filepath.Join(filepath.Join(workspace.GetWorkspaceFileFolder(), "cache"), "plugins")
	registryMap := pkgCRegistries.CreateRegistryMap(cachePluginsPath, wksCachePluginsPath, runtimeCtx.GetGoos(), runtimeCtx.GetGoarch())
	pluginKeys := []string{}
	for _, plugin := range wksModel.Workspace.Plugins {
		registryKey := plugin.Registry
		if len(registryKey) == 0 {
			registryKey, _ = runtimeCtx.GetSettings(pkgCContexts.CFG_NAME_PLUGINS_REGISTRY)
		}
		pluginKeys = append(pluginKeys, pkgCRegistries.GetPackageName(plugin.Name, plugin.Version))
		err = registryMap.AddDependency(registryKey, plugin.Name, plugin.Version)
		if err != nil {
			return err
		}
	}
	tItems := pkgCTerminal.TerminalItemsOutput{
		Messages:    []string{"Configured plugins:"},
		Suggestions: []string{},
		ItemsType:   pkgCTerminal.Info,
		Items:       pluginKeys,
	}
	pkgCTerminal.Print(&pkgCTerminal.TerminalOutput{
		Messages:    []string{"Plugins are going to be installed", currentWorkspaceTxt},
		ItemsOutput: []pkgCTerminal.TerminalItemsOutput{tItems},
	})

	tAction := &pkgCTerminal.TerminalActionOutput{
		Step:    "Downloading plugins",
		Outputs: []string{},
	}
	pkgCTerminal.PrintActions(tAction)
	printOkFunc := func(text string) {
		tAction.Outputs = append(tAction.Outputs, fmt.Sprintf("• %v", text))
		pkgCTerminal.PrintActions(tAction)
	}
	printKoFunc := func(text string) {
		tAction.Outputs = append(tAction.Outputs, fmt.Sprintf("• %v", pkgCTerminal.ConvertToRedColor(text)))
		pkgCTerminal.PrintActions(tAction)
	}
	return registryMap.Download(printOkFunc, printKoFunc)
}

func init() {
	pluginsCmd.AddCommand(installWorkspaceCmd)
}
