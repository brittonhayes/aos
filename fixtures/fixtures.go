package fixtures

import (
	"embed"
)

//go:embed *.yaml
var FS embed.FS

const (
	ALLEGIANCES = "allegiances.yaml"
	ALLIANCES   = "alliances.yaml"
	CITIES      = "cities.yaml"
	STRATEGIES  = "strategies.yaml"
	UNITS       = "units.yaml"
	WARSCROLLS  = "warscrolls.yaml"
)
