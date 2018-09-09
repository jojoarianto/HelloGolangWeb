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
	// dashboard search
	http.HandleFunc("/search", search)
	// dashboard search
	http.HandleFunc("/search-get", searchGet)

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

// method for showing search form
func search(w http.ResponseWriter, r *http.Request) {
	var data = map[string]string{
		"title":   "Search",
		"name":    "Irianto",
		"message": "Welcome to go web app",
	}

	// parsing template with 2 return instance template, error (if exist)
	var t, err = template.ParseFiles("index.html")
	if err != nil {
		// shoing error msg on console
		fmt.Println(err.Error())
		return
	}

	// make result of parsing template show on web browser
	t.Execute(w, data)
}

// method to retrieve get value from a form
func searchGet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// call parseform() to parse the raw query and update r.PostForm and r.Form
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err : $v", err)
			return
		}

		email := r.FormValue("email")
		fmt.Fprintf(w, "Email is %s", email)
	default:
		fmt.Fprintf(w, "Sorry, only get and post method are supported")
	}
}
