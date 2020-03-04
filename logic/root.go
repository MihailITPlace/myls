package logic

import (
	"io/ioutil"
	"log"
	"os"
)

func GetAllFiles() []os.FileInfo {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	return files
}
