package main

import (
	"io/ioutil"
	"talisman/git_repo"

	log "github.com/Sirupsen/logrus"

	"github.com/bmatcuk/doublestar"
)

type DirectoryHook struct{}

func NewDirectoryHook() *DirectoryHook {
	return &DirectoryHook{}
}

func (p *DirectoryHook) GetFilesFromDirectory(globPattern string) []git_repo.Addition {
	var result []git_repo.Addition

	files, _ := doublestar.Glob(globPattern)
	for _, file := range files {
		data, err := ReadFile(file)

		if err != nil {
			continue
		}

		newAddition := git_repo.NewAddition(file, data)
		result = append(result, newAddition)
	}

	return result
}

func ReadFile(filepath string) ([]byte, error) {
	log.Debugf("reading file %s", filepath)
	return ioutil.ReadFile(filepath)
}
