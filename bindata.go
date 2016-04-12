// Code generated by go-bindata.
// sources:
// config/emoji.toml
// DO NOT EDIT!

package main

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

var _configEmojiToml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x90\x5d\x6e\xec\x20\x0c\x46\xdf\x67\x15\xa3\x59\xc6\x48\x77\x05\xb7\x3b\xa8\xaa\x91\xe3\x18\x62\x05\x70\x64\x4c\xd3\xee\xbe\x06\xa6\x3f\x8f\x39\x27\x3a\xfe\xc4\x2b\x4a\xce\x6c\xff\xb9\xac\xf5\xed\x12\x44\x33\xd8\xf5\xdf\xf5\x76\x07\xb5\xfb\xed\x72\x90\x0e\x56\x90\x06\x55\x40\xda\x44\x2b\xb9\x5b\x05\xeb\x80\x8b\xc8\x5e\x1d\x2c\x2d\x06\xfe\x98\xa8\x45\x07\xa8\x0d\x19\xd2\xec\xe5\xa5\xa5\xde\x71\xae\x94\xe5\x7d\x06\x03\x6b\x27\x46\xd5\x66\xec\xdc\xd8\xe8\x81\x1b\xe1\xfe\xc8\xa0\xbb\xcb\x4a\xd8\x94\xed\x73\xf8\x24\xd8\x59\xe3\xf9\xc5\x47\x35\x1e\xe4\xe4\x63\x20\x94\x52\xcd\x0f\x1b\x4b\xe9\x65\x88\xbf\x2b\x9f\x41\x2e\x6c\xdf\xbb\x0c\x56\x70\x94\x24\x46\x2e\x71\xa0\x7a\x10\xec\xa4\x3f\x43\x5f\xfe\xb8\xdc\xac\xef\x0d\x04\xd6\x94\x9e\xbf\x7b\x35\x51\x7f\x01\xbf\x1d\x38\x36\x85\x7e\x7c\xca\x22\x67\x48\x9e\x73\xfb\x15\x00\x00\xff\xff\x70\x12\x19\x39\x6c\x01\x00\x00")

func configEmojiTomlBytes() ([]byte, error) {
	return bindataRead(
		_configEmojiToml,
		"config/emoji.toml",
	)
}

func configEmojiToml() (*asset, error) {
	bytes, err := configEmojiTomlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/emoji.toml", size: 364, mode: os.FileMode(420), modTime: time.Unix(1460468452, 0)}
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
	"config/emoji.toml": configEmojiToml,
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
	"config": &bintree{nil, map[string]*bintree{
		"emoji.toml": &bintree{configEmojiToml, map[string]*bintree{}},
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

