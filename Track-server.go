package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		str := r.URL.EscapedPath()[1:] 
		url, err := base64.StdEncoding.DecodeString(str)
		if err != nil {
			var redirectTo string = string(url)
			fmt.Printf("Redirect to: %s Base64: %s\n", redirectTo, strings.TrimPrefix(r.URL.Path, "/"))
			http.Redirect(w, r, "http://"+redirectTo, http.StatusSeeOther)
		}

		http.Redirect(w, r, "http://"+string(url), http.StatusFound) 
	})
	http.ListenAndServe(":8085", nil)
}
