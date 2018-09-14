package main
 
import (
	"log"
	"html/template"
	"os"
)

type Package struct {
	Name string
	NumFuncs int
	NumVars int
}

func main() {
	tmpl, err := template.New("go-web").Parse(`Hello world {{.Name}}`)
	if err != nil {
		log.Fatalf("Parse: %v", err)
	}
	err = tmpl.Execute(os.Stdout, &Package{
		Name: "go-web",
		NumFuncs: 12,
		NumVars: 1200,
	})
	if err != nil {
		log.Fatalf("Execute: %v", err)
	}
}