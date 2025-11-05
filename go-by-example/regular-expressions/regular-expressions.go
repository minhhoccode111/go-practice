package main

import (
	"bytes"
	"fmt"
	"regexp"
)

var p = fmt.Println

func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	p(match)

	r, _ := regexp.Compile("p([a-z]+)ch")
	p(r.MatchString("peach"))

	p(r.FindString("peach punch"))
	p("idx:", r.FindStringIndex("peach punch"))
	p(r.FindStringSubmatch("peach punch"))
	p(r.FindAllString("peach punch pinch", -1))
	p("all:", r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	p(r.FindAllString("peach punch pinch", 2))
	p(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	p("regexp:", r)

	p(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	p(string(out))

	/*
	   $ go run regular-expressions.go
	   true
	   true
	   peach
	   idx: [0 5]
	   [peach ea]
	   [peach punch pinch]
	   all: [[0 5 1 3] [6 11 7 9] [12 17 13 15]]
	   [peach punch]
	   true
	   regexp: p([a-z]+)ch
	   a <fruit>
	   a PEACH
	*/
}
