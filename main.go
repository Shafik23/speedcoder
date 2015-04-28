package speedcoder

import (
	"fetcher"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/snippet/", snippetHandler)

	http.HandleFunc("/scripts/", scriptsHandler)
	http.HandleFunc("/", homeHandler)

	// seed the RNG correctly
	rand.Seed(time.Now().UTC().UnixNano())
}

func snippetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fetcher.GetCodeSnippet(r, "json", "python", 200, 300))
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
