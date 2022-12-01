package main

import (
	"fmt"
	"os"
)

var theError int = 0

func (d Dir) ReadMe() bool {
	readme, err := os.Open("README.md")
	if err != nil {
		theError = 1
		fmt.Println(err)
	}
	fil, err := os.ReadFile("README.md")
	if err != nil {
		theError = 1
		fmt.Println(err)
		return false
	}
	fmt.Println(readme)
	fmt.Println(fil)
	return true

}

func (sd StrDir) ReadMe() bool {
	readme, err := os.Open("README.md")
	if err != nil {
		fmt.Println(err)
	}
	fil, err := os.ReadFile("README.md")
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println(readme)
	fmt.Println(fil)
	return true

}
