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
	"os"
	"path/filepath"
)

func inverseRecursiveFindFilesInPaths(rootPath string, targetPath string, folderName string, fileName string, files []string) {
	rel, _ := filepath.Rel(rootPath, targetPath)
	if rel == "." {
		return
	}
	searchPath := fmt.Sprintf("%v/%v", targetPath, folderName)
	filePath := fmt.Sprintf("%v/%v", searchPath, fileName)
	_, err := os.Stat(filePath)
	if err == nil || os.IsExist(err) {
		files = append(files, fmt.Sprintf("%v%v/%v/%v", rootPath, rel, folderName, fileName))
	} else {
		filePath := fmt.Sprintf("%v/%v", targetPath, fileName)
		_, err := os.Stat(filePath)
		if err == nil || os.IsExist(err) {
			files = append(files, fmt.Sprintf("%v%v/%v", rootPath, rel, fileName))
		}
	}
	targetPath += "/.."
	inverseRecursiveFindFilesInPaths(rootPath, targetPath, folderName, fileName, files)
}

func inverseRecursiveFindFiles(targetPath string, folderName string, fileName string) (files []string) {
	basePath := "/"
	paths := []string{}
	inverseRecursiveFindFilesInPaths(basePath, targetPath, folderName, fileName, paths)
	return paths
}
