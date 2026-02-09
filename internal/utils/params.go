package utils

import "net/http"

type GetParams struct {
	Path string
	Cmd  string
}

func (p *GetParams) Parse(r *http.Request) {
	query := r.URL.Query()
	p.Path = query.Get("path")
	p.Cmd = query.Get("cmd")
}
