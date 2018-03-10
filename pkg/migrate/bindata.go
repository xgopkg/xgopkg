// Code generated by go-bindata. DO NOT EDIT.
// sources:
// .DS_Store
// README.MD
// README_zh-CN.MD
// migrations/201803042326_add_package.down.sql
// migrations/201803042326_add_package.up.sql
// test/clean.sql
// test/data.sql
package migrate

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _Ds_store = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x98\x4d\x6a\x84\x40\x10\x85\x5f\x75\x5c\x34\xc9\xa6\x97\x59\xf6\x15\x72\x83\x46\xcc\x09\xbc\x40\x02\x01\x11\x12\x05\xf3\xb3\x76\x95\x73\xe5\x68\x41\xfa\x85\x08\xea\xac\x66\x18\x67\x78\x1f\x34\xdf\xc2\x2a\x95\x5e\x74\x57\x15\x00\x2b\x3f\x5f\x1e\x80\x00\xc0\x23\x1b\xdf\x58\xc5\x73\x2d\x70\xb4\xe5\x15\x80\x67\x7c\xa0\x45\x8f\x0e\xef\xeb\xef\x5a\x30\xe5\xde\xe2\x0d\x2d\x1a\x0c\xf3\xfc\xaf\x7a\xe8\x5e\xfb\xae\x61\x8c\x10\x42\x08\x21\x8e\x03\xef\x55\x7f\x77\xee\x1f\x11\x42\xec\x8e\xe9\x7c\x88\x74\xa2\xc7\x6c\xe3\x73\x47\x17\xb3\x9c\x40\x47\x3a\xd1\x63\xb6\x31\xce\xd1\x05\xed\xe9\x40\x47\x3a\xd1\x63\x36\x0f\x2d\x63\xf3\x61\xfc\xb2\xb1\x43\xb1\x40\x47\x3a\x9d\x66\x6f\x84\xb8\x74\x6e\xb2\xc2\x74\xff\x3f\x6e\xf7\xff\x42\x88\x2b\xc6\x8a\xaa\xae\xca\x03\x83\x36\xc7\x42\xe0\x89\x31\x3f\x7f\x89\x1b\x85\x80\xcb\x03\xc3\x7b\xfc\xc7\xa9\x18\x10\x62\x47\xfc\x06\x00\x00\xff\xff\x4d\xd6\x4f\xb3\x04\x18\x00\x00")

func Ds_storeBytes() ([]byte, error) {
	return bindataRead(
		_Ds_store,
		".DS_Store",
	)
}

func Ds_store() (*asset, error) {
	bytes, err := Ds_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1520176529, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _readmeMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x70\x49\x2c\x49\x4c\x4a\x2c\x4e\x55\x48\xc9\x4f\x2e\xcd\x4d\xcd\x2b\x01\x04\x00\x00\xff\xff\x57\x6d\xe5\x94\x13\x00\x00\x00")

func readmeMdBytes() ([]byte, error) {
	return bindataRead(
		_readmeMd,
		"README.MD",
	)
}

func readmeMd() (*asset, error) {
	bytes, err := readmeMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "README.MD", size: 19, mode: os.FileMode(420), modTime: time.Unix(1520139157, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _readme_zhCnMd = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x78\x36\x75\xc3\xb3\xde\x75\x4f\x77\x4d\x7e\x36\xad\xfd\xd9\xc2\xc5\x5c\x5c\xca\x48\x62\x2f\xf6\x37\x3e\x5f\xbe\x1b\x10\x00\x00\xff\xff\xc8\xfa\xce\xd5\x25\x00\x00\x00")

func readme_zhCnMdBytes() ([]byte, error) {
	return bindataRead(
		_readme_zhCnMd,
		"README_zh-CN.MD",
	)
}

func readme_zhCnMd() (*asset, error) {
	bytes, err := readme_zhCnMdBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "README_zh-CN.MD", size: 37, mode: os.FileMode(420), modTime: time.Unix(1520139141, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations201803042326_add_packageDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\xa8\x48\x2f\x88\x2f\x48\x4c\xce\x4e\x4c\x4f\xb5\x06\x04\x00\x00\xff\xff\x50\xaa\xe6\x38\x21\x00\x00\x00")

func migrations201803042326_add_packageDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations201803042326_add_packageDownSql,
		"migrations/201803042326_add_package.down.sql",
	)
}

func migrations201803042326_add_packageDownSql() (*asset, error) {
	bytes, err := migrations201803042326_add_packageDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/201803042326_add_package.down.sql", size: 33, mode: os.FileMode(420), modTime: time.Unix(1520600739, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations201803042326_add_packageUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcf\x41\x4b\xc0\x20\x1c\x05\xf0\xfb\x3e\xc5\xff\xa8\xd0\x61\x0d\x76\xea\x64\xcb\x60\xe4\x5c\x88\x3b\xec\xa4\xa6\x62\x32\xb6\xc4\xb9\xe8\xe3\xc7\x3a\xd4\x82\xe8\xf6\xe0\xfd\x78\xf0\x3a\x41\x89\xa4\x20\xc9\x3d\xa3\xa0\x3f\x42\x52\xc9\xd8\xc5\x04\xaf\x01\x55\x00\x3a\x3a\x0d\x2f\x31\xc4\xad\xa0\xa6\xc6\xc0\x47\x09\x7c\x62\x0c\xc8\x24\x47\xd5\xf3\x4e\xd0\x81\x72\x79\x73\xd2\xb4\x04\xb5\x99\xd5\x6b\xb0\xaf\x26\xa3\xf6\xc2\xbf\xfb\xfd\xed\xc8\xd6\x6b\x78\x37\xf9\x0b\xdd\xd6\x7f\x29\xe7\x77\xfb\x63\x9a\xb6\xc5\xf0\x40\x1f\xc9\xc4\x2e\xee\x48\xce\x14\xef\x94\x29\x1a\xce\x54\xe2\xea\x7f\x2f\xd9\xec\xff\x17\xcf\xa2\x1f\x88\x98\xe1\x89\xce\x80\xce\xa7\xb8\xc2\x77\x9f\x01\x00\x00\xff\xff\xb5\xbd\x64\x2f\x11\x01\x00\x00")

func migrations201803042326_add_packageUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations201803042326_add_packageUpSql,
		"migrations/201803042326_add_package.up.sql",
	)
}

func migrations201803042326_add_packageUpSql() (*asset, error) {
	bytes, err := migrations201803042326_add_packageUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/201803042326_add_package.up.sql", size: 273, mode: os.FileMode(420), modTime: time.Unix(1520687128, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testCleanSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x4e\xce\x48\xcd\x4d\x8c\xcf\xcd\x4c\x2f\x4a\x2c\xc9\xcc\xcf\x2b\xb6\xe6\xc2\xaa\xae\x22\xbd\x20\xbe\x20\x31\x39\x3b\x31\x3d\xd5\x1a\x10\x00\x00\xff\xff\x0b\xdb\xaa\xf9\x49\x00\x00\x00")

func testCleanSqlBytes() ([]byte, error) {
	return bindataRead(
		_testCleanSql,
		"test/clean.sql",
	)
}

func testCleanSql() (*asset, error) {
	bytes, err := testCleanSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/clean.sql", size: 73, mode: os.FileMode(420), modTime: time.Unix(1520683996, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testDataSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x8b\x31\x4e\xc3\x30\x18\x46\xaf\xf2\x6f\x6e\xa5\x34\x8e\xc3\x42\xc3\xc4\xd0\xa1\x12\x2a\x12\x0d\xac\xb1\x65\x5b\xbf\xad\x24\xb6\x95\xd8\x52\xae\xc1\xc8\x0d\x98\xd8\xe1\x3c\xe4\x1e\x28\x42\xc9\xd6\xed\xbd\xa7\xef\x3b\x5f\xae\xa7\x97\x1a\xce\x97\xfa\x19\xf8\x84\x3e\xb4\xc8\x73\x3e\x61\x68\x82\x90\xad\x40\xcd\x77\xdc\x2a\x9e\x01\x0f\x2d\x36\x4e\xf4\x7a\xe5\xd1\xa7\x41\x6e\xa6\xf4\x28\x17\x4e\x41\x89\xa8\x55\x23\xe2\x62\x72\xd0\xab\xed\xe1\xed\xf1\xe9\xf5\x74\x85\x1d\xcb\x80\xf4\x22\x0e\x76\x22\x19\x10\x13\x63\x18\x2b\x4a\xd1\xc6\x5c\x1a\xe1\xd0\x78\x87\xb9\xf5\x54\x1a\xd9\xf9\xa4\xe8\xff\x94\x86\x2e\xa1\x75\x94\x6c\x67\x98\xbf\xbe\xe7\x8f\xcf\xdf\x9f\xf7\xa5\x95\x05\xbb\x3f\x14\x77\x87\xe2\x08\x25\xab\x58\x59\xb1\xe3\xad\xbc\x7f\xf8\x0b\x00\x00\xff\xff\xab\x0d\x7f\x9c\xf5\x00\x00\x00")

func testDataSqlBytes() ([]byte, error) {
	return bindataRead(
		_testDataSql,
		"test/data.sql",
	)
}

func testDataSql() (*asset, error) {
	bytes, err := testDataSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/data.sql", size: 245, mode: os.FileMode(420), modTime: time.Unix(1520687350, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	".DS_Store":                                    Ds_store,
	"README.MD":                                    readmeMd,
	"README_zh-CN.MD":                              readme_zhCnMd,
	"migrations/201803042326_add_package.down.sql": migrations201803042326_add_packageDownSql,
	"migrations/201803042326_add_package.up.sql":   migrations201803042326_add_packageUpSql,
	"test/clean.sql":                               testCleanSql,
	"test/data.sql":                                testDataSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	".DS_Store":       {Ds_store, map[string]*bintree{}},
	"README.MD":       {readmeMd, map[string]*bintree{}},
	"README_zh-CN.MD": {readme_zhCnMd, map[string]*bintree{}},
	"migrations": {nil, map[string]*bintree{
		"201803042326_add_package.down.sql": {migrations201803042326_add_packageDownSql, map[string]*bintree{}},
		"201803042326_add_package.up.sql":   {migrations201803042326_add_packageUpSql, map[string]*bintree{}},
	}},
	"test": {nil, map[string]*bintree{
		"clean.sql": {testCleanSql, map[string]*bintree{}},
		"data.sql":  {testDataSql, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
