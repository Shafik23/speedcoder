package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	//"github.com/davecgh/go-spew/spew"
)

const (
	codesearch_url_template = "https://searchcode.com/api/codesearch_I/?q=%s lang:%s &loc=%d&loc2=%d"
)

// Given a keyword, a programming language, and a min/max loc, returns
// a random code snippet fitting the input criteria from searchcode.com
func GetCodeSnippet(keyword string, lang string, min_loc int, max_loc int) string {
	url := fmt.Sprintf(codesearch_url_template, keyword, lang, min_loc, max_loc)
	url = strings.Replace(url, " ", "%20", -1)

	fmt.Printf("Language is %s, Keyword is %s\n", lang, keyword)
	fmt.Println("Reading from: ", url)

	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		var f interface{}
		json.Unmarshal(body, &f)
		json_map := f.(map[string]interface{})
		results := json_map["results"].([]interface{})

		result := results[rand.Intn(len(results))].(map[string]interface{})
		snippet_url := strings.Replace(result["url"].(string), "view", "raw", -1)

		resp, err = http.Get(snippet_url)
		defer resp.Body.Close()

		snippet, _ := ioutil.ReadAll(resp.Body)
		return string(snippet)
	} else {
		return "GetCodeSnippet Error - http.Get() could not GET"
	}
}
