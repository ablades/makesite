package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

type post struct {
	User    string
	Content string
}

func readFile(name string) string {
	fileContents, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func renderTemplate(content string) *template.Template {
	paths := []string{
		"template.tmpl",
	}
	//will get the template at filename and store it in t. t can then be executed to show the template.
	t := template.Must(template.New("template.tmpl").ParseFiles(paths...))
	err := t.Execute(os.Stdout, post{User: "Audi", Content: content})
	if err != nil {
		panic(err)
	}

	return t
}

func main() {

	content := readFile("first-post.txt")
	t := renderTemplate(content)
	print(t)

}
