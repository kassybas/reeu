package common

import (
	"io/ioutil"
	"log"
)

func LoadFile(path string) []byte {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("ERROR: File not found:     %v ", err)
	}
	return f
}
