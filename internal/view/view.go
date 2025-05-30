package view

import (
	"embed"
	_ "embed"
)

//go:embed index.templ
var IndexTemplate embed.FS
