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

	"nitroci/pkg/core/workspaces"

	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

const (
	WorkspaceFileFolder = ".nitroci"
	WorkspaceFileName   = "workspace.yml"
)

func FindWorkspaceFiles() (workspaceFiles []string) {
	targetPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return inverseRecursiveFindFiles(targetPath, WorkspaceFileFolder, WorkspaceFileName)
}

func FindCurrentWorkspaceFile(workspaceDepth int) (file string) {
	workspaceFiles := FindWorkspaceFiles()
	if len(workspaceFiles) == 0 || len(workspaceFiles) <= workspaceDepth {
		empty := ""
		return empty
	}
	return (workspaceFiles)[workspaceDepth]
}

func LoadWorkspaceFile(workspaceFile string) (*workspaces.WorkspaceModel, error) {
	var workspaceModel = &workspaces.WorkspaceModel{}
	yamlFile, err := ioutil.ReadFile(workspaceFile)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, workspaceModel)
	if err != nil {
		return nil, err
	}
	return workspaceModel, nil
}

func SaveWorkspaceFile(workspaceModel *workspaces.WorkspaceModel, workspaceFile string) {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	workspaceModel.Workspace.ID = uuid
	yamlData, _ := yaml.Marshal(&workspaceModel)
	fmt.Println(string(yamlData))
}
