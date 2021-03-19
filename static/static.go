package static

import (
	"embed"
)

//go:embed "css"
//go:embed "javascript"
var FS embed.FS
