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
	//"os"
	//"strings"

	//"github.com/nitroci/nitroci-core/pkg/core/io/terminal"
	//"github.com/nitroci/nitroci-core/pkg/internal/config"

	//"github.com/nitroci/nitroci-core/pkg/core/net/http"

	"github.com/spf13/cobra"
)

var (
	FlagJFrogDomain, FlagJFrogUsername, FlagJFrogPassword string
)

var jfrogConfigureCmd = &cobra.Command{
	Use:   "jfrog",
	Short: "Configure JFrog",
	Long:  `Configure JFrog`,
	Run: func(cmd *cobra.Command, args []string) {
		configureJFrogRunner()
	},
}

func configureJFrogRunner() {
	/*
	domain := FlagJFrogDomain
	if len(domain) == 0 {
		_, domain = config.PromptGlobalConfigKey(runtimeContext.Cli.Profile, "Domain", false)
	}
	username := FlagJFrogUsername
	if len(username) == 0 {
		_, username = config.PromptGlobalConfigKey(runtimeContext.Cli.Profile, "Username", false)

	}
	password := FlagJFrogPassword
	if len(password) == 0 {
		_, password = config.PromptGlobalConfigKey(runtimeContext.Cli.Profile, "Password", true)
	}
	httpResult, err := http.HttpGet("https://"+domain+".jfrog.io/"+domain+"/api/npm/auth", username, password)
	if err != nil || httpResult.StatusCode != 200 {
		errMessage := "Operation cannot be completed. Please verify the inputs."
		terminal.Print(&terminal.TerminalOutput{
			Messages:    []string{errMessage},
			MessageType: terminal.Error,
		})
		os.Exit(1)
	}
	for _, line := range strings.Split(strings.TrimSuffix(*httpResult.Body, "\n"), "\n") {
		s := strings.Split(line, " = ")
		if s[0] == "_auth" {
			config.SetGlobalConfigString(runtimeContext.Cli.Profile, "jfrog_secret", s[1])
		} else if s[0] == "email" {
			config.SetGlobalConfigString(runtimeContext.Cli.Profile, "jfrog_username", s[1])
		}
	}
	*/
}

func init() {
	jfrogConfigureCmd.Flags().StringVar(&FlagJFrogDomain, "domain", "", "domain")
	jfrogConfigureCmd.Flags().StringVar(&FlagJFrogUsername, "user", "", "username")
	jfrogConfigureCmd.Flags().StringVar(&FlagJFrogPassword, "pass", "", "password")
}
