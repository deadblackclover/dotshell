// Copyright (c) 2026, DEADBLACKCLOVER.

// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
// General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
package handlers

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"

	. "github.com/deadblackclover/dotshell/internal/utils"
)

//go:embed index.tmpl
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
