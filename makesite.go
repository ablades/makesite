package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
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

	//Run template save contents to buffer
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

func directorySearch(directory string) {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			fmt.Println(file.Name())
		}

	}
}

func main() {
	//Defines a flag called filePtr
	filePtr := flag.String("file", "first-post.txt", "name of file contents to read")

	dirPtr := flag.String("dir", ".", "directory to pull files from")
	//Called after all flags have been defined
	flag.Parse()

	content := readFile(*filePtr)

	template := renderTemplate(content)

	directorySearch(*dirPtr)
	//Gets name of file and changes extension
	fileName := strings.SplitN(*filePtr, ".", 2)[0] + ".html"
	saveFile(template, fileName)
}
