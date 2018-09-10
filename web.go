package main

import (
	"fmt"
	// package for web
	"net/http"
	// package for html template
	"html/template"
	// package to connect with mongo db
	"gopkg.in/mgo.v2"
	// package bson
	"gopkg.in/mgo.v2/bson"
)

// model for data user
type User struct {
	// bson property to define name on field from collection db
	Id       bson.ObjectId `bson:"_id,omitempty"`
	Email    string        `bson:"email"`
	Password string        `bson:"password"`
	Phone    string        `bson:"phone_number"`
}

// method to create session to connect on database
func connect() (*mgo.Session, error) {
	// create connection object session
	// parameter is connection string from server mongodb
	var session, err = mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		return nil, err
	}

	return session, nil
}

// main method
func main() {
	// how to write route in go
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello u request : %s\n", r.URL.Path)
		// FindByEmail("nadiramuvidah@gmail.com")
	})

	// another route (home route) call home method
	http.HandleFunc("/home", home)
	// search form
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

// method for showing search form
func search(w http.ResponseWriter, r *http.Request) {
	var data = map[string]string{
		"title":   "Search",
		"name":    "Irianto",
		"message": "Welcome to go web app",
	}

	// parsing template with 2 return instance template, error (if exist)
	var t, err = template.ParseFiles("view/index.html")
	if err != nil {
		// showing error msg on console
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

		// get email value on post request
		email := r.FormValue("email")
		// set var and showing the result of find user
		if email, phone, password := FindByEmail(email); email != "" || phone != "" || password != "" {

			// set data for template result
			var data = map[string]string{
				"title":    "Search Result",
				"email":    email,
				"password": password,
				"phone":    phone,
				// "message": "Welcome to go web app",
			}

			// parsing template and send data to show the result
			var t, err = template.ParseFiles("view/result.html")
			if err != nil {
				// showing error msg on console
				fmt.Println(err.Error())
				return
			}

			// make result of parsing template show on web browser
			t.Execute(w, data)
		}

		// fmt.Fprintf(w, "Email is %s", email)
	default:
		fmt.Fprintf(w, "Sorry, only get and post method are supported")
	}
}

// find method by email
func FindByEmail(email string) (string, string, string) {
	// create session
	session, err := connect()
	if err != nil {
		fmt.Println("Error!", err.Error())
		return "", "", ""
	}
	// close session
	defer session.Close()

	// retrive data from database sibiti, collection users
	collection := session.DB("sibiti").C("users")
	var user User
	selector := bson.M{"email": email}
	err = collection.Find(selector).One(&user)
	if err != nil {
		fmt.Println("Error!", err.Error())
		return "", "", ""
	}

	return user.Email, user.Phone, user.Password
}
