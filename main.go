package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

var reDate = regexp.MustCompile(`\d{8}`)
var reDupe = regexp.MustCompile(`\(1\)`)

func main() {
	// Get current dir, and source + output dir paths
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	sourceDir := filepath.Join(currentDir, "files")
	outputDir := filepath.Join(currentDir, "output")

	// Open source dir & list files
	f, err := os.Open(sourceDir)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	files, err := f.Readdir(0)
	if err != nil {
		panic(err)
	}

	// For each file, extract month, date, and dupe indicator (1)
	// Rename & move into output dir
	for _, file := range files {
		date := reDate.FindString(file.Name())
		dupe := reDupe.FindString(file.Name())
		month := date[2:4]
		year := date[4:]

		newFileName := fmt.Sprintf("Bulletins %s-%s%s.pdf", month, year, dupe)

		err := os.Rename(filepath.Join(sourceDir, file.Name()), filepath.Join(outputDir, newFileName))
		if err != nil {
			panic(err)
		}

		fmt.Printf("%-40s %s\n", file.Name(), newFileName)
	}
}