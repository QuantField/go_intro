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
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> This is my website! </h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	//handler := nil
	http.ListenAndServe(":3000", nil)
}
