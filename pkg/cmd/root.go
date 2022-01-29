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
	"github.com/nitroci/nitroci-core/pkg/core/terminal"
	cliPlugins "github.com/nitroci/nitroci-cli/pkg/plugins"
	"github.com/spf13/cobra"
)

var runtimeContext *contexts.RuntimeContext

var rootCmd = &cobra.Command{
	Use:   "nitroci",
	Short: "NitroCI - Boost your software development",
	Long: `NitroCI - Boost your software development
	
The command line interface implementes devops practices 
and it is not tied to a particolar language or farmework.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },ls
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true
	err := rootCmd.Execute()
	if err != nil {
		terminal.Println(terminal.ConvertToRedColor(err.Error()))
		os.Exit(1)
	}
}

func initRoot() {
	profile, _ := rootCmd.PersistentFlags().GetString("profile")
	verbose, _ := rootCmd.PersistentFlags().GetBool("verbose")
	workspaceDepth, _ := rootCmd.PersistentFlags().GetInt("workspace")
	var err error
	runtimeContext, err = contexts.LoadRuntimeContext(profile, "", workspaceDepth-1, verbose)
	if err != nil {
		terminal.Println(terminal.ConvertToRedColor(err.Error()))
		os.Exit(1)
	}
	workspace, err := runtimeContext.GetCurrentWorkspace()
	if err != nil {
		terminal.Println(terminal.ConvertToRedColor(err.Error()))
		os.Exit(1)
	}
	pluginPath := filepath.Join(filepath.Join(workspace.WorkspaceFileFolder, "cache"), "plugins")
	pluginModels, err := cliPlugins.LoadPlugins(runtimeContext, pluginPath)
	if err != nil {
		terminal.Println(terminal.ConvertToRedColor(err.Error()))
		os.Exit(1)
	}
	terminal.Println(len(pluginModels))
}

func init() {
	cobra.OnInitialize(initRoot)
	rootCmd.PersistentFlags().StringP("profile", "p", "default", "set a specific profile")
	rootCmd.PersistentFlags().IntP("workspace", "w", 1, "set current workspace")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "output verbose output")
}
