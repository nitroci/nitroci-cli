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
package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

const (
	projectFolderName = ".nitroci"
	projectFileName = "nitroci.yml"
)

func FindProjectFiles() (files *[]string) {
	targetPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return InverseRecursiveFindFiles(targetPath, projectFolderName, projectFileName)
}

func FindCurrentProjectFile(depth int) (file *string) {
	projectFiles := FindProjectFiles()
	if len(*projectFiles) == 0 || len(*projectFiles) <= depth {
		empty := ""
		return &empty
	}
	return &(*projectFiles)[depth]
}

func (config *ProjectConfiguration) LoadProject(path *string) *ProjectConfiguration {
	yamlFile, err := ioutil.ReadFile(*path)
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
	return config
}

func (config *ProjectConfiguration) SaveProject(path *string) *ProjectConfiguration {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	config.Workspace.ID = uuid
	yamlData, _ := yaml.Marshal(&config)
	fmt.Println(string(yamlData))
	return config
}
