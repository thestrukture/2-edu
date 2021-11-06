package assets

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

// tmpl_core_pageend_tmpl reads file data from disk. It returns an error on failure.
func tmpl_core_pageend_tmpl() ([]byte, error) {
	return bindata_read(
		"/data/data/com.termux/files/home/go/src/2-edu/tmpl/core/PageEnd.tmpl",
		"tmpl/core/PageEnd.tmpl",
	)
}

// tmpl_core_pagestart_tmpl reads file data from disk. It returns an error on failure.
func tmpl_core_pagestart_tmpl() ([]byte, error) {
	return bindata_read(
		"/data/data/com.termux/files/home/go/src/2-edu/tmpl/core/PageStart.tmpl",
		"tmpl/core/PageStart.tmpl",
	)
}

// tmpl_ui_feed_tmpl reads file data from disk. It returns an error on failure.
func tmpl_ui_feed_tmpl() ([]byte, error) {
	return bindata_read(
		"/data/data/com.termux/files/home/go/src/2-edu/tmpl/ui/Feed.tmpl",
		"tmpl/ui/Feed.tmpl",
	)
}

// tmpl_ui_navbar_tmpl reads file data from disk. It returns an error on failure.
func tmpl_ui_navbar_tmpl() ([]byte, error) {
	return bindata_read(
		"/data/data/com.termux/files/home/go/src/2-edu/tmpl/ui/NavBar.tmpl",
		"tmpl/ui/NavBar.tmpl",
	)
}

// tmpl_ui_urltool_tmpl reads file data from disk. It returns an error on failure.
func tmpl_ui_urltool_tmpl() ([]byte, error) {
	return bindata_read(
		"/data/data/com.termux/files/home/go/src/2-edu/tmpl/ui/URLTool.tmpl",
		"tmpl/ui/URLTool.tmpl",
	)
}

// web_index_tmpl reads file data from disk. It returns an error on failure.
func web_index_tmpl() ([]byte, error) {
	return bindata_read(
		"/data/data/com.termux/files/home/go/src/2-edu/web/index.tmpl",
		"web/index.tmpl",
	)
}

// web_url_short_tmpl reads file data from disk. It returns an error on failure.
func web_url_short_tmpl() ([]byte, error) {
	return bindata_read(
		"/data/data/com.termux/files/home/go/src/2-edu/web/url-short.tmpl",
		"web/url-short.tmpl",
	)
}

// web_your_404_page_tmpl reads file data from disk. It returns an error on failure.
func web_your_404_page_tmpl() ([]byte, error) {
	return bindata_read(
		"/data/data/com.termux/files/home/go/src/2-edu/web/your-404-page.tmpl",
		"web/your-404-page.tmpl",
	)
}

// web_your_500_page_tmpl reads file data from disk. It returns an error on failure.
func web_your_500_page_tmpl() ([]byte, error) {
	return bindata_read(
		"/data/data/com.termux/files/home/go/src/2-edu/web/your-500-page.tmpl",
		"web/your-500-page.tmpl",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"tmpl/core/PageEnd.tmpl":   tmpl_core_pageend_tmpl,
	"tmpl/core/PageStart.tmpl": tmpl_core_pagestart_tmpl,
	"tmpl/ui/Feed.tmpl":        tmpl_ui_feed_tmpl,
	"tmpl/ui/NavBar.tmpl":      tmpl_ui_navbar_tmpl,
	"tmpl/ui/URLTool.tmpl":     tmpl_ui_urltool_tmpl,
	"web/index.tmpl":           web_index_tmpl,
	"web/url-short.tmpl":       web_url_short_tmpl,
	"web/your-404-page.tmpl":   web_your_404_page_tmpl,
	"web/your-500-page.tmpl":   web_your_500_page_tmpl,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() ([]byte, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"tmpl": &_bintree_t{nil, map[string]*_bintree_t{
		"core": &_bintree_t{nil, map[string]*_bintree_t{
			"PageEnd.tmpl":   &_bintree_t{tmpl_core_pageend_tmpl, map[string]*_bintree_t{}},
			"PageStart.tmpl": &_bintree_t{tmpl_core_pagestart_tmpl, map[string]*_bintree_t{}},
		}},
		"ui": &_bintree_t{nil, map[string]*_bintree_t{
			"Feed.tmpl":    &_bintree_t{tmpl_ui_feed_tmpl, map[string]*_bintree_t{}},
			"NavBar.tmpl":  &_bintree_t{tmpl_ui_navbar_tmpl, map[string]*_bintree_t{}},
			"URLTool.tmpl": &_bintree_t{tmpl_ui_urltool_tmpl, map[string]*_bintree_t{}},
		}},
	}},
	"web": &_bintree_t{nil, map[string]*_bintree_t{
		"index.tmpl":         &_bintree_t{web_index_tmpl, map[string]*_bintree_t{}},
		"url-short.tmpl":     &_bintree_t{web_url_short_tmpl, map[string]*_bintree_t{}},
		"your-404-page.tmpl": &_bintree_t{web_your_404_page_tmpl, map[string]*_bintree_t{}},
		"your-500-page.tmpl": &_bintree_t{web_your_500_page_tmpl, map[string]*_bintree_t{}},
	}},
}}
