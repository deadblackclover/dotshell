package utils

import (
	"log"
	"os"
)

type File struct {
	Name  string
	IsDir bool
}

func GetFiles(path string) []File {
	log.Printf("Getting a list of files in %s", path)

	dirEntry, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	var result []File

	for _, entry := range dirEntry {
		var file File

		file.Name = entry.Name()
		file.IsDir = entry.IsDir()

		result = append(result, file)
	}

	return result
}
