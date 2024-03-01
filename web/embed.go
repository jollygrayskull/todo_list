//go:build !dev

package web

import (
	"embed"
)

//go:embed views
var ViewsFolder embed.FS

//go:embed static
var StaticFolder embed.FS
