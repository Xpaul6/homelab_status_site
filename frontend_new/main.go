package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No env file provided")
	}
}

func main() {
	target, _ := url.Parse(os.Getenv("BACKEND_URL"))

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
	mux.Handle("/", http.FileServer(http.Dir("./static/")))

	log.Fatal(http.ListenAndServe(":5555", mux))
}
