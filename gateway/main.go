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
	fmt.Println("gateway listening to :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func forwardUserServiceRequest(w http.ResponseWriter, r *http.Request) {

	host := strings.Split(r.Host, ":")[0]
	host = "http://" + host
	path := r.RequestURI
	var url *url.URL
	if strings.Contains(path, "api/user") {
		url = getUserServiceURL(host)
	}

	if strings.Contains(path, "api/invoice") {
		url = getInvoiceServiceURL(host)
	}

	rProxy := httputil.NewSingleHostReverseProxy(url)
	rProxy.ServeHTTP(w, r)
}

func getUserServiceURL(host string) *url.URL {
	url, _ := url.Parse(host + ":9001")
	return url
}

func getInvoiceServiceURL(host string) *url.URL {
	url, _ := url.Parse(host + ":9002")
	return url
}
