package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/", forwardUserServiceRequest)
	fmt.Println("gateway listening to :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func forwardUserServiceRequest(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI
	var url *url.URL
	fmt.Println(path)
	if strings.Contains(path, "api/user") {
		url = getUserServiceURL(path)
	}

	if strings.Contains(path, "api/invoice") {
		url = getInvoiceServiceURL(path)
	}

	rProxy := httputil.NewSingleHostReverseProxy(url)
	rProxy.ServeHTTP(w, r)
}

func getUserServiceURL(path string) *url.URL {
	url, _ := url.Parse("http://127.0.0.1:9001" + path)
	return url
}

func getInvoiceServiceURL(path string) *url.URL {
	url, _ := url.Parse("http://127.0.0.1:9002" + path)
	return url
}
