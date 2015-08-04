package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"
)

const kpxDateFormat = "2006-01-02T15:04:05"

type Entry struct {
	Title      string  `xml:"title"`
	Username   string  `xml:"username"`
	Password   string  `xml:"password"`
	Url        string  `xml:"url"`
	Comment    string  `xml:"comment"`
	Icon       int     `xml:"icon"`
	Creation   kpxTime `xml:"creation"`
	LastAccess kpxTime `xml:"lastaccess"`
	LastMod    kpxTime `xml:"lastmod"`
	Expire     kpxTime `xml:"expire"`
}

func (e *Entry) Merge(o Entry) (int, error) {
	if o.LastMod.Time().After(e.LastMod.Time()) {
		log.Printf("%s (challenger) more recent than %s (champion). Challenger wins.\n", o.String(), e.String())
		e.Comment = fmt.Sprintf("%s<br />Conflict (at %s): %s", e.Comment, e.LastMod, e.String())
		e.Password = o.Password
		e.Url = o.Url
		e.Icon = o.Icon
		e.Creation = o.Creation
		e.LastAccess = o.LastAccess
		e.LastMod = kpxTime(time.Now())
		e.Expire = o.Expire
		return 1, nil
	} else if o.LastMod.Time().Before(e.LastMod.Time()) {
		log.Printf("%s (challenger older than %s (champion). Garbage.\n", o.String(), e.String())
		e.Comment = fmt.Sprintf("%s<br />Garbage (at %s): %s", e.Comment, o.LastMod, o.String())
		return 1, nil
	}
	return 0, nil
}

func (e Entry) Ident() string {
	return fmt.Sprintf("%s - %s", e.Title, e.Username)
}

func (e Entry) Equals(o Entry) bool {
	if e.Title == o.Title && e.Username == o.Username {
		return true
	}
	return false
}

func (e Entry) String() string {
	return fmt.Sprintf("Title: %s - Username: %s - Password: %s - URL: %s - Last Modified: %s", e.Title, e.Username, e.Password, e.Url, e.LastMod)
}

type kpxTime time.Time

func (k kpxTime) String() string {
	return k.Time().String()
}

func (t kpxTime) Time() time.Time {
	return (time.Time)(t)
}

func (k *kpxTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	t := (time.Time)(*k)
	e.EncodeElement(t.Format(kpxDateFormat), start)
	return nil
}

func (k *kpxTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(kpxDateFormat, v)
	if err != nil {
		parse = time.Time{}
	}
	*k = kpxTime(parse)
	return nil
}
