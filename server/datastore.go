package server

import (
	"errors"
	"github.com/ultimateanu/sesame-server/filesystem"
)

type Store struct {
	NameMap    map[string][]*filesystem.File
	FilesIndex string
}

func MakeStore(files []*filesystem.File) *Store {
	m := make(map[string][]*filesystem.File)
	for _, f := range files {
		urlSafeName := UrlSafe(f.Name)
		m[urlSafeName] = append(m[urlSafeName], f)
	}
	return &Store{m, ""}
}

func (s *Store) Contains(fileName string) bool {
	if len((*s).NameMap[fileName]) == 0 {
		return false
	}
	return true
}

func (s *Store) Count(fileName string) int {
	return len((*s).NameMap[fileName])
}

func (s *Store) Get(fileName string) []*filesystem.File {
	if !s.Contains(fileName) {
		return nil
	}
	return s.NameMap[fileName]
}

func (s *Store) GetFile(fileName string, index int) (*filesystem.File, error) {
	if index >= s.Count(fileName) {
		return nil, errors.New("store: index out of bounds")
	}
	return s.NameMap[fileName][index], nil
}
