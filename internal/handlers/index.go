package handlers

import (
	"html/template"
	"net/http"

	. "github.com/deadblackclover/dotshell/internal/utils"
)

type Data struct {
	System System
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	d := getData()
	t, _ := template.ParseFiles("internal/templates/index.html")
	t.Execute(w, d)
}

func getData() Data {
	var data Data
	data.System = GetSystemData()
	return data
}
