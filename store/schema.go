package store

//
// Welcome to our badass cms.
//

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// CreateSchema Remove/create schema from scratch.
func CreateSchema() {
	log.Println("-> Create DB schema.")
	file, err := ioutil.ReadFile(schemaFilePath())

	if err != nil {
		panic(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		_, err := DB.Exec(request)

		if err != nil {
			panic(err)
		}
	}
}

func schemaFilePath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/sql/schema.sql", dir)
}
