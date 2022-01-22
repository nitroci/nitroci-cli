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
	"nitroci/pkg/internal/config"

	"github.com/spf13/cobra"
)

var (
	FlagBitbucketDomain, FlagBitbucketUsername, FlagBitbucketPassword string
)

var bitbucketConfigureCmd = &cobra.Command{
	Use:   "bitbucket",
	Short: "Configure Bitbucket",
	Long:  `Configure Bitbucket`,
	Run: func(cmd *cobra.Command, args []string) {
		configureBitbucketRunner()
	},
}

func configureBitbucketRunner() {
	domain := FlagBitbucketDomain
	if len(domain) == 0 {
		_, domain = config.PromptGlobalConfigKey(runtimeContext.Cli.Profile, "Workspace", false)
		config.SetGlobalConfigString(runtimeContext.Cli.Profile, "bitbucket_workspace", domain)
	}
	username := FlagBitbucketUsername
	if len(username) == 0 {
		_, username = config.PromptGlobalConfigKey(runtimeContext.Cli.Profile, "Username", false)
		config.SetGlobalConfigString(runtimeContext.Cli.Profile, "bitbucket_username", domain)
	}
	password := FlagBitbucketPassword
	if len(password) == 0 {
		_, password = config.PromptGlobalConfigKey(runtimeContext.Cli.Profile, "Password", true)
		config.SetGlobalConfigString(runtimeContext.Cli.Profile, "bitbucket_secret", domain)
	}
}

func init() {
	bitbucketConfigureCmd.Flags().StringVar(&FlagBitbucketDomain, "workspace", "", "workspace")
	bitbucketConfigureCmd.Flags().StringVar(&FlagBitbucketUsername, "user", "", "username")
	bitbucketConfigureCmd.Flags().StringVar(&FlagBitbucketPassword, "pass", "", "password")
}
