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
package context

import (
	"errors"
	"path/filepath"
	"strings"

	"nitroci/pkg/internal/config"
)

type WorkspaceContext struct {
	WorkspaceFileHome string
	WorkspaceFile     string
	WorkspaceHome     string
	Version           int
	Id                string
	Name              string
}

type VirtualContext struct {
	Workspaces []*WorkspaceContext
}

func loadWorkspaceModel(projectFile string) (*config.WorkspaceModel, error) {
	return config.LoadProject(projectFile)
}

func (w *WorkspaceContext) LoadWorkspaceModel() (*config.WorkspaceModel, error) {
	return loadWorkspaceModel(w.WorkspaceFile)
}

func (v *VirtualContext) loadVirtualContext(workspaceDepth int) *VirtualContext {
	projFiles := config.FindProjectFiles()
	projFilesCount := len(*projFiles)
	v.Workspaces = make([]*WorkspaceContext, projFilesCount)
	if projFilesCount == 0 {
		return v
	}
	for i, projFile := range *projFiles {
		workspaceModel, _ := loadWorkspaceModel(projFile)
		var wksContext = WorkspaceContext{}
		wksContext.WorkspaceFile = projFile
		wksContext.WorkspaceFileHome = filepath.Dir(projFile)
		wksContext.WorkspaceHome = wksContext.WorkspaceFileHome
		if strings.HasSuffix(wksContext.WorkspaceHome, config.ProjectFolderName) {
			wksContext.WorkspaceHome = filepath.Dir(wksContext.WorkspaceHome)
		}
		wksContext.Version = workspaceModel.Version
		wksContext.Id = workspaceModel.Workspace.ID
		wksContext.Name = workspaceModel.Workspace.Name
		v.Workspaces[i] = &wksContext
	}
	return v
}

func (v *VirtualContext) hasWorkspaces() bool {
	return v.Workspaces != nil && len(v.Workspaces) > 0
}

func (v *VirtualContext) getWorkspaces() ([]*WorkspaceContext, error) {
	if v.Workspaces == nil {
		return nil, errors.New("invalid workspace depth")
	}
	return v.Workspaces, nil
}

func (v *VirtualContext) getWorkspace(workspaceDepth int) (*WorkspaceContext, error) {
	if v.Workspaces == nil || len(v.Workspaces) <= workspaceDepth {
		return nil, errors.New("invalid workspace depth")
	}
	return v.Workspaces[workspaceDepth], nil
}
