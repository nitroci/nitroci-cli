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
	pkgCCtx "github.com/nitroci/nitroci-core/pkg/core/contexts"
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
		ctx, err := pkgCCore.CreateWorspaceContext(ctxInput)
		if err != nil {
			return err
		}
		return cleanRunner(ctx)
	},
}

func cleanRunner(ctx pkgCCtx.CoreContexter) error {
	if !cleanGlobalCache && !cleanLocalCache {
		return nil
	}
	runtimeCtx := ctx.GetRuntimeCtx()
	workspace, err := runtimeCtx.GetCurrentWorkspace()
	if err != nil {
		return err
	}
	terminal := ctx.GetTerminal()
	currentWorkspaceTxt := fmt.Sprintf("Your curent workspace is set to %v", terminal.ConvertToCyanColor(workspace.GetWorkspacePath()))
	terminal.Print(&pkgCTerminal.TerminalOutput{
		Messages: []string{"Cache is going to be cleaned", currentWorkspaceTxt},
	})
	tAction := &pkgCTerminal.TerminalActionOutput{
		Step:    "Cleaning cache",
		Outputs: []string{},
	}
	terminal.PrintActions(tAction)
	if cleanGlobalCache {
		cachePluginsPath, _ := runtimeCtx.GetSettings("NITROCI_CACHE")
		err := os.RemoveAll(cachePluginsPath)
		if err != nil {
			tAction.Outputs = append(tAction.Outputs, fmt.Sprintf("• %v", terminal.ConvertToRedColor(cachePluginsPath)))
			terminal.PrintActions(tAction)
			return err
		}
		tAction.Outputs = append(tAction.Outputs, fmt.Sprintf("• %v", cachePluginsPath))
		terminal.PrintActions(tAction)
	}
	if cleanLocalCache {
		wksCachePluginsPath := filepath.Join(workspace.GetWorkspaceFileFolder(), "cache")
		err := os.RemoveAll(wksCachePluginsPath)
		if err != nil {
			tAction.Outputs = append(tAction.Outputs, fmt.Sprintf("• %v", terminal.ConvertToRedColor(wksCachePluginsPath)))
			terminal.PrintActions(tAction)
			return err
		}
		tAction.Outputs = append(tAction.Outputs, fmt.Sprintf("• %v", wksCachePluginsPath))
		terminal.PrintActions(tAction)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().BoolVar(&cleanGlobalCache, "global-cache", false, "clean global cache")
	cleanCmd.Flags().BoolVar(&cleanLocalCache, "local-cache", false, "clean local cache")
}
