package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	target, _ := url.Parse("http://dash-backend:8080")

	proxy := &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {
			pr.SetURL(target)

			path := strings.TrimPrefix(pr.In.URL.Path, "/api")
			if path == "" {
				path = "/"
			}

			pr.Out.URL.Path = path
			pr.Out.Host = pr.Out.URL.Host
		},
	}

	mux := http.NewServeMux()
	mux.Handle("/api/", proxy)
	mux.Handle("/", http.FileServer(http.Dir("./")))

	log.Fatal(http.ListenAndServe(":5555", mux))
}
