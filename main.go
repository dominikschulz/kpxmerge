package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stderr)
	// check args
	if len(os.Args) < 2 {
		log.Fatalf("Usage: kpxmerge <list of xml files>")
	}
	// resolve arguments
	files := os.Args[1:]
	if len(files) < 2 {
		log.Fatalf("Need at least two files to merge. Got: %v", files)
	}
	// read first file into master db
	db0, err := NewDatabase(files[0])
	if err != nil {
		log.Fatalf("Failed to parse %s: %s", files[0], err)
	}
	log.Printf("Master-Database (%s) has %d entries", files[0], db0.Count())
	// read all other files and merge to master db
	for i := 1; i < len(files); i++ {
		dbN, err := NewDatabase(files[i])
		if err != nil {
			log.Printf("Failed to parse %s: %s. Skipping.", files[i], err)
			continue
		}
		changes, err := db0.Merge(dbN)
		if err != nil {
			log.Printf("Failed to merge %d entries from %s into master db: %s", dbN.Count(), files[i], err)
		} else {
			log.Printf("Merged %d entries from %s (total %d)", changes, files[i], dbN.Count())
		}
	}
	log.Printf("Master-Database has %d entries after merging", db0.Count())
	// write out merged XML to STDOUT
	fmt.Println(db0.String())
}
