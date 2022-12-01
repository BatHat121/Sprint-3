package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Validation interface {
	Lic()
	ReadMe()
	GoFileCheck()
	LineCheck()
	CommentCheck()
}

type Dir string
type StrDir struct {
	strDirect strings.Builder
}

func rootdir() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return dir
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var userinput = ""

	for userinput != "val" {
		scanner.Scan()
		userinput = scanner.Text()
		fmt.Println("In: ", userinput)
		if userinput == "val detail" {
			fmt.Println("Details: ")
			fmt.Println("This is the License Validation.")
			var v Validation
			v = Dir(rootdir())
			fmt.Println("This is the Go File Validation.")
			v.GoFileCheck()
			fmt.Println("This is the License File Validation.")
			v.Lic()
			fmt.Println("This is the ReadMe File Validation.")
			v.ReadMe()
			fmt.Println("This is the Comments Validation.")
			fmt.Println("If a comment shows up in this section, it is validated")
			v.CommentCheck()
			fmt.Println("This is the Line Check Validation.")
			v.LineCheck()
		}

		if userinput == "help" {
			fmt.Println("Help: ")
			fmt.Println("")
			fmt.Println("Lic() checks for a License File.")
			fmt.Println("GoFileCheck() checks for a Go File.")
			fmt.Println("ReadMe() checks for a ReadMe File.")
			fmt.Println("CommentCheck() checks for Comments in your File.")
			fmt.Println("LineCheck() checks for number of Characters in a Line.")
			fmt.Println("")
			fmt.Println("At the end of the program there will be a Validation Summary.")
			fmt.Println("Val Detail will reveal summary of errors and where to find them.")
		}

		if userinput == "val" {
			var v Validation
			dir := Dir(rootdir())
			v = dir
			v.Lic()
			v.GoFileCheck()
			v.ReadMe()
			v.CommentCheck()
			v.LineCheck()
			str := StrDir{rootdir()}
			v.Lic()
			v.GoFileCheck()
			v.ReadMe()
			v.CommentCheck()
			v.LineCheck()
		}
	}
}
