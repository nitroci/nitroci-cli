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
package contexts

type CliContext struct {
	ConfigHome  	string
	ConfigFile  	string
	CacheHome   	string
	Profile     	string
	Verbose     	bool
	WorkspaceDepth 	int
	Environment		string
}

func (c *CliContext) loadCliContext(profile string, verbose bool, workspaceDepth int) *CliContext {
	c.Profile = profile
	c.Verbose = verbose
	c.WorkspaceDepth = workspaceDepth
	return c
}
