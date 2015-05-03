package speedcoder

import (
	"appengine"
	"fetcher"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

const (
	min_length = 100
	max_length = 500
)

func init() {
	http.HandleFunc("/snippet/", snippetHandler)

	http.HandleFunc("/scripts/", scriptsHandler)
	http.HandleFunc("/", homeHandler)

	// seed the RNG correctly
	rand.Seed(time.Now().UTC().UnixNano())
}

func snippetHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	lang := r.FormValue("lang")
	keyword := r.FormValue("keyword")

	c.Debugf("(lang, keyword) == (%v, %v)", lang, keyword)

	fmt.Fprint(w, fetcher.GetCodeSnippet(r, keyword, lang, min_length, max_length))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	dumpFileRaw(w, "gui.html")
}

func scriptsHandler(w http.ResponseWriter, r *http.Request) {
	dumpFileRaw(w, r.URL.Path[1:])
}

func dumpFileRaw(w http.ResponseWriter, file string) {
	data, err := ioutil.ReadFile(file)
	if err == nil {
		fmt.Fprintf(w, string(data))
	} else {
		fmt.Fprintf(w, "Could not find page: %q", file)
	}
}
