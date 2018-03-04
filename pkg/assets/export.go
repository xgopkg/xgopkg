package assets

import (
	"github.com/elazarl/go-bindata-assetfs"
)

//FS export assetFS function
func FS(path string) *assetfs.AssetFS {
	afs := &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: AssetInfo, Prefix: path}
	return afs
}
