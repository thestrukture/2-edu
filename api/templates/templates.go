// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import (
	"fmt"
	"html/template"
	"strings"

	gosweb "github.com/cheikhshift/gos/web"

	"2-edu/api/assets"
	methods "2-edu/api/methods"
	"2-edu/types"
)

var TemplateFuncStore template.FuncMap

var templateCache = gosweb.NewTemplateCache()
var WebCache = gosweb.NewCache()

func StoreNetfn() int {
	// List of pipelines linked to each template.
	TemplateFuncStore = template.FuncMap{"a": gosweb.Netadd, "s": gosweb.Netsubs, "m": gosweb.Netmultiply, "d": gosweb.Netdivided, "js": gosweb.Netimportjs, "css": gosweb.Netimportcss, "sd": gosweb.NetsessionDelete, "sr": gosweb.NetsessionRemove, "sc": gosweb.NetsessionKey, "ss": gosweb.NetsessionSet, "sso": gosweb.NetsessionSetInt, "sgo": gosweb.NetsessionGetInt, "sg": gosweb.NetsessionGet, "form": gosweb.Formval, "eq": gosweb.Equalz, "neq": gosweb.Nequalz, "lte": gosweb.Netlt, "IsActive": methods.IsActive, "GetState": methods.GetState, "PageStart": netPageStart, "bPageStart": netbPageStart, "cPageStart": netcPageStart, "PageEnd": netPageEnd, "bPageEnd": netbPageEnd, "cPageEnd": netcPageEnd, "NavBar": netNavBar, "bNavBar": netbNavBar, "cNavBar": netcNavBar, "Feed": netFeed, "bFeed": netbFeed, "cFeed": netcFeed, "URLTool": netURLTool, "bURLTool": netbURLTool, "cURLTool": netcURLTool, "PageComp": types.NewPageComp, "isPageComp": types.CastPageComp, "Feed": types.NewFeed, "isFeed": types.CastFeed}
	return 0
}

var FuncStored = StoreNetfn()

// load a page from your web root by passing the
// path, relative to your web root.
func LoadPage(title string) (*gosweb.Page, error) {

	if lPage, ok := WebCache.Get(title); ok {
		return &lPage, nil
	}

	var nPage = gosweb.Page{}
	if roottitle := (title == "/"); roottitle {
		webbase := "web/"
		fname := fmt.Sprintf("%s%s", webbase, "index.html")
		body, err := assets.Asset(fname)
		if err != nil {
			fname = fmt.Sprintf("%s%s", webbase, "index.tmpl")
			body, err = assets.Asset(fname)
			if err != nil {
				return nil, err
			}
			nPage.Body = body
			WebCache.Put(title, nPage)
			body = nil
			return &nPage, nil
		}
		nPage.Body = body
		nPage.IsResource = true
		WebCache.Put(title, nPage)
		body = nil
		return &nPage, nil

	}

	filename := fmt.Sprintf("web%s.tmpl", title)

	if body, err := assets.Asset(filename); err != nil {
		filename = fmt.Sprintf("web%s.html", title)

		if body, err = assets.Asset(filename); err != nil {
			filename = fmt.Sprintf("web%s", title)

			if body, err = assets.Asset(filename); err != nil {
				return nil, err
			} else {
				if strings.Contains(title, ".tmpl") {
					return nil, nil
				}
				nPage.Body = body
				nPage.IsResource = true
				WebCache.Put(title, nPage)
				body = nil
				return &nPage, nil
			}
		} else {
			nPage.Body = body
			nPage.IsResource = true
			WebCache.Put(title, nPage)
			body = nil
			return &nPage, nil
		}
	} else {
		nPage.Body = body
		WebCache.Put(title, nPage)
		body = nil
		return &nPage, nil
	}

}
