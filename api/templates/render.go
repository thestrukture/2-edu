// File generated by Gopher Sauce
// DO NOT EDIT!!
package templates

import gosweb "github.com/cheikhshift/gos/web"

// Renders an HTML response with variables provided.
func RenderTemplate(w http.ResponseWriter, p *gosweb.Page) {
	defer func() {
		if n := recover(); n != nil {
			color.Red(fmt.Sprintf("Error loading template in path : web%s.tmpl reason : %s", p.R.URL.Path, n))

			DebugTemplate(w, p.R, fmt.Sprintf("web%s", p.R.URL.Path))
			w.WriteHeader(http.StatusInternalServerError)

			pag, err := LoadPage("/your-500-page")

			if err != nil {
				log.Println(err.Error())
				return
			}

			if pag.IsResource {
				w.Write(pag.Body)
			} else {
				pag.R = p.R
				pag.Session = p.Session
				RenderTemplate(w, pag) ///your-500-page"

			}
		}
	}()

	if _, ok := templateCache.Get(p.R.URL.Path); !ok || !Prod {
		var tmpstr = string(p.Body)
		var localtemplate = template.New(p.R.URL.Path)

		localtemplate.Funcs(TemplateFuncStore)
		localtemplate.Parse(tmpstr)
		templateCache.Put(p.R.URL.Path, localtemplate)
	}

	outp := new(bytes.Buffer)
	err := templateCache.JGet(p.R.URL.Path).Execute(outp, p)
	if err != nil {
		log.Println(err.Error())
		DebugTemplate(w, p.R, fmt.Sprintf("web%s", p.R.URL.Path))
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/html")
		pag, err := LoadPage("/your-500-page")

		if err != nil {
			log.Println(err.Error())
			return
		}
		pag.R = p.R
		pag.Session = p.Session

		if pag.IsResource {
			w.Write(pag.Body)
		} else {
			RenderTemplate(w, pag)

		}
		return
	}

	var outps = outp.String()
	var outpescaped = html.UnescapeString(outps)
	outp = nil
	fmt.Fprintf(w, outpescaped)

}
