package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const sourceDirName = "downloaded"
const destDirName = "artists"

const artistNameDelimiter = " - "

func main() {
	if info, err := os.Stat(sourceDirName); err != nil {
		log.Fatalf("%v", err)
	} else if !info.IsDir() {
		log.Fatalf("%s is not a directory. [error] %v", sourceDirName, err)
	}

	files, _ := ioutil.ReadDir(sourceDirName)
	for _, f := range files {
		artistName := strings.Split(f.Name(), artistNameDelimiter)[0]
		artistDirName := destDirName + "/" + artistName
		if _, err := os.Stat(artistDirName); os.IsNotExist(err) {
			os.Mkdir(artistDirName, os.ModePerm)
		}

		oldPath := sourceDirName + "/" + f.Name()
		newPath := artistDirName + "/" + f.Name()
		if err := os.Rename(oldPath, newPath); err != nil {
			log.Printf("%v", err)
		}
	}
	fmt.Println("Done!")
}
