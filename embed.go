package warhammer

import (
	"embed"
)

//go:embed fixtures/*.yaml
var FIXTURES embed.FS

//go:embed web/index.html
var OPEN_API_DOCS []byte
