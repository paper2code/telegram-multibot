// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// views/general.tmpl
// views/image.tmpl
// views/music.tmpl
// views/news.tmpl
// views/science.tmpl
// views/social.tmpl
// views/start.tmpl
// views/technology.tmpl
// views/video.tmpl
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

var _viewsGeneralTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\x52\x70\x4f\xcd\x4b\x2d\x4a\xcc\x51\xd0\xe2\x02\x04\x00\x00\xff\xff\x82\xaa\x89\x5d\x0c\x00\x00\x00")

func viewsGeneralTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsGeneralTmpl,
		"views/general.tmpl",
	)
}

func viewsGeneralTmpl() (*asset, error) {
	bytes, err := viewsGeneralTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/general.tmpl", size: 12, mode: os.FileMode(420), modTime: time.Unix(1600582070, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsImageTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\x52\xf0\xcc\x4d\x4c\x4f\x2d\x56\xd0\x02\x04\x00\x00\xff\xff\xc1\xc6\x7d\x19\x0a\x00\x00\x00")

func viewsImageTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsImageTmpl,
		"views/image.tmpl",
	)
}

func viewsImageTmpl() (*asset, error) {
	bytes, err := viewsImageTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/image.tmpl", size: 10, mode: os.FileMode(420), modTime: time.Unix(1600582079, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsMusicTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\x52\xf0\x2d\x2d\xce\x4c\x56\xd0\x02\x04\x00\x00\xff\xff\x95\x4c\x61\x8a\x09\x00\x00\x00")

func viewsMusicTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsMusicTmpl,
		"views/music.tmpl",
	)
}

func viewsMusicTmpl() (*asset, error) {
	bytes, err := viewsMusicTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/music.tmpl", size: 9, mode: os.FileMode(420), modTime: time.Unix(1600582098, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsNewsTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\x52\xf0\x4b\x2d\x2f\x56\xd0\x02\x04\x00\x00\xff\xff\x12\xa8\xc9\x48\x08\x00\x00\x00")

func viewsNewsTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsNewsTmpl,
		"views/news.tmpl",
	)
}

func viewsNewsTmpl() (*asset, error) {
	bytes, err := viewsNewsTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/news.tmpl", size: 8, mode: os.FileMode(420), modTime: time.Unix(1600582115, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsScienceTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\x52\x08\x4e\xce\x4c\xcd\x4b\x4e\x55\xd0\x02\x04\x00\x00\xff\xff\x25\x5d\xf7\x1c\x0b\x00\x00\x00")

func viewsScienceTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsScienceTmpl,
		"views/science.tmpl",
	)
}

func viewsScienceTmpl() (*asset, error) {
	bytes, err := viewsScienceTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/science.tmpl", size: 11, mode: os.FileMode(420), modTime: time.Unix(1600582084, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsSocialTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\x52\x08\xce\x4f\xce\x4c\xcc\x51\xf0\x4d\x4d\xc9\x4c\x54\xd0\x02\x04\x00\x00\xff\xff\x6d\x93\x19\xac\x10\x00\x00\x00")

func viewsSocialTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsSocialTmpl,
		"views/social.tmpl",
	)
}

func viewsSocialTmpl() (*asset, error) {
	bytes, err := viewsSocialTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/social.tmpl", size: 16, mode: os.FileMode(420), modTime: time.Unix(1600582134, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsStartTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\x52\x08\x2e\x49\x2c\x2a\x51\xd0\xe2\xe2\x72\xce\xcf\xcd\x4d\xcc\x4b\x29\xb6\xe2\x02\x04\x00\x00\xff\xff\xf0\x04\x5f\xcb\x15\x00\x00\x00")

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

	info := bindataFileInfo{name: "views/start.tmpl", size: 21, mode: os.FileMode(420), modTime: time.Unix(1600582203, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsTechnologyTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\x0a\x49\x4d\xce\xc8\xcb\xcf\xc9\xcf\x4f\xaf\xd4\x02\x04\x00\x00\xff\xff\x36\x3c\x90\x92\x0d\x00\x00\x00")

func viewsTechnologyTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsTechnologyTmpl,
		"views/technology.tmpl",
	)
}

func viewsTechnologyTmpl() (*asset, error) {
	bytes, err := viewsTechnologyTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/technology.tmpl", size: 13, mode: os.FileMode(420), modTime: time.Unix(1600582067, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsVideoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd2\x52\x08\xcb\x4c\x49\xcd\x2f\x56\xd0\x02\x04\x00\x00\xff\xff\xe2\xe3\x19\x0d\x0a\x00\x00\x00")

func viewsVideoTmplBytes() ([]byte, error) {
	return bindataRead(
		_viewsVideoTmpl,
		"views/video.tmpl",
	)
}

func viewsVideoTmpl() (*asset, error) {
	bytes, err := viewsVideoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/video.tmpl", size: 10, mode: os.FileMode(420), modTime: time.Unix(1600582155, 0)}
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
	"views/general.tmpl":    viewsGeneralTmpl,
	"views/image.tmpl":      viewsImageTmpl,
	"views/music.tmpl":      viewsMusicTmpl,
	"views/news.tmpl":       viewsNewsTmpl,
	"views/science.tmpl":    viewsScienceTmpl,
	"views/social.tmpl":     viewsSocialTmpl,
	"views/start.tmpl":      viewsStartTmpl,
	"views/technology.tmpl": viewsTechnologyTmpl,
	"views/video.tmpl":      viewsVideoTmpl,
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
		"general.tmpl":    &bintree{viewsGeneralTmpl, map[string]*bintree{}},
		"image.tmpl":      &bintree{viewsImageTmpl, map[string]*bintree{}},
		"music.tmpl":      &bintree{viewsMusicTmpl, map[string]*bintree{}},
		"news.tmpl":       &bintree{viewsNewsTmpl, map[string]*bintree{}},
		"science.tmpl":    &bintree{viewsScienceTmpl, map[string]*bintree{}},
		"social.tmpl":     &bintree{viewsSocialTmpl, map[string]*bintree{}},
		"start.tmpl":      &bintree{viewsStartTmpl, map[string]*bintree{}},
		"technology.tmpl": &bintree{viewsTechnologyTmpl, map[string]*bintree{}},
		"video.tmpl":      &bintree{viewsVideoTmpl, map[string]*bintree{}},
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
