package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func (d Dir) CommentCheck() {
	fmt.Println("Go File Check is Starting...")
	files, err := os.ReadDir(string(d))
	if err != nil {
		fmt.Println(err)
	}
	var text = ""
	var gofile = ""
	for _, fi := range files {
		if strings.Contains(fi.Name(), ".go") {
			text = "Is a Go File."
			gofile = fi.Name()
		} else {
			text = ""
		}
		fmt.Println("File: ", gofile, " ", text)
	}
	fil, err := os.Open(gofile)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		lines := scanner.Text()
	comment: "//"
		if strings ContainsAny(lines, comments) == true {
			fmt.Println("Comment Found.")
		}
		if strings.ContainsAny(lines, comment) == false {
			fmt.Println("No Comments Found.")
		}
	}
}

func (sd StrDir) CommentCheck() {
	fmt.Println("Go File Check is Starting...")
	files, err := os.ReadDir(sd.strDirect)
	if err != nil {
		fmt.Println(err)
	}
	var text = ""
	var gofile = ""
	for _, fi := range files {
		if strings.Contains(fi.Name(), ".go") {
			text = "Is a Go File"
			gofile = fi.Name()
		} else {
			text = ""
		}
		fmt.Println("File: ", gofile " ", text)
	}
	fil, err := os.Open(gofile)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		lines := scanner.Text()
		comment := "//"
		if strings.ContainsAny(lines, comment) == true {
			fmt.Println("Comment Found.")
		}
		if strings.ContainsAny(lines, comment) == false {
			fmt.Println("No Comments Found.")
		}
	}
}
