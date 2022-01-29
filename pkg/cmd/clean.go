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
	"os"
	"path/filepath"

	"github.com/nitroci/nitroci-core/pkg/core/contexts"
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
		return cleanRunner()
	},
}

func cleanRunner() error {
	workspace, err := runtimeContext.GetCurrentWorkspace()
	if err != nil {
		return err
	}
	if cleanGlobalCache {
		cachePluginsPath := runtimeContext.Cli.Settings[contexts.CFG_NAME_CACHE_PATH]
		err := os.RemoveAll(cachePluginsPath)
		if err != nil {
			return err
		}
	}
	if cleanLocalCache {
		wksCachePluginsPath := filepath.Join(workspace.WorkspaceFileFolder, "cache")
		err := os.RemoveAll(wksCachePluginsPath)
		if err != nil {
			return err
		}
	}
	return nil
}


func init() {
	rootCmd.AddCommand(cleanCmd)
	cleanCmd.Flags().BoolVar(&cleanGlobalCache, "global-cache", false, "clean global cache")
	cleanCmd.Flags().BoolVar(&cleanLocalCache, "local-cache", false, "clean local cache")
}
