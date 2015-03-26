package fetcher

import (
	"fmt"
	"testing"
)

func TestBasic(t *testing.T) {
	fmt.Println("fetcher module: designed to fetch an arbitrary piece of code")
}

func TestFetch(t *testing.T) {
	fmt.Println("Fetching code snippet")
	fmt.Println(GetCodeSnippet("json", "python", 200, 300))
}

