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
package plugins

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	pkgCContexts "github.com/nitroci/nitroci-core/pkg/core/contexts"
	pkgCPlugins "github.com/nitroci/nitroci-core/pkg/core/plugins"
)

func LoadPlugins(runtimeContext *pkgCContexts.RuntimeContext, pluginPath string) ([]*pkgCPlugins.PluginModel, error) {
	files, err := ioutil.ReadDir(pluginPath)
	if err != nil {
		return nil, err
	}
	pluginModels := []*pkgCPlugins.PluginModel{}
	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			filePath := filepath.Join(pluginPath, fileInfo.Name(), "manifest.yml")
			if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
				continue
			}
			pluginModel, err := pkgCPlugins.LoadPluginFile(filePath)
			if err != nil {
				return nil, err
			}
			pluginModels = append(pluginModels, pluginModel)
		}
	}
	return pluginModels, nil
}
