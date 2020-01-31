package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	if path == "." {
		path = "testdata"
	}
	if printFiles {
		printDirTreeWithFiles("", out, path)
	} else {
		printDirTree("", out, path)
	}
	return nil
}

func printDirTreeWithFiles(prefix string, out io.Writer, path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			dirString := createOutputStringForDir(file)
			fmt.Fprintln(out, prefix+dirString)
			printDirTreeWithFiles("│	", out, createPath(file, path))
		} else {
			fileNameString := createOutputString(file)
			fmt.Fprintln(out, prefix+fileNameString)
		}
	}
}

func printDirTree(prefix string, out io.Writer, path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			dirString := createOutputStringForDirOnly(file)
			fmt.Fprintln(out, prefix+dirString)
			printDirTree("│	", out, createPath(file, path))
		}
	}
}

func createOutputStringForDir(file os.FileInfo) string {
	return "├───" + file.Name()
}

func createOutputStringForDirOnly(file os.FileInfo) string {
	return "	├───" + file.Name()
}

func createPath(file os.FileInfo, path string) string {
	return path + "/" + file.Name()
}

func createOutputString(file os.FileInfo) string {
	var output = "	├───" + file.Name()
	var size string
	if file.Size() == 0 {
		size = " (empty)"
	} else {
		size = " (" + strconv.FormatInt(int64(file.Size()), 10) + "b)"
	}
	return output + size
}
