// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import "2-edu/types"
import "2-edu/api/assets"
import gosweb "github.com/cheikhshift/gos/web"

// Template path
var templateIDFeed = "tmpl/ui/Feed.tmpl"

//
// Renders HTML of template
// Feed with struct gosweb.NoStruct
func Feed(d gosweb.NoStruct) string {
	return netbFeed(d)
}

// Render template with JSON string as
// data.
func netFeed(args ...interface{}) string {

	// Get data from JSON
	var d = netcFeed(args...)
	return netbFeed(d)

}

// template render function
func netbFeed(d gosweb.NoStruct) string {
	localid := templateIDFeed
	name := "Feed"
	defer templateRecovery(name, localid, &d)

	// render and return template result
	return executeTemplate(name, localid, &d)
}

// Unmarshal a json string to the template's struct
// type
func netcFeed(args ...interface{}) (d gosweb.NoStruct) {

	if len(args) > 0 {
		jsonData := args[0].(string)
		err := parseJSON(jsonData, &d)
		if err != nil {
			log.Println("error:", err)
			return
		}
	}

	return
}
