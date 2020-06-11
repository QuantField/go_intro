package main

import (
	"fmt"
	"os"
	"io/ioutil"
)


func check_error(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func main(){

	stringToDump := `[
	{
		"species": "pigeon",
		"decription": "likes to perch on rocks"
	},
	{
		"species":"eagle",
		"description":"bird of prey"
	}
]
`   // creating a file
	f, err := os.Create("dump_test.txt")	
	check_error(err)
	
	l, err := f.WriteString(stringToDump)
	check_error(err)
	fmt.Println(l, "bytes written successfully")
	f.Close()
	
	// reading a file
	dat, err := ioutil.ReadFile("dump_test.txt")
	check_error(err)
	fmt.Println(string(dat))

	

}


