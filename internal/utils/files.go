package utils

import (
	"os"
	"path/filepath"
	"time"
)

type FileInfo struct {
	Name    string
	Path    string
	ModTime string
	Perm    string
	Size    int64
	IsDir   bool
}

func (f *FileInfo) initFromPath(path string) error {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return err
	}

	f.Name = fileInfo.Name()
	f.Path = path
	f.IsDir = fileInfo.IsDir()

	return nil
}

func (f *FileInfo) initFromDirEntry(entry os.DirEntry) error {
	fileInfo, err := entry.Info()
	if err != nil {
		return err
	}

	f.Name = fileInfo.Name()
	f.IsDir = fileInfo.IsDir()
	f.Size = fileInfo.Size()
	f.ModTime = fileInfo.ModTime().Format(time.RFC822Z)
	f.Perm = fileInfo.Mode().Perm().String()

	return nil
}

type File struct {
	Back     string
	Data     string
	FileInfo *FileInfo
	Files    []*FileInfo
}

func NewFile(path string) (*File, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	fileInfo := new(FileInfo)
	err = fileInfo.initFromPath(absPath)
	if err != nil {
		return nil, err
	}

	file := new(File)
	file.FileInfo = fileInfo
	file.Back = filepath.Join(absPath, "../")

	if fileInfo.IsDir {
		err = file.getFiles()
		if err != nil {
			return nil, err
		}
	} else {
		err = file.read()
		if err != nil {
			return nil, err
		}
	}

	return file, nil
}

func (f *File) read() error {
	bytes, err := os.ReadFile(f.FileInfo.Path)
	if err != nil {
		return err
	}

	f.Data = string(bytes)

	return nil
}

func (f *File) getFiles() error {
	dirEntry, err := os.ReadDir(f.FileInfo.Path)
	if err != nil {
		return err
	}

	for _, entry := range dirEntry {
		files := new(FileInfo)
		files.initFromDirEntry(entry)
		files.Path = filepath.Join(f.FileInfo.Path, files.Name)

		f.Files = append(f.Files, files)
	}

	return nil
}
