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
	"os"
	"path/filepath"

	pkgCCore "github.com/nitroci/nitroci-core/pkg/core"
	pkgCContexts "github.com/nitroci/nitroci-core/pkg/core/contexts"
	pkgCTerminal "github.com/nitroci/nitroci-core/pkg/core/terminal"

	"github.com/spf13/cobra"
)

var (
	cleanGlobalCache, cleanLocalCache bool
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove object files and cached files",
	Long:  `Remove object files and cached files`,
	RunE: func(cmd *cobra.Command, args []string) error {
		runtimeCtx, err := pkgCCore.CreateAndInitalizeContext(pkgCContexts.CORE_BUILDER_WORKSPACE_TYPE)
		if err != nil {
			return err
		}
		return cleanRunner(runtimeCtx)
	},
}

func cleanRunner(runtimeCtx pkgCContexts.RuntimeContexter) error {
	if !cleanGlobalCache && !cleanLocalCache {
		return nil
	}
	workspace, err := runtimeCtx.GetCurrentWorkspace()
	if err != nil {
		return err
	}
	currentWorkspaceTxt := fmt.Sprintf("Your curent workspace is set to %v", pkgCTerminal.ConvertToCyanColor(workspace.GetWorkspacePath()))
	pkgCTerminal.Print(&pkgCTerminal.TerminalOutput{
		Messages: []string{"Cache is going to be cleaned", currentWorkspaceTxt},
	})
	tAction := &pkgCTerminal.TerminalActionOutput{
		Step:    "Cleaning cache",
		Outputs: []string{},
	}
	pkgCTerminal.PrintActions(tAction)
	if cleanGlobalCache {
		cachePluginsPath, _ := runtimeCtx.GetSettings(pkgCContexts.CFG_NAME_CACHE_PATH)
		err := os.RemoveAll(cachePluginsPath)
		if err != nil {
			tAction.Outputs = append(tAction.Outputs, fmt.Sprintf("• %v", pkgCTerminal.ConvertToRedColor(cachePluginsPath)))
			pkgCTerminal.PrintActions(tAction)
			return err
		}
		tAction.Outputs = append(tAction.Outputs, fmt.Sprintf("• %v", cachePluginsPath))
		pkgCTerminal.PrintActions(tAction)
	}
	if cleanLocalCache {
		wksCachePluginsPath := filepath.Join(workspace.GetWorkspaceFileFolder(), "cache")
		err := os.RemoveAll(wksCachePluginsPath)
		if err != nil {
			tAction.Outputs = append(tAction.Outputs, fmt.Sprintf("• %v", pkgCTerminal.ConvertToRedColor(wksCachePluginsPath)))
			pkgCTerminal.PrintActions(tAction)
			return err
		}
		tAction.Outputs = append(tAction.Outputs, fmt.Sprintf("• %v", wksCachePluginsPath))
		pkgCTerminal.PrintActions(tAction)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().BoolVar(&cleanGlobalCache, "global-cache", false, "clean global cache")
	cleanCmd.Flags().BoolVar(&cleanLocalCache, "local-cache", false, "clean local cache")
}
