package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// StringPrompt asks for a string value using the label
func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

func main() {
	name := StringPrompt("Please select a Directory. \n")
	fmt.Printf("Chose directory: %s!\n", name)
	readDir(name)
}

func readDir(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		filename := filepath.Join(path, file.Name())
		if !file.IsDir() {
			dat, err := os.ReadFile(filename)
			checkErr(err)
			fmt.Print(http.DetectContentType(dat) + " : ")
			format := file.ModTime().Format("2006-01-02")
			fileName := file.Name()
			fmt.Println(fileName, format)
			folderPath := filepath.Join(path, format)
			oldPath := filepath.Join(path, fileName)
			err = os.MkdirAll(folderPath, os.ModePerm)
			checkErr(err)
			newPath := filepath.Join(folderPath, fileName)
			fmt.Println(oldPath)
			fmt.Println(newPath)
			err = os.Rename(oldPath, newPath)
			checkErr(err)
		}
	}
}
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
