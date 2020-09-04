package 

// {{}} represents a place holder in templates



import (
	"log"
	"os"
	"text/template"
)

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func main() {

	tpl, err := template.ParseFiles("myletter.txt")
	checkError(err)

	err = tpl.Execute(os.Stdout, "John")
	checkError(err)

}
