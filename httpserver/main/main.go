package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/golang/httpserver/models"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server........")
	http.HandleFunc("/startup", sayHelloName)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	glog.V(2).Info("Access to URI healthz")
	io.WriteString(w, fmt.Sprintf("200"))
	remoteAccessAddress := r.RemoteAddr
	glog.V(2).Info("Remote access address is ", remoteAccessAddress)

	// Get environment variable "VERSION"
	VERSION := os.Getenv("VERSION")
	w.Header().Set("os env variable version", VERSION)

	// Set response header as request header setting
	for rHKey, rHValue := range r.Header {
		for _, value := range rHValue {
			w.Header().Set(rHKey, value)
		}
	}
	w.WriteHeader(200)

}
func sayHelloName(w http.ResponseWriter, r *http.Request) {
	glog.V(2).Info("Access to URI sayHelloName")
	r.ParseForm()
	user := r.URL.Query().Get("user")
	remoteAccessAddress := r.RemoteAddr
	glog.V(2).Info("Remote access address is ", remoteAccessAddress)

	// Get environment variable "VERSION"
	VERSION := os.Getenv("VERSION")
	w.Header().Set("os env variable version", VERSION)

	// Set response header as request header setting
	for rHKey, rHValue := range r.Header {
		for _, value := range rHValue {
			w.Header().Set(rHKey, value)
		}
	}

	// MVC, some models
	employees := `{
		"NBA Players":[
			{"Name":"Kobe","Birthday":"1988-03-14","City":"Kobe Bean Bryant"},
			{"Name":"Durant","Birthday":"1988-09-29","City":"Texas-Austin/United States"},
			{"Name":"LeBron","Birthday":"1984-12-30","City":"t.Vincent-St. Mary HS (OH)/United States"},
			{"Name":"Curry","Birthday":"1988-03-14","City":"Davidson/United States"},
			{"Name":"姚明","Birthday":"1989-09-12","City":"上海"}
		  ]
		}`

	persons := make(map[string][]models.Person)
	err := json.Unmarshal([]byte(employees), &persons)
	if err != nil {
		glog.V(2).Info("JSON String to Struct error ")
	}

	if user != "" {
		for _, person := range persons["NBA Players"] {
			if user == person.Name {
				io.WriteString(w, fmt.Sprintf("Name is %s, City is %s\n", person.GetMethod("Name"), person.GetMethod("City")))
				glog.V(2).Info("Response status code ", 200)
				break
			} else {
				io.WriteString(w, fmt.Sprintf("Name is %s, City is %s\n", user, "Undefined!!"))
				glog.V(2).Info("Response status code ", 200)

				break
			}
		}
	} else {
		io.WriteString(w, fmt.Sprintf("hello stranger\n"))
		glog.V(2).Info("Response status code ", 200)
	}
	w.WriteHeader(200)
}
