package main

import (
	"bytes"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
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

//Returns a list of files in a given directory
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

//Converts markdown files to html
func parseMarkdown(fileName string) string {

	extension := strings.SplitN(fileName, ".", 2)[1]

	//Verify file is a markdown file
	if extension != "md" {
		err := fmt.Errorf("file %v is not a markdown file", fileName)
		panic(err.Error())
	}

	mdBytes := []byte(readFile(fileName))
	output := markdown.ToHTML(mdBytes, nil, nil)

	return string(output)

}

func main() {
	//Defines a flag called filePtr
	filePtr := flag.String("file", "first-post.txt", "name of file contents to read")
	dirPtr := flag.String("dir", ".", "directory to pull files from")
	markdownPtr := flag.String("md", "test.md", "markdown file to convert to html")
	//Called after all flags have been defined
	flag.Parse()

	//Parse flag actions
	if activeFlag("dir") {
		files := directorySearch(*dirPtr)

		//Create templates for all files in directory
		for _, file := range files {
			content := readFile(file)

			template := renderTemplate(content)

			fileName := strings.SplitN(file, ".", 2)[0] + ".html"
			saveFile(template, fileName)
		}
		//Exit program successfully
		os.Exit(0)
	} else if activeFlag("md") {

		mdHTML := parseMarkdown(*markdownPtr)
		template := renderTemplate(mdHTML)

		//Save as an html file
		fileName := strings.SplitN(*markdownPtr, ".", 2)[0] + ".html"
		saveFile(template, fileName)

		//Exit program successfully
		os.Exit(0)

	}

	content := readFile(*filePtr)
	template := renderTemplate(content)
	//Gets name of file and changes extension
	fileName := strings.SplitN(*filePtr, ".", 2)[0] + ".html"
	saveFile(template, fileName)

}
