package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mattbaird/gosaml"
	//"log"
	"io/ioutil"
	"net/http"
	"os"
	//"strings"
	//	"unicode"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/saml", SAMLRequestHandler)
	r.HandleFunc("/proxy", ProxyHandler)

	http.ListenAndServe(":8080", r)

}

func HomeHandler(w http.ResponseWriter, req *http.Request) {

	w.Write([]byte("hihihi"))

}

func SAMLRequestHandler(w http.ResponseWriter, req *http.Request) {
	// Configure the app and account settings
	appSettings := saml.NewAppSettings("http://www.onelogin.net", "issuer")
	accountSettings := saml.NewAccountSettings("cert", "http://www.onelogin.net")

	// Construct an AuthnRequest
	authRequest := saml.NewAuthorizationRequest(*appSettings, *accountSettings)

	// Return a SAML AuthnRequest as a string
	saml, err := authRequest.GetRequest(false)

	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write([]byte(saml))
}

func ProxyHandler(w http.ResponseWriter, r *http.Request) {

	response, err := http.Get(r.URL.RawQuery)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		w.Write([]byte(contents))
	}

}
