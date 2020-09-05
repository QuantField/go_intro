/*
sever using template but withoug parsing any data.
hence the nil in ExecuteTemplate method.
*/

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

// init() function is special function, usually used to initilaise
// global varialbles visible at package level, can have as many as
// required in the same file. They are always executed in the order
// they are created.

func init() {
	// Must is an error handler wrapper, if error in not nil
	// will print error and exit with code 1
	tpl = template.Must(template.ParseGlob("public/*.gohtml"))
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("End point Hit: Home page")
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("End point Hit: Contact")
	tpl.ExecuteTemplate(w, "contact.gohtml", nil)
}

func main() {
	fmt.Println("Server 1.0.0")
	http.HandleFunc("/", index)          // registers index to "/"
	http.HandleFunc("/contact", contact) // registers index to "/"
	http.ListenAndServe(":3000", nil)

}
