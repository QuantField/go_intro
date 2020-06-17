// this is the course on https://tutorialedge.net/golang/creating-restful-api-with-golang/

// We’ll be creating a REST API that allows us to CREATE, READ, UPDATE and DELETE the
// articles on our website. When we talk about CRUD APIs we are referring to an API that
// can handle all of these tasks: Creating, Reading, Updating and Deleting.

/*
Simple web server
inside the current folder web_server type:
go build
an executable file web_server // same name as folder
will be created the run it
./web_server
to stop Ctrl-C
you can check on localhost:3000

*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func main() {

	// populating Articels
	Articles = []Article{
		Article{Title: "Virus outbreak", Desc: "Covid-19", Content: "Research objective"},
		Article{Title: "Advances in Astrophyics", Desc: "Physics", Content: "Age of universe"},
	}
	// going to localhost:3000 is a request for which a handler (response) is needed
	http.HandleFunc("/", homePage) // routing "/" to homePage func

	// if we go to localhost:3000/articles we should see Articls in json format
	http.HandleFunc("/articles", returnAllArticles)

	// log.Fatal not necessary here but good to have in case
	// something goes wrong
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// returning our newly populated Articles variable, encoded in JSON format:
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}
