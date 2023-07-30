package aos

import (
	"embed"
)

//go:embed fixtures/*.yaml
var FIXTURES embed.FS

//go:embed web/docs.html
var OPEN_API_DOCS []byte

//go:embed web/index.html
var HOMEPAGE []byte
