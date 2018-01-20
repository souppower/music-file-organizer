package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const targetDir = "BPMSupreme"

func main() {
	if info, err := os.Stat(targetDir); err != nil {
		log.Fatalf("%v", err)
	} else if !info.IsDir() {
		log.Fatalf("%s is not a directory. [error] %v", targetDir, err)
	}

	files, _ := ioutil.ReadDir(targetDir)
	for _, f := range files {
		artistName := strings.Split(f.Name(), " - ")[0]

		if _, err := os.Stat(artistName); os.IsNotExist(err) {
			os.Mkdir(artistName, os.ModePerm)
		}

		oldPath := targetDir + "/" + f.Name()
		newPath := artistName + "/" + f.Name()
		if err := os.Rename(oldPath, newPath); err != nil {
			log.Printf("%v", err)
		}
	}
	fmt.Println("Done!")
}
