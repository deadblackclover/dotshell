package handlers

import (
	"fmt"
	"log"
	"net/http"

	. "github.com/deadblackclover/dotshell/internal/utils"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	// Get all the values
	params := &GetParams{}
	params.Parse(r)

	log.Printf("Download=%s", params.Path)

	// Getting data from the path
	if file, err := NewFile(params.Path); err == nil {
		if !file.FileInfo.IsDir {
			value := fmt.Sprintf("attachment; filename=%s", file.FileInfo.Name)
			w.Header().Set("Content-Disposition", value)
			http.ServeFile(w, r, params.Path)
		}
	} else {
		log.Println(err)
	}
}
