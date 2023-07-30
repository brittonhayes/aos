package aos

import (
	"embed"
)

//go:embed fixtures/*.yaml
var FIXTURES embed.FS

//go:embed web/src/docs.html
var HTML_DOCS []byte

//go:embed web/src/index.html
var HTML_HOME []byte
