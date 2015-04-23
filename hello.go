package speedcoder

import (
	"fmt"
	"net/http"
	"fetcher"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fetcher.GetCodeSnippet(r, "json", "python", 200, 300))
}
