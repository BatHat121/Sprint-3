package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
)

var LicenseError int = 0
var LicenseErrorL string = ""

func (d Dir) Lic() bool {
	contents, err := os.Open("LICENSE")
	if err != nil {
		LicenseError = 1
		fmt.Println(err)
	}
	if err == nil {
		fmt.Println("PASS")
		scanner := bufio.NewScanner(contents)
		for scanner.Scan() {
			inside := scanner.Text()
			fmt.Println("Inside: ", inside)
		}
		if err != nil {
			LicenseError = 1
			pc, filename, line, _ := runtime.Caller(1)
			LicenseErrorL = fmt.Sprintf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
			fmt.Println("")
			fmt.Println("FAIL")
			return false
		}
	}
	return true
}

func (sd StrDir) Lic() bool {
	contents, err := os.Open("LICENSE")
	if err == nil {
		fmt.Println("PASS")
		scanner := bufio.NewScanner(contents)
		for scanner.Scan() {
			inside := scanner.Text()
			fmt.Println("Inside: ", inside)
		}
		if err != nil {
			LicenseError = 1
			pc, filename, line, _ := runtime.Caller(1)
			log.Printf("[error] in %s[%s:%d] %v", runtime.FuncForPC(pc).Name(), filename, line, err)
			fmt.Println("")
			fmt.Println("FAIL")
			return false
		}
	}
	return true
}
