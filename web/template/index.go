package template

import "embed"

//go:embed *.twig

var Contents embed.FS
