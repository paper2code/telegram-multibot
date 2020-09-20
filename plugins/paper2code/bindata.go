// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// views/add.tmpl
// views/feeds.tmpl
// views/home.tmpl
// views/mashup.tmpl
// views/question.tmpl
// views/search.tmpl
// views/start.tmpl
// views/trending.tmpl
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _viewsAddTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xf0\x48\xad\x54\x48\x4c\x49\x51\x50\x04\x04\x00\x00\xff\xff\x4d\xad\x9c\x3e\x0b\x00\x00\x00")

func viewsAddTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsAddTmpl,
		"views/add.tmpl",
	)
}

func viewsAddTmpl() (*asset, error) {
	bytes, err := viewsAddTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/add.tmpl", size: 11, mode: os.FileMode(420), modTime: time.Unix(1600577322, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsFeedsTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xf0\x48\xad\x54\x70\x4b\x4d\x4d\x29\x56\x50\x04\x04\x00\x00\xff\xff\x31\x3e\xec\xf2\x0d\x00\x00\x00")

func viewsFeedsTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsFeedsTmpl,
		"views/feeds.tmpl",
	)
}

func viewsFeedsTmpl() (*asset, error) {
	bytes, err := viewsFeedsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/feeds.tmpl", size: 13, mode: os.FileMode(420), modTime: time.Unix(1600577319, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsHomeTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8c\xcd\x4a\x03\x41\x10\x84\xef\xfd\x14\x7d\x4c\x02\xc9\xe2\xbf\x06\xc4\x17\xf0\x20\x82\x27\x09\x99\xc9\x4e\xef\x6c\xe3\x4c\xf7\x30\x76\x98\xf5\xed\x65\x17\xc1\x83\x17\x2f\x45\x15\x7c\x5f\x6d\x4e\x9a\x02\x1a\x4d\xb6\x01\x38\xb2\xf9\xc4\xfd\x32\x8f\x00\xef\x2c\x89\x85\xf0\xed\xf5\xf9\xb0\x1a\xcd\xca\xbe\xeb\x5a\x6b\x3b\x9a\x7c\x2e\x89\x76\xbd\xe6\x6e\xfd\x8b\x65\x12\x63\x15\xd4\x01\x3d\x9e\x3f\xa9\x1e\x56\x16\xf7\x5d\x37\xd7\x27\x0e\x8f\x17\x97\x57\xd7\x37\xb7\x77\xf7\x0f\x6b\x00\xf7\xe3\x0c\x3c\x51\xd8\x36\x0e\x36\x62\xaf\x81\x1c\x80\x73\x0e\x4a\xa5\xed\xa0\x35\x7b\x33\x0a\x7f\x20\x3c\x25\xed\x3f\x16\x70\x8e\xf2\x65\xa3\xca\xbf\x1c\x6c\x95\xcd\x48\x90\x05\x6d\x24\x7c\x59\x54\x2c\x55\x63\xf5\x39\xb3\x44\x4c\x5e\xe2\xd9\x47\x9a\x9f\xbf\x03\x00\x00\xff\xff\x55\xee\xf9\xec\x1d\x01\x00\x00")

func viewsHomeTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsHomeTmpl,
		"views/home.tmpl",
	)
}

func viewsHomeTmpl() (*asset, error) {
	bytes, err := viewsHomeTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/home.tmpl", size: 285, mode: os.FileMode(420), modTime: time.Unix(1600579827, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsMashupTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xf0\x48\xad\x54\xf0\x4d\x2c\xce\x28\x2d\x50\x50\x04\x04\x00\x00\xff\xff\x38\x67\xd1\x26\x0e\x00\x00\x00")

func viewsMashupTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsMashupTmpl,
		"views/mashup.tmpl",
	)
}

func viewsMashupTmpl() (*asset, error) {
	bytes, err := viewsMashupTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/mashup.tmpl", size: 14, mode: os.FileMode(420), modTime: time.Unix(1600577313, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsQuestionTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\xf0\x48\xad\x54\xc8\x48\x2c\x4b\x55\x28\x2c\x4d\x2d\x2e\xc9\xcc\xcf\x53\xb0\x07\x04\x00\x00\xff\xff\x81\xd5\xd9\xb2\x15\x00\x00\x00")

func viewsQuestionTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsQuestionTmpl,
		"views/question.tmpl",
	)
}

func viewsQuestionTmpl() (*asset, error) {
	bytes, err := viewsQuestionTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/question.tmpl", size: 21, mode: os.FileMode(420), modTime: time.Unix(1600577308, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsSearchTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x08\x4f\xcc\x2b\x51\x28\xc9\x57\x28\x4e\x4d\x2c\x4a\xce\x50\x28\xce\xcf\x4d\x2d\xc9\xc8\xcc\x4b\x57\xb0\x07\x04\x00\x00\xff\xff\xf1\xa3\xb5\xe5\x1c\x00\x00\x00")

func viewsSearchTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsSearchTmpl,
		"views/search.tmpl",
	)
}

func viewsSearchTmpl() (*asset, error) {
	bytes, err := viewsSearchTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/search.tmpl", size: 28, mode: os.FileMode(420), modTime: time.Unix(1600577305, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsStartTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\x08\x48\x2c\x48\x2d\x32\x4a\xce\x4f\x49\x4d\x50\xc8\xcd\x4f\x29\xcd\x49\x55\xc8\xc9\x4f\x4c\x49\x4d\xd1\xd3\xd3\x03\x04\x00\x00\xff\xff\xf9\x54\xff\x7a\x1d\x00\x00\x00")

func viewsStartTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsStartTmpl,
		"views/start.tmpl",
	)
}

func viewsStartTmpl() (*asset, error) {
	bytes, err := viewsStartTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/start.tmpl", size: 29, mode: os.FileMode(420), modTime: time.Unix(1600578132, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsTrendingTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x08\xcf\x48\x2c\x51\xc8\x2c\x56\x28\x29\x4a\xcd\x4b\xc9\xcc\x4b\x57\xb0\x07\x04\x00\x00\xff\xff\x1c\x42\x03\x26\x14\x00\x00\x00")

func viewsTrendingTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsTrendingTmpl,
		"views/trending.tmpl",
	)
}

func viewsTrendingTmpl() (*asset, error) {
	bytes, err := viewsTrendingTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/trending.tmpl", size: 20, mode: os.FileMode(420), modTime: time.Unix(1600577337, 0)}
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
	"views/add.tmpl":      viewsAddTmpl,
	"views/feeds.tmpl":    viewsFeedsTmpl,
	"views/home.tmpl":     viewsHomeTmpl,
	"views/mashup.tmpl":   viewsMashupTmpl,
	"views/question.tmpl": viewsQuestionTmpl,
	"views/search.tmpl":   viewsSearchTmpl,
	"views/start.tmpl":    viewsStartTmpl,
	"views/trending.tmpl": viewsTrendingTmpl,
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
	"views": &bintree{nil, map[string]*bintree{
		"add.tmpl":      &bintree{viewsAddTmpl, map[string]*bintree{}},
		"feeds.tmpl":    &bintree{viewsFeedsTmpl, map[string]*bintree{}},
		"home.tmpl":     &bintree{viewsHomeTmpl, map[string]*bintree{}},
		"mashup.tmpl":   &bintree{viewsMashupTmpl, map[string]*bintree{}},
		"question.tmpl": &bintree{viewsQuestionTmpl, map[string]*bintree{}},
		"search.tmpl":   &bintree{viewsSearchTmpl, map[string]*bintree{}},
		"start.tmpl":    &bintree{viewsStartTmpl, map[string]*bintree{}},
		"trending.tmpl": &bintree{viewsTrendingTmpl, map[string]*bintree{}},
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
