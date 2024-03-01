//go:build dev

package web

import (
	"io/fs"
	"os"
	"path/filepath"
)

var ViewsFolder = getWebFolderDirFs()
var StaticFolder = getWebFolderDirFs()

// Assumes you are in the project root
func getWebFolderDirFs() fs.FS {
	wd, _ := os.Getwd()
	p := filepath.Join(wd, "web/")
	return os.DirFS(p)
}
