package utils

import (
	"os"
	"path/filepath"
)

type File struct {
	Name  string
	Path  string
	IsDir bool
	Files []string
	Data  string
}

func NewFile(path string) (*File, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	fileInfo, err := os.Stat(absPath)
	if err != nil {
		return nil, err
	}

	isDir := fileInfo.IsDir()

	file := &File{
		Name:  fileInfo.Name(),
		Path:  absPath,
		IsDir: isDir,
	}

	if isDir {
		err = file.GetFiles()
		if err != nil {
			return nil, err
		}
	} else {
		err = file.Read()
		if err != nil {
			return nil, err
		}
	}

	return file, nil
}

func (f *File) Read() error {
	bytes, err := os.ReadFile(f.Path)
	if err != nil {
		return err
	}

	f.Data = string(bytes)

	return nil
}

func (f *File) GetFiles() error {
	dirEntry, err := os.ReadDir(f.Path)
	if err != nil {
		return err
	}

	for _, entry := range dirEntry {
		f.Files = append(f.Files, entry.Name())
	}

	return nil
}
