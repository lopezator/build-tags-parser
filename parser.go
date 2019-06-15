package parser

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Parse parses all the go files inside the given directory
// and returns a list of build tags inside it.
func Parse(path string) ([]string, error) {
	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

func walk(path string) {
	
}
