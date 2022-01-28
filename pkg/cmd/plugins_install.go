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
	"path"

	"github.com/nitroci/nitroci-core/pkg/core/contexts"
	"github.com/nitroci/nitroci-core/pkg/core/registries"
	"github.com/nitroci/nitroci-core/pkg/core/registries/plugins"
	"github.com/spf13/cobra"
)

var installWorkspaceCmd = &cobra.Command{
	Use:   "install",
	Short: "Install plugins",
	Long:  `Install plugins`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return pluginsInstallRunner()
	},
}

func pluginsInstallRunner() error {
	if !runtimeContext.HasWorkspaces() {
		workspaceRunner()
		return nil
	}
	workspace, _ := runtimeContext.GetCurrentWorkspace()
	wksModel, _ := workspace.CreateWorkspaceInstance()
	cachePluginsPath := runtimeContext.Cli.Settings[contexts.CFG_NAME_CACHE_PLUGINS_PATH]
	for _,m := range wksModel.Workspace.Plugins {
		registryName := m.Registry
		if len(registryName) == 0 {
			registryName = runtimeContext.Cli.Settings[contexts.CFG_NAME_PLUGINS_REGISTRY]
		}
		registry, err := registries.GetRegistry(registryName)
		if err != nil {
			return fmt.Errorf("invalid registry %v", registryName)
		}
		wksCachePluginsPath := path.Join(path.Join(workspace.WorkspaceFileFolder, "cache"), "packages")
		err = registries.Download(cachePluginsPath, wksCachePluginsPath, registry, plugins.DownloadPlugins)
		if err != nil {
			return err
		}
	}
	return nil
}

func init() {
	pluginsCmd.AddCommand(installWorkspaceCmd)
}