// https://tutorialedge.net/golang/creating-restful-api-with-golang/
// usging routers

// Implementation of a Restful API with mux pacakge

// useful commands for testing

/*
CURL is  a simple but effective command line tool.
Rest implementation test commands :

curl -i -X GET http://rest-api.io/items
curl -i -X GET http://rest-api.io/items/5069b47aa892630aae059584
curl -i -X DELETE http://rest-api.io/items/5069b47aa892630aae059584
curl -i -X POST -H 'Content-Type: application/json' -d '{"name": "New item", "year": "2009"}' http://rest-api.io/items
curl -i -X PUT -H 'Content-Type: application/json' -d '{"name": "Updated item", "year": "2010"}' http://rest-api.io/items/5069b47aa892630aae059584
*/

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database.
// using maps is better than slices, just a little change
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "<h1>Welcome to the HomePage!</h1>")
	tpl, _ := template.ParseFiles("index.gohtml")
	tpl.Execute(w, "John")
	fmt.Println("Endpoint Hit: homePage")

}

// returning our newly populated Articles variable, encoded in JSON format:
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)

	//adding .Methods("POST") to the end of our route to specify that
	//we only want to call this function when the incoming request is a HTTP POST request:
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")

	myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")

	myRouter.HandleFunc("/article/{id}", updateArticle).Methods("PUT")

	// localhost:3000/article/1 will display article with Id=1
	// or using curl:
	// curl http:/localhost:3000/article/1
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)

	//instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		Article{Id: "1", Title: "Virus outbreak", Desc: "Covid-19", Content: "Research objective"},
		Article{Id: "2", Title: "Advances in Astrophyics", Desc: "Physics", Content: "Age of universe"},
	}

	handleRequests()
}

// Create Read Update Delete CRUD

// Reading a single articel
func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}

	}
	//fmt.Fprintf(w, "Key: " + key)
}

// creating new aticle, here adding a new article via a json
/* this is how to do a POST
curl -H "Content-Type: application/json" -X POST -d \
 '{"Id": "3", "Title": "Newly Created Post",  "desc": "The description for my new post",  "content": "my articles content" }' \
 http://localhost:3000/article
*/
// but running the above POST, it is updated in  Articles
func createNewArticle(w http.ResponseWriter, r *http.Request) {
	/// get the body of our POST request
	// unmarshal this into a new Article struct
	// append this to our Articles array.
	reqBody, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(reqBody, &article)

	// update our global Articles array to include
	// our new Article
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
	//fmt.Fprintf(w, "%+v", string(reqBody))
}

// Reading a single articel
// Curl command to delete articel id=1
//curl -X DELETE http://localhost:3000/article/1
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	// we then need to loop through all our articles
	for index, article := range Articles {
		// if our id path parameter matches one of our
		// articles
		if article.Id == id {
			// updates our Articles array to remove the
			// article
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

// curl command to update article id=1
// curl -H "Content-Type: application/json" -X PUT -d \
// '{"Id": "1", "Title": "Newly Updated Post",  "desc": "The description for my updated  post",  "content": "my articles content" }' \
// http://localhost:3000/article/1
func updateArticle(w http.ResponseWriter, r *http.Request) {
	/// get the body of our UPDATE request
	// unmarshal this into a new Article struct
	// update Artices
	reqBody, _ := ioutil.ReadAll(r.Body)

	var updArticle Article
	json.Unmarshal(reqBody, &updArticle)

	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	for index, article := range Articles {
		// if our id path parameter matches one of our
		// articles
		if article.Id == id {
			// updates our Articles array to remove the
			// article
			tmp := Articles[index+1:]
			Articles = append(Articles[:index], updArticle)
			Articles = append(Articles, tmp...)
		}
	}

	json.NewEncoder(w).Encode(updArticle)
	//fmt.Fprintf(w, "%+v", string(reqBody))
}
