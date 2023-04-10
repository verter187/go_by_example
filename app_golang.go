package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	p("Contains:	", s.Contains("test", "es"))
	p("Count:		", s.Count("test", "t"))
	p("HasPrefix:	", s.HasPrefix("test", "te"))
	p("HasSuffix:	", s.HasSuffix("test", "st"))
	p("Index:		", s.Index("test", "e"))
	p("Join:		", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:		", s.Repeat("a", 5))
	p("Replace:		", s.Replace("foooo", "o", "0", -1))
	p("Replace:		", s.Replace("foooo", "o", "0", 3))
	p("Split:		", s.Split("test str", " "))
	p("ToLower:		", s.ToLower("TEST"))
	p("ToUpper:		", s.ToUpper("test"))
	p()
	p("Len:		", len("test"))
	p("Char:		", "test"[1])
}
