package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func main() {

	var photoFiles []string
	var e error

	src := "NIKON Z 6_2"
	dst := "NIKON Z 6  " // two spaces in the end are intentional

	photoFiles, e = filepath.Glob("*.[nN][eE][fF]")
	checkError(e)
	for _, pFile := range photoFiles {
		fmt.Print("Processing: " + pFile)

		fileContents, errRead := ioutil.ReadFile(pFile)
		checkError(errRead)

		newContents := strings.Replace(string(fileContents), src, dst, -1)

		errWrite := ioutil.WriteFile(pFile, []byte(newContents), 0)
		checkError(errWrite)

		fmt.Println(" - 100%")
	}

	fmt.Println("Done!")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
