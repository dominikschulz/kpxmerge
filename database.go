package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

type Database struct {
	XMLName xml.Name `xml:"database"`
	Groups  []Group  `xml:"group"`
	Entries []Entry  `xml:"entry"`
}

func NewDatabase(fn string) (Database, error) {
	d := Database{}

	body, err := ioutil.ReadFile(fn)
	if err != nil {
		return d, err
	}

	err = xml.Unmarshal(body, &d)
	if err != nil {
		return d, err
	}

	return d, nil
}

func (d Database) Count() int {
	sum := 0
	for _, g := range d.Groups {
		sum += g.Count()
	}
	sum += len(d.Entries)
	return sum
}

func (d Database) String() string {
	out, err := xml.MarshalIndent(d, "", "  ")
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
	return "<!DOCTYPE KEEPASSX_DATABASE>\n" + string(out)
}

func (d *Database) Merge(o Database) (int, error) {
	var mergedEntries int
	var mergedGroups int
	// merge top-level entries
	d.Entries, mergedEntries = mergeEntries(d.Entries, o.Entries)
	d.Groups, mergedGroups = mergeGroups(d.Groups, o.Groups)
	return mergedEntries + mergedGroups, nil
}
