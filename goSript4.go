package main

import (
    "fmt"
    "io/ioutil"
    "regexp"
    //"bytes"
)


func main() {
    b, err := ioutil.ReadFile("congress.txt") // just pass the file name
    if err != nil {
        fmt.Print(err)
    }

    str := string(b) // convert content to a 'string'
    
    
    reg := regexp.MustCompile("[[:alpha:]]")
    indexes := reg.FindAllStringIndex(str, -1)
    
    
    //match, _ := regexp.MatchString("[[:alpha:]]", str)
    fmt.Println(indexes)

    //fmt.Println(str) // print the content as a 'string'
}