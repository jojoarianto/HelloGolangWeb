package main

import (
	"fmt"
	// package for web
	"net/http"
	// package for html template
	"html/template"
)

// main method
func main() {
	// how to write route in go
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello u request : %s\n", r.URL.Path)
	})

	// another route (home route) call home method
	http.HandleFunc("/home", home)
	// dashboard route
	http.HandleFunc("/dashboard", dashboard)

	// log start web
	fmt.Println("Starting web server at http://localhost:8000")
	// start / turn on go web server
	http.ListenAndServe(":8000", nil)
}

// method for home route
func home(w http.ResponseWriter, r *http.Request) {
	var data = map[string]string{
		"title":   "Home",
		"name":    "Irianto",
		"message": "Welcome to go web app",
	}

	// parsing template with 2 return instance template, error (if exist)
	var t, err = template.ParseFiles("template.html")
	if err != nil {
		// shoing error msg on console
		fmt.Println(err.Error())
		return
	}

	// make result of parsing template show on web browser
	t.Execute(w, data)
}

// method for dashboard route
func dashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Dashboard")
}
