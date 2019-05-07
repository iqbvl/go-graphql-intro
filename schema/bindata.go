// Code generated by go-bindata.
// sources:
// schema/schema.graphql
// schema/type/admin-user.graphql
// DO NOT EDIT!

package schema

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

var _schemaSchemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x3c\x8d\x31\x0a\x02\x31\x10\x45\xfb\x40\xee\xf0\x4b\x05\x4f\x30\x9d\xa5\x85\x85\x88\x95\x58\x2c\xbb\x83\x06\xcc\x44\x27\x93\x62\x09\xde\x5d\x76\x13\xb6\x99\xe2\x31\xef\x3f\xef\x6c\xfe\x30\x2e\x85\x75\x46\x85\x77\x00\xf0\x0e\xd9\x8e\x53\x0c\x72\xcb\xac\xbb\xa4\x13\x2b\xe1\x6a\x1a\xe4\x79\xc0\x98\x8a\x18\xe1\x24\xb6\x27\xdc\xb7\xb7\x87\x77\xde\xfd\x96\xb3\x0e\x9e\x8b\x0d\x16\x92\x54\x00\x9d\xe7\xf1\xc5\x71\x40\x6d\x8d\xef\x52\xa4\x16\x6e\x24\x76\x85\x36\xb9\x8b\xde\xfd\x03\x00\x00\xff\xff\x6e\x41\x8d\x79\xa6\x00\x00\x00")

func schemaSchemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaSchemaGraphql,
		"schema/schema.graphql",
	)
}

func schemaSchemaGraphql() (*asset, error) {
	bytes, err := schemaSchemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/schema.graphql", size: 166, mode: os.FileMode(438), modTime: time.Unix(1557215650, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaTypeAdminUserGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x70\x4c\xc9\xcd\xcc\x0b\x2d\x4e\x2d\x52\xa8\xe6\xe5\x52\x50\x50\x50\xc8\x4c\x51\xc0\x0a\xac\x14\x82\x4b\x8a\x32\xf3\xd2\x21\xaa\x12\x41\xda\xfc\x12\x73\x53\x09\x2b\x2d\x2d\x4e\x2d\xca\xc3\x50\x89\xaa\xa6\x20\xb1\xb8\xb8\x3c\xbf\x28\x05\x9f\x9a\xe2\x92\xc4\x92\xd2\x62\x4c\xbb\x3c\xf3\x4a\x20\x0a\x92\x8b\x52\x13\x4b\x52\x53\x5c\x12\x4b\x52\x71\x9b\x02\x55\xe4\x54\x49\xd0\xd9\x05\x29\xd8\x4c\xc3\xae\x0a\xc5\x38\x98\x0a\x05\x5e\xae\x5a\x40\x00\x00\x00\xff\xff\xb8\x15\x2f\xb1\x64\x01\x00\x00")

func schemaTypeAdminUserGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaTypeAdminUserGraphql,
		"schema/type/admin-user.graphql",
	)
}

func schemaTypeAdminUserGraphql() (*asset, error) {
	bytes, err := schemaTypeAdminUserGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema/type/admin-user.graphql", size: 356, mode: os.FileMode(438), modTime: time.Unix(1557223764, 0)}
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
	"schema/schema.graphql": schemaSchemaGraphql,
	"schema/type/admin-user.graphql": schemaTypeAdminUserGraphql,
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
	"schema": &bintree{nil, map[string]*bintree{
		"schema.graphql": &bintree{schemaSchemaGraphql, map[string]*bintree{}},
		"type": &bintree{nil, map[string]*bintree{
			"admin-user.graphql": &bintree{schemaTypeAdminUserGraphql, map[string]*bintree{}},
		}},
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

