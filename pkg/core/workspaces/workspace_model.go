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
package workspaces

type WorkspaceModel struct {
	Version      int            `yaml:"version"`
	Workspace    Workspace      `yaml:"workspace,omitempty"`
	Environments []Environments `yaml:"environments,omitempty"`
	Commands     []Commands     `yaml:"commands"`
}
type Bitbucket struct {
	Workspace string `yaml:"workspace"`
	Slug      string `yaml:"slug"`
}
type Pipelines struct {
	Platform  string    `yaml:"platform,omitempty"`
	Bitbucket Bitbucket `yaml:"bitbucket,omitempty"`
	Suffix    string    `yaml:"suffix,omitempty"`
}
type Packages struct {
	Platform string `yaml:"platform,omitempty"`
}
type SettingsEncryption struct {
	SecretKey string `yaml:"secret_key,omitempty"`
}
type Settings struct {
	Pipelines  Pipelines          `yaml:"pipelines,omitempty"`
	Packages   Packages           `yaml:"packages,omitempty"`
	Encryption SettingsEncryption `yaml:"encryption,omitempty"`
}
type Workspace struct {
	ID       string   `yaml:"id"`
	Name     string   `yaml:"name,omitempty"`
	Settings Settings `yaml:"settings,omitempty"`
}
type Encryption struct {
	Paths []string `yaml:"paths,omitempty"`
}
type Environments struct {
	Name       string     `yaml:"name,omitempty"`
	Encryption Encryption `yaml:"encryption,omitempty"`
}
type Steps struct {
	Cwd     string   `yaml:"cwd,omitempty"`
	Scripts []string `yaml:"scripts,omitempty"`
	Script  string   `yaml:"script,omitempty"`
}
type Commands struct {
	Name        string  `yaml:"name,omitempty"`
	Description string  `yaml:"description,omitempty"`
	Steps       []Steps `yaml:"steps,omitempty"`
}
