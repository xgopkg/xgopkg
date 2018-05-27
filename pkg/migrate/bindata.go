// Code generated by go-bindata. DO NOT EDIT.
// sources:
// .DS_Store
// README.MD
// README_zh-CN.MD
// migrations/201803042326_add_package_table.down.sql
// migrations/201803042326_add_package_table.up.sql
// migrations/201805261449_alter_package_table.down.sql
// migrations/201805261449_alter_package_table.up.sql
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

var _migrations201803042326_add_package_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\xa8\x48\x2f\x88\x2f\x48\x4c\xce\x4e\x4c\x4f\xb5\x06\x04\x00\x00\xff\xff\x50\xaa\xe6\x38\x21\x00\x00\x00")

func migrations201803042326_add_package_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations201803042326_add_package_tableDownSql,
		"migrations/201803042326_add_package_table.down.sql",
	)
}

func migrations201803042326_add_package_tableDownSql() (*asset, error) {
	bytes, err := migrations201803042326_add_package_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/201803042326_add_package_table.down.sql", size: 33, mode: os.FileMode(420), modTime: time.Unix(1521984615, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations201803042326_add_package_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcc\xcb\x4a\xc4\x30\x14\x87\xf1\x7d\x9f\xe2\xbf\x4c\xc0\xc5\x38\xd0\x95\xab\x33\x9d\xa3\x06\xd3\x8c\x84\x53\x61\x56\x21\x24\xa1\x0c\x45\x0d\xbd\x80\x8f\x2f\x08\xde\x40\x5c\x7f\x3f\xbe\xce\x33\x09\x43\xe8\x60\x19\x6f\x63\x0d\x35\xa6\x29\x8e\x05\xaa\x01\x2e\x19\x07\x73\x67\x9c\xa8\xfd\x4e\xc3\x9d\x04\x6e\xb0\x16\x34\xc8\x29\x18\xd7\x79\xee\xd9\xc9\x55\x03\xd4\x69\x0c\x2f\xf1\xb9\xa0\xbb\x27\xaf\xda\x1f\xf8\xb3\x2e\xaf\xdb\x9c\x0a\x9e\xc8\x7f\x90\xeb\xdd\x1f\x26\x97\x25\x7d\x89\x7d\xdb\x6a\x1c\xf9\x96\x06\xfb\xad\xb6\x9a\xe3\x5a\x72\x88\x2b\x8e\x24\x2c\xa6\xe7\x5f\x97\x34\x97\x7f\xfb\xa3\x37\x3d\xf9\x33\x1e\xf8\x0c\x75\xc9\xba\xd1\x37\xef\x01\x00\x00\xff\xff\xa6\x01\x22\x60\x01\x01\x00\x00")

func migrations201803042326_add_package_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations201803042326_add_package_tableUpSql,
		"migrations/201803042326_add_package_table.up.sql",
	)
}

func migrations201803042326_add_package_tableUpSql() (*asset, error) {
	bytes, err := migrations201803042326_add_package_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/201803042326_add_package_table.up.sql", size: 257, mode: os.FileMode(420), modTime: time.Unix(1521984622, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations201805261449_alter_package_tableDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\xa8\x48\x2f\x88\x2f\x48\x4c\xce\x4e\x4c\x4f\x4d\x50\x70\x09\xf2\x0f\x50\x48\x28\x2d\x4e\x2d\x8a\xcf\x4c\x49\xb0\x56\xe0\x02\x0b\x40\x54\x7a\xba\x29\xb8\x46\x78\x06\x87\x04\x43\xf4\x80\x14\x25\x58\x03\x02\x00\x00\xff\xff\xc9\x57\x77\x19\x4b\x00\x00\x00")

func migrations201805261449_alter_package_tableDownSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations201805261449_alter_package_tableDownSql,
		"migrations/201805261449_alter_package_table.down.sql",
	)
}

func migrations201805261449_alter_package_tableDownSql() (*asset, error) {
	bytes, err := migrations201805261449_alter_package_tableDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/201805261449_alter_package_table.down.sql", size: 75, mode: os.FileMode(420), modTime: time.Unix(1527324600, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _migrations201805261449_alter_package_tableUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\xd0\xcb\x6a\xeb\x30\x10\x06\xe0\xbd\x9f\x62\x96\x36\x9c\x45\x0e\x85\x52\x08\x59\x28\xd6\x38\x15\xb5\x95\x56\x1e\x95\x66\x25\x09\x5b\x35\x21\xd4\x31\xbe\xd0\x3c\x7e\xf1\xa5\x50\x08\x69\x57\x12\xe2\xd3\x3f\xfc\xc3\x52\x42\x05\xc4\xb6\x29\x82\xbd\x54\x8d\x69\x5c\x71\x72\x95\xb7\xc0\x38\x07\x3b\x74\xbe\x35\xc7\xd2\x42\xfc\xc8\x54\x78\x77\x1f\x81\xdc\x13\x48\x9d\xa6\xc0\x92\xf1\xa7\x6d\x4e\x95\x29\x7d\x57\xd8\x75\x10\x2b\x64\x84\x4b\x98\x48\x26\x8a\x6f\x22\xa7\x7c\x8e\x1e\xc3\x6c\x18\x00\x00\xd8\x31\x53\x48\x02\xd0\x32\x17\x3b\x89\x1c\x98\xa6\xbd\x11\x32\x56\x98\xa1\xa4\x7f\x33\xbb\x3d\xff\x07\xa8\xdd\x87\xb7\xf0\xca\xd4\x84\xfe\xaf\x56\x57\xaa\x71\x5d\xf7\x79\x6e\xcb\xdf\x55\xd1\x7a\xd7\x7b\xe3\x7a\x0b\x9c\x11\x92\xc8\xf0\x6a\x5c\x53\xfe\x41\x9e\x95\xc8\x98\x3a\xc0\x13\x1e\x20\x1c\x6b\x46\xf3\xbb\x96\xe2\x45\x23\x0c\xc7\xba\xf4\x17\xb3\xf4\x0a\x97\xf3\xa6\x19\xab\x85\xdf\x97\x28\x88\x50\xee\x84\xc4\x8d\xa8\xeb\x33\xdf\x02\xc7\x84\xe9\x94\xa6\xe5\xe4\x48\x9b\xa1\x7f\x7f\x58\x7f\x05\x00\x00\xff\xff\x89\xe9\x75\xb6\xd2\x01\x00\x00")

func migrations201805261449_alter_package_tableUpSqlBytes() ([]byte, error) {
	return bindataRead(
		_migrations201805261449_alter_package_tableUpSql,
		"migrations/201805261449_alter_package_table.up.sql",
	)
}

func migrations201805261449_alter_package_tableUpSql() (*asset, error) {
	bytes, err := migrations201805261449_alter_package_tableUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "migrations/201805261449_alter_package_table.up.sql", size: 466, mode: os.FileMode(420), modTime: time.Unix(1527423236, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testCleanSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x28\x4e\xce\x48\xcd\x4d\x8c\xcf\xcd\x4c\x2f\x4a\x2c\xc9\xcc\xcf\x2b\xb6\xe6\xc2\xaa\xae\x22\xbd\x20\xbe\x20\x31\x39\x3b\x31\x3d\x15\x8f\x8a\xd2\xe2\xd4\x22\x6b\x40\x00\x00\x00\xff\xff\x26\xe6\x5c\xa8\x68\x00\x00\x00")

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

	info := bindataFileInfo{name: "test/clean.sql", size: 104, mode: os.FileMode(420), modTime: time.Unix(1527337296, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testDataSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x90\x3f\x6b\xc3\x30\x14\xc4\xf7\x7e\x0a\x6d\x4a\xc0\x8a\x25\x3b\x4e\x2c\x77\xea\x90\x21\x50\x52\x68\xd2\xae\xd6\x5f\x94\x60\x52\x0b\xbf\x67\xf0\xc7\x2f\x0a\x4d\xb7\x42\x97\x42\x07\x6f\x77\xc7\x1d\xbc\xf7\xdb\x1f\x8e\xbb\xd7\x13\xd9\x1f\x4e\x2f\x44\x4d\xa1\x8f\x5d\x50\x2b\x35\x85\xd8\x46\x6d\x3b\x1d\xbc\x5a\xa8\x8b\x53\x19\x51\xb1\x0b\xed\x87\xbe\xfa\xbb\x86\x7e\x1c\xec\xb7\x73\x1e\x6c\xd2\x23\xf8\xa1\x4d\x03\x35\x46\xa7\xd1\xbb\x56\x63\xca\xed\xe0\xef\x6e\x49\xde\x9f\x9e\xdf\x76\x47\xb2\x10\x19\xa1\xe8\x01\x0b\x9a\x11\x7a\x46\x8c\xd0\xe4\x79\xb8\xe0\x79\x34\x2b\xdb\x5f\xf3\xa9\xbf\x1d\x94\xa7\x0e\xfd\xea\x92\xd8\x05\x9a\x51\xa1\x79\x6d\xca\x8a\x33\xee\x9d\x60\x6b\x5d\x49\x66\xb8\x95\xac\xda\xf2\x4d\x5d\x17\xc6\x09\x69\xd2\xa4\xe0\xa2\x66\xbc\x64\x5c\x92\x42\x34\xa2\x68\x84\xfc\x29\x5e\x3e\x3e\xfc\x86\xc6\x1f\x70\xb8\x3d\x56\xce\x10\x00\xd7\x33\x04\xc0\x6a\x86\x00\xb8\x99\x21\x00\x6e\xff\x09\x84\xcf\x00\x00\x00\xff\xff\x85\x84\x39\x35\x26\x06\x00\x00")

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

	info := bindataFileInfo{name: "test/data.sql", size: 1574, mode: os.FileMode(420), modTime: time.Unix(1527423171, 0)}
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
	".DS_Store":                                            Ds_store,
	"README.MD":                                            readmeMd,
	"README_zh-CN.MD":                                      readme_zhCnMd,
	"migrations/201803042326_add_package_table.down.sql":   migrations201803042326_add_package_tableDownSql,
	"migrations/201803042326_add_package_table.up.sql":     migrations201803042326_add_package_tableUpSql,
	"migrations/201805261449_alter_package_table.down.sql": migrations201805261449_alter_package_tableDownSql,
	"migrations/201805261449_alter_package_table.up.sql":   migrations201805261449_alter_package_tableUpSql,
	"test/clean.sql":                                       testCleanSql,
	"test/data.sql":                                        testDataSql,
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
		"201803042326_add_package_table.down.sql":   {migrations201803042326_add_package_tableDownSql, map[string]*bintree{}},
		"201803042326_add_package_table.up.sql":     {migrations201803042326_add_package_tableUpSql, map[string]*bintree{}},
		"201805261449_alter_package_table.down.sql": {migrations201805261449_alter_package_tableDownSql, map[string]*bintree{}},
		"201805261449_alter_package_table.up.sql":   {migrations201805261449_alter_package_tableUpSql, map[string]*bintree{}},
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
