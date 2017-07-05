// Code generated by go-bindata.
// sources:
// assets/Makefile
// assets/formula.rb
// DO NOT EDIT!

package command

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

var _assetsMakefile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xc1\x6a\x83\x40\x10\x86\xcf\x9d\xa7\x18\xd0\x83\x1e\xb4\x77\x61\x0f\x85\xb6\x5a\x68\xb5\x28\x2d\xe4\xb8\xc9\x8e\x51\xb2\x51\x71\x55\x02\xc3\xbe\x7b\x30\x68\xf4\x32\xc3\x0c\x7c\xff\xf7\x3b\xe0\xe0\x8f\xbc\x50\x59\x6b\x02\x07\xfe\x3f\xf2\xe2\x2b\x4b\x51\xa0\x69\x64\x67\xaa\x76\x80\x38\xc9\x3f\xbf\xdf\xe2\x02\x05\x80\xa2\x52\x8e\x7a\x88\xf0\x38\xd6\x5a\x01\x84\xbf\x49\x96\x1e\xd6\xf3\x31\x23\x78\x39\xb7\xb7\x13\x06\x4a\x30\x87\xef\x64\x06\x6b\x31\xe8\x26\xe1\x7a\x4b\xb8\xbf\x71\x3d\x69\x92\x86\x60\xd9\x33\x5b\xf5\xc8\x5c\x97\x18\xfe\x19\xea\x53\x79\xa5\x19\x1f\x91\x79\xff\x60\xa6\x46\x59\x8b\xae\xb7\xb6\xf3\x71\xda\x04\xf8\x34\xbf\xee\xac\xf7\x00\x00\x00\xff\xff\x06\x0c\xa5\x85\xeb\x00\x00\x00")

func assetsMakefileBytes() ([]byte, error) {
	return bindataRead(
		_assetsMakefile,
		"assets/Makefile",
	)
}

func assetsMakefile() (*asset, error) {
	bytes, err := assetsMakefileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/Makefile", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsFormulaRb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x93\xcf\x8a\xd4\x40\x10\xc6\xef\x79\x8a\x22\x97\xb9\xa5\x41\xd7\x41\x82\xe0\x61\x60\x5c\x0f\xae\x8b\xb8\x5e\x44\x86\x4e\x77\x4d\x52\x6c\xff\x59\xbb\x3a\xdb\x0e\x49\xde\x5d\x3a\xd9\xdd\x89\x88\xa0\x17\x2f\x9b\x43\x9a\xaa\xef\x57\x1f\x55\x5d\x74\xc0\xef\x3d\x05\x84\x4d\x68\x94\x77\x47\x6a\x37\x85\x32\x92\x19\x86\xa1\xba\x96\xea\x56\xb6\x08\x23\x7c\xa6\x68\x70\x9a\xe0\x0d\xec\x7d\xb0\xbd\x91\x05\x80\x46\x56\x50\x96\x05\x40\xe7\x2d\xde\x65\xb0\xec\x62\xbc\xe3\x5a\x88\x96\x62\xd7\x37\x95\xf2\x56\x0c\x43\x75\xc3\x18\xae\xa4\xc5\x69\x12\x67\xd7\x69\xca\xa5\xf7\x18\x98\xbc\x83\x72\x18\xca\x61\xa8\xbe\x2c\xe1\x34\x95\x59\x2e\x00\xe8\x08\x97\x32\xe8\x24\x03\xd6\xf5\xee\xfa\xa6\x22\x3e\x6c\x2f\x0e\x0d\xc5\xb7\x05\x00\x80\x92\x8c\xf0\xa9\xd9\xcd\xad\xd7\xf5\xee\xe3\xd5\xfe\xfd\xbb\xaf\x9b\xce\x73\x3c\x78\xde\x7c\x9b\xa1\xd4\xa1\x03\x61\x39\x91\x1b\x2d\x9f\x78\xb4\xe4\xda\x34\xaa\x53\x9b\x33\x8d\x52\xf9\x48\xe4\x14\x8e\x68\x95\x98\x6b\x00\xea\x44\x4e\xfb\xc4\x2b\x0b\x2d\xc3\xec\x21\x15\x78\x7e\xe4\xfa\x60\xfe\x71\x70\x11\xd0\xa0\x64\x64\xa1\x7d\x72\xc6\x4b\x2d\xee\x7f\x9f\x5f\x2c\xa9\x0f\x52\x6d\x2f\xaa\x3d\x19\x5c\x9c\xca\xe5\xe2\xf2\xc7\x9d\x7c\xf1\x6a\xfb\x78\x77\x0b\x78\x29\xb9\x5b\x41\x4b\xdf\x86\x5c\xff\xe3\x69\xae\x39\x5a\xa9\xec\x8d\x0c\xc4\x63\xc3\xfa\x89\xe9\x1d\x2d\x08\x1a\xc6\x73\xf2\xd6\xf9\xe4\x96\xbc\xd3\xc5\x4a\x7d\x16\x7b\x78\xf9\x7a\xfb\x97\x8b\xc8\xe4\x7f\xde\x84\xd3\xc5\xfc\x28\x8f\x40\x8e\xa3\x34\x66\x16\x1b\x72\xd5\x43\x9c\xdb\xfb\xf5\xf1\x3d\xd4\x44\xe4\x08\xda\xcf\x3c\x9f\x38\xa2\xfd\x03\x9a\x7f\x3f\x03\x00\x00\xff\xff\xdc\x93\xb1\x25\x30\x04\x00\x00")

func assetsFormulaRbBytes() ([]byte, error) {
	return bindataRead(
		_assetsFormulaRb,
		"assets/formula.rb",
	)
}

func assetsFormulaRb() (*asset, error) {
	bytes, err := assetsFormulaRbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/formula.rb", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"assets/Makefile": assetsMakefile,
	"assets/formula.rb": assetsFormulaRb,
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
	"assets": &bintree{nil, map[string]*bintree{
		"Makefile": &bintree{assetsMakefile, map[string]*bintree{}},
		"formula.rb": &bintree{assetsFormulaRb, map[string]*bintree{}},
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

