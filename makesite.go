package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
)

type post struct {
	User    string
	Content string
}

//Reads in a file and returns contents as string
func readFile(fileName string) string {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

//Renders the template with given content
func renderTemplate(content string) string {
	paths := []string{
		"template.tmpl",
	}

	buff := new(bytes.Buffer)
	//will get the template at filename and store it in t. t can then be executed to show the template.
	t := template.Must(template.New("template.tmpl").ParseFiles(paths...))
	err := t.Execute(buff, post{User: "Audi", Content: content})
	if err != nil {
		panic(err)
	}

	return buff.String()
}

//Saves rendered template to a file
func saveFile(template string, fileName string) bool {
	bytesToWrite := []byte(template)
	err := ioutil.WriteFile(fileName, bytesToWrite, 0644)

	if err != nil {
		return false
	}

	return true
}

func main() {
	content := readFile("first-post.txt")
	t := renderTemplate(content)
	print(t)
	saveFile(t, "first-post.html")

}
