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
	pkgCCore "github.com/nitroci/nitroci-core/pkg/core"
	pkgCCtx "github.com/nitroci/nitroci-core/pkg/core/contexts"

	"github.com/spf13/cobra"
)

var environmentsPipelinesCmd = &cobra.Command{
	Use:   "envs",
	Short: "Print workspace environments information",
	Long:  `Print workspace environments information`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, err := pkgCCore.CreateWorspaceContext(ctxInput)
		if err != nil {
			return err
		}
		return environmentsRunner(ctx)
	},
}

func environmentsRunner(ctx pkgCCtx.CoreContexter) error {
	return nil
}

func init() {
	rootCmd.AddCommand(environmentsPipelinesCmd)
}
