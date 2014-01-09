package server

import (
	"github.com/ultimateanu/sesame-server/filesystem"
)

type Store *map[int]*filesystem.File

func MakeStore(files []*filesystem.File) Store {
	m := make(map[int]*filesystem.File)
	for index, f := range files {
		m[index] = f
	}
	return &m
}
