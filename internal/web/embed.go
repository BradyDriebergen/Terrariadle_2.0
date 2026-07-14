package web

import (
	"embed"
	"io/fs"
)

//go:embed all:build
var embeddedFiles embed.FS

func Assets() fs.FS {
	sub, err := fs.Sub(embeddedFiles, "build")
	if err != nil {
		panic(err)
	}
	return sub
}
