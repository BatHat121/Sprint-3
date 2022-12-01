package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

var LineErrorCheckL string

func LineNumber(l int) (error, bool) {
	if l > 100 {
		return errors.New("Line Number is too long."), false
	}
	return nil, true
}

var LineCheckError int = 0

func (d Dir) LineCheck() bool {
	fmt.Println("Go File Check is Starting...")
	files, err := os.ReadDir(string(d))
	if err != nil {
		LineCheckError = 1
		fmt.Println(err)
	}
	fmt.Println(files[9].Name())
	var text = ""
	var gofile = ""
	for _, fi := range files {
		if strings.Contains(fi.Name(), ".go") {
			text = "Go File Found."
			gofile = fi.Name()
		} else {
			text = " "
		}
		fmt.Println("file:", gofile, " ", text)
	}
	fil, err := os.Open(gofile)
	if err != nil {
		LineCheckError = 1
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		fmt.Println("Inside:", scanner.Text())
		lineCount := len(scanner.Text())
		fmt.Println("Total Lines: ", lineCount)
		lc, _ := LineNumber(lineCount)
		fmt.Println("Line Check: ")
		fmt.Println("")
		if lc == nil {
			fmt.Println("Pass.")
		}
		if lc != nil {
			LineCheckError++
			pc, filename, line, _ := runtime.Caller(1)
			LineErrorCheckL = fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
			fmt.Println("Fail.")
			return false
		}
	}
	return true
}

func (sd StrDir) LineCheck() bool {
	fmt.Println("Go File Check is Starting...")
	files, err := os.ReadDir(sd.strDirect)
	if err != nil {
		LineCheckError = 1
		fmt.Println(err)
	}
	fmt.Println(files[9].Name())
	var text = ""
	var gofile = ""
	for _, fi := range files {
		if strings.Contains(fi.Name(), ".go") {
			text = "Go File Found."
			gofile = fi.Name()
		} else {
			text = " "
		}
		fmt.Println("File: ", gofile, " ", text)
	}
	fil, err := os.Open(gofile)
	if err != nil {
		LineCheckError = 1
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		fmt.Println("Inside:", scanner.Text())
		lineCount := len(scanner.Text())
		fmt.Println("Total Lines: ", lineCount)
		lc, _ := LineNumber(lineCount)
		fmt.Println("Line Check is Starting...")
		fmt.Println("")
		if lc == nil {
			fmt.Println("")
			fmt.Println("Pass.")
		}
		if lc != nil {
			LineCheckError++
			pc, filename, line, _ := runtime.Caller(1)
			LineErrorCheckL = fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
			fmt.Println("")
			fmt.Println("Fail.")
			return false
		}
	}
	return true
}
