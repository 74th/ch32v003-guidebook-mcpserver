package main

import (
	"embed"
)

//go:embed documents.yaml
var DocumentManifestBlob []byte

//go:embed documents/*/*.md
var DocumentFS embed.FS

//go:embed sample_codes/*/main.c sample_codes/*/src/main.c
var CodeFS embed.FS
