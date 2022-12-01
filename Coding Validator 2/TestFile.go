package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestLineCheck(t *testing.T) {
	var v Validation
	dir := Dir(rootdir())
	v = dir
	lines := v.LineCheck()

	if lines != true {
		t.Error("Line is over 100 characters.")
	} else {
		t.Log("Line is under 100 characters.")
	}
}

func TestReadMe(t *testing.T) {
	var strPath string
	strPath = rootdir() + "\\" + "README.md"
	var v Validation
	v = Dir(strPath)
	file := v.ReadMe()

	if file != true {
		t.Error("ReadMe File Not Found.")
	} else {
		t.Log("ReadMe File is Found.")
	}
}

func TestCommentCheck(t *testing.T) {
	var v Validation
	dir := Dir(rootdir())
	v = dir
	Comment := v.CommentCheck()

	if Comment != true {
		t.Error("No Comments.")
	} else {
		t.Log("Comments Found.")
	}

}
func TestLicense(t *testing.T) {
	var strPath string
	strPath = rootdir() + "\\" + "LICENSE"
	var v Validation
	v = Dir(strPath)
	file := v.Lic()

	if file != true {
		t.Error("No License File.")
	} else {
		t.Log("License File Found.")
	}
}

func TestRootdir(t *testing.T) {
	result := rootdir()
	currentdir, err := os.Getwd()
	if result != currentdir {
		fmt.Println(err)
		t.Error("Select Current Directory.")
	} else {
		t.Log("In Current Directory.")
	}
}

func TestReadMeStrDirect(t *testing.T) {
	var strPath string
	strPath = rootdir() + "\\" + "README.md"
	var v Validation
	v = StrDir{strPath}
	file := v.ReadMe()

	if file != true {
		t.Error("No ReadMe File.")
	} else {
		t.Log("ReadMe File Found.")
	}
}
func TestLicStrDirect(t *testing.T) {
	var strPath string
	strPath = rootdir() + "\\" + "LICENSE"
	var v Validation
	v = StrDir{strPath}
	file := v.Lic()

	if file != true {
		t.Error("No License File.")
	} else {
		t.Log("License File Found.")
	}
}

func TestLineCheckStrDir(t *testing.T) {
	var v Validation
	str := StrDir{rootdir()}
	v = str
	lines := v.LineCheck()

	if lines != true {
		t.Error("Lines are too long.")
	} else {
		t.Log("Lines are in range.")
	}
}

func TestCommentCheckStrDir(t *testing.T) {
	var v Validation
	str := StrDir{rootdir()}
	v = str
	comment := v.CommentCheck()

	if comment != true {
		t.Error("No Comment Found.")
	} else {
		t.Log("Comment Found.")
	}
}

func TestLineNumber(t *testing.T) {
	fil, err := os.Open("val.go")
	if err != nil {
		t.Error("Val.go File Will NOT Open.")
	}
	scanner := bufio.NewScanner(fil)
	for scanner.Scan() {
		lines := scanner.Text()
		line := len(lines)
		l, _ := LineNumber(line)

		if l != nil {
			t.Error("Lines are too long.")
		} else {
			t.Log("Lines are in range.")
		}
	}
}
