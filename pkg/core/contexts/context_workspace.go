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

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/nitroci/nitroci-cli/pkg/core/workspaces"
	"github.com/nitroci/nitroci-cli/pkg/internal/config"
)

type WorkspaceContext struct {
	WorkspaceFileFolder string
	WorkspaceFileName   string
	WorkspaceHome       string
	Version             int
	Id                  string
	Name                string
}

type VirtualContext struct {
	Workspaces []*WorkspaceContext
}

func loadWorkspaceFile(projectFile string) (*workspaces.WorkspaceModel, error) {
	return config.LoadWorkspaceFile(projectFile)
}

func (w *WorkspaceContext) LoadWorkspaceFile() (*workspaces.WorkspaceModel, error) {
	return loadWorkspaceFile(w.WorkspaceFileName)
}

func (v *VirtualContext) loadVirtualContext(workspaceDepth int) *VirtualContext {
	projFiles := config.FindWorkspaceFiles()
	projFilesCount := len(projFiles)
	v.Workspaces = make([]*WorkspaceContext, projFilesCount)
	if projFilesCount == 0 {
		return v
	}
	for i, projFile := range projFiles {
		workspaceModel, _ := loadWorkspaceFile(projFile)
		var wksContext = WorkspaceContext{}
		wksContext.WorkspaceFileName = projFile
		wksContext.WorkspaceFileFolder = filepath.Dir(projFile)
		wksContext.WorkspaceHome = wksContext.WorkspaceFileFolder
		if strings.HasSuffix(wksContext.WorkspaceHome, wksContext.WorkspaceFileFolder) {
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
