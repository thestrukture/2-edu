// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import (
	"2-edu/api/assets"
	"2-edu/types"
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"text/template"

	"github.com/fatih/color"
)

//
// Renders HTML of template
// PageStart with struct types.PageComp
func PageStart(d types.PageComp) string {
	return netbPageStart(d)
}

// recovery function used to log a
// panic.
func templateFNPageStart(localid string, d interface{}) {
	if n := recover(); n != nil {
		color.Red(fmt.Sprintf("Error loading template in path (core/PageStart) : %s", localid))
		// log.Println(n)
		DebugTemplatePath(localid, d)
	}
}

var templateIDPageStart = "tmpl/core/PageStart.tmpl"

// Render template with JSON string as
// data.
func netPageStart(args ...interface{}) string {

	localid := templateIDPageStart
	var d *types.PageComp
	defer templateFNPageStart(localid, d)
	if len(args) > 0 {
		jso := args[0].(string)
		var jsonBlob = []byte(jso)
		err := json.Unmarshal(jsonBlob, d)
		if err != nil {
			return err.Error()
		}
	} else {
		d = &types.PageComp{}
	}

	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("PageStart")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		color.Red(fmt.Sprintf("Error processing template %s", localid))
		DebugTemplatePath(localid, d)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = nil
	output.Reset()
	output = nil
	args = nil
	return outpescaped

}

// alias of template render function.
func bPageStart(d types.PageComp) string {
	return netbPageStart(d)
}

//

// template render function
func netbPageStart(d types.PageComp) string {
	localid := templateIDPageStart
	defer templateFNPageStart(localid, d)
	output := new(bytes.Buffer)

	if _, ok := templateCache.Get(localid); !ok || !Prod {

		body, er := assets.Asset(localid)
		if er != nil {
			return ""
		}
		var localtemplate = template.New("PageStart")
		localtemplate.Funcs(TemplateFuncStore)
		var tmpstr = string(body)
		localtemplate.Parse(tmpstr)
		body = nil
		templateCache.Put(localid, localtemplate)
	}

	erro := templateCache.JGet(localid).Execute(output, d)
	if erro != nil {
		log.Println(erro)
	}
	var outps = output.String()
	var outpescaped = html.UnescapeString(outps)
	d = types.PageComp{}
	output.Reset()
	output = nil
	return outpescaped
}

// Unmarshal a json string to the template's struct
// type
func netcPageStart(args ...interface{}) (d types.PageComp) {
	if len(args) > 0 {
		var jsonBlob = []byte(args[0].(string))
		err := json.Unmarshal(jsonBlob, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	} else {
		d = types.PageComp{}
	}
	return
}

// Create a struct variable of template.
func cPageStart(args ...interface{}) (d types.PageComp) {
	if len(args) > 0 {
		d = netcPageStart(args[0])
	} else {
		d = netcPageStart()
	}
	return
}
