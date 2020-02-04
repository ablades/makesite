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

	//Run template save contents to buffer(temporary storage)
	err := t.Execute(buff, post{User: "Audi", Content: content})
	if err != nil {
		panic(err)
	}

	return buff.String()
}

//Saves rendered template to a file
func saveFile(buffer string, fileName string) bool {
	bytesToWrite := []byte(buffer)
	err := ioutil.WriteFile(fileName, bytesToWrite, 0644)

	if err != nil {
		return false
	}

	return true
}

func directorySearch(directory string) []string {
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	var textFiles []string

	//Check files in directory for a txt extension
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".txt" {
			fmt.Println(file.Name())
			textFiles = append(textFiles, file.Name())
		}
	}

	return textFiles
}

//Checks if a specified flag is active
func activeFlag(name string) bool {
	active := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			active = true
		}
	})

	return active
}

func main() {
	//Defines a flag called filePtr
	filePtr := flag.String("file", "first-post.txt", "name of file contents to read")
	dirPtr := flag.String("dir", ".", "directory to pull files from")
	//Called after all flags have been defined
	flag.Parse()

	//Parse given directory
	if activeFlag("dir") {
		files := directorySearch(*dirPtr)

		//Create templates for all files in directory
		for _, file := range files {
			content := readFile(file)

			template := renderTemplate(content)

			fileName := strings.SplitN(file, ".", 2)[0] + ".html"
			saveFile(template, fileName)
		}

	} else {
		content := readFile(*filePtr)

		template := renderTemplate(content)
		//Gets name of file and changes extension
		fileName := strings.SplitN(*filePtr, ".", 2)[0] + ".html"
		saveFile(template, fileName)
	}

}
