package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type FileLocator struct {
	paths []string
}

func NewFileLocator(paths []string) *FileLocator {
	return &FileLocator{paths: paths}
}

func (this *FileLocator) LocateAll(name string) []string {
	return this.locate(name, nil, false)
}

func (this *FileLocator) LocateFirst(name string) []string {
	return this.locate(name, nil, true)
}

func (this *FileLocator) LocateFirstFrom(name string, currentPath string) []string {
	return this.locate(name, &currentPath, true)
}

func (this *FileLocator) locate(name string, curr *string, first bool) []string {
	if filepath.IsAbs(name) {
		if !FileExists(name) {
			panic(fmt.Sprintf("The file %v doesn't exist.", name))
		}

		return []string{name}
	}

	current := ""
	if curr != nil {
		current = *curr
	}

	filepaths := make([]string, 0)

	addFileFunc := func(p string) {
		file := path.Join(p, name)
		if FileExists(file) {
			filepaths = append(filepaths, file)
		}
	}

	if curr != nil {
		addFileFunc(current)
		if first {
			return filepaths
		}
	}

	for _, p := range this.paths {
		addFileFunc(p)
		if first {
			return filepaths
		}
	}

	if len(filepaths) == 0 {
		panic(fmt.Sprintf("The file %v doesn't exist in %v, %v", name, current, strings.Join(this.paths, ", ")))
	}

	return filepaths
}

func FileExists(name string) bool {
	stat, _ := os.Stat(name)
	return stat != nil
}
