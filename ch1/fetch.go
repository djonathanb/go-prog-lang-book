// Fetch prints the content found at a URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	const (
		httpPrefix = "http://"
	)

	for _, url := range os.Args[1:] {
		var urlToFetch = url
		if !strings.HasPrefix(url, httpPrefix) {
			urlToFetch = fmt.Sprintf("%s%s", httpPrefix, url)
		}

		resp, err := http.Get(urlToFetch)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		w, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("exit: response status %d. finished printing %d characters\n", resp.StatusCode, w)
	}
}
