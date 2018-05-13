package callhandle

import (
	"io/ioutil"
	"net/http"
	"strings"

	"sessiontools"

	"github.com/julienschmidt/httprouter"
)

func StaticSessionFiles(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if sessiontools.CheckSessionUser(r.Header.Get("Sessiontopgoid")) {
		path := r.URL.Path[1:]
		data, err := ioutil.ReadFile(string(path))

		if err == nil {
			var contentType string

			if strings.HasSuffix(path, ".css") {
				contentType = "text/css"
			} else if strings.HasSuffix(path, ".html") {
				contentType = "text/html"
			} else if strings.HasSuffix(path, ".js") {
				contentType = "application/javascript"
			} else if strings.HasSuffix(path, ".png") {
				contentType = "image/png"
			} else if strings.HasSuffix(path, ".svg") {
				contentType = "image/svg+xml"
			} else {
				contentType = "text/plain"
			}

			w.Header().Add("Content Type", contentType)
			w.Write(data)
		} else {
			http.Error(w, "404 page not found - nie ma pliku", 404)
		}
	} else {
		http.Error(w, "404 page not found - nie ma sesji", 404)
	}
}
