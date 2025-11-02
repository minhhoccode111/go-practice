package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	p("contains : ", s.Contains("test", "t"))
	p("count    : ", s.Count("test", "t"))
	p("hasprefix: ", s.HasPrefix("test", "te"))
	p("hassuffix: ", s.HasSuffix("test", "st"))
	p("index    : ", s.Index("test", "e"))
	p("join     : ", s.Join([]string{"test", "tai", "vi", "sao"}, ";"))
	p("repeat   : ", s.Repeat("test;", 5))
	p("replace  : ", s.Replace("foo", "o", "0", -1))
	p("replace  : ", s.Replace("foo", "o", "0", 1))
	p("split    : ", s.Split("a-b-c-d-e", "-"))
	p("tolower  : ", s.ToLower("TEST"))
	p("toupper  : ", s.ToUpper("test"))

	/*
	   $ go run string-functions.go
	   contains :  true
	   count    :  2
	   hasprefix:  true
	   hassuffix:  true
	   index    :  1
	   join     :  test;tai;vi;sao
	   repeat   :  test;test;test;test;test;
	   replace  :  f00
	   replace  :  f0o
	   split    :  [a b c d e]
	   tolower  :  test
	   toupper  :  TEST
	*/
}
