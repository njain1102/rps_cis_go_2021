package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))

	fmt.Println(r.FindStringIndex("peach punch"))

	fmt.Println(r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))

	matched, err := regexp.MatchString(`a.b`, "aaxbb")
	fmt.Println(matched) // true
	fmt.Println(err)
	matched, _ = regexp.MatchString(`^a.b$`, "aaxbb")
	fmt.Println(matched)
	matched, _ = regexp.MatchString(`^a.b$`, "aaxbb")
	fmt.Println(matched)

	//re1, err := regexp.Compile(`regexp`) // error if regexp invalid
	//re2 := regexp.MustCompile(`regexp`)  // panic if regexp invalid

	re := regexp.MustCompile(`foo.?`)
	fmt.Printf("%q\n", re.FindString("seafood fool")) // "food"
	fmt.Printf("%q\n", re.FindString("meat"))

	re = regexp.MustCompile(`ab?`)
	fmt.Println(re.FindStringIndex("tablett"))    // [1 3]
	fmt.Println(re.FindStringIndex("foo") == nil) // true

	re = regexp.MustCompile(`a.`)
	fmt.Printf("%q\n", re.FindAllString("paranormal", -1)) // ["ar" "an" "al"]
	fmt.Printf("%q\n", re.FindAllString("paranormal", 2))  // ["ar" "an"]
	fmt.Printf("%q\n", re.FindAllString("graal", -1))      // ["aa"]
	fmt.Printf("%q\n", re.FindAllString("none", -1))       // [] (nil slice)

	re = regexp.MustCompile(`ab*`)
	fmt.Printf("%q\n", re.ReplaceAllString("-a-abb-", "T")) // "-T-T-"

	a := regexp.MustCompile(`a`)
	fmt.Printf("%q\n", a.Split("banana", -1)) // ["b" "n" "n" ""]
	fmt.Printf("%q\n", a.Split("banana", 0))  // [] (nil slice)
	fmt.Printf("%q\n", a.Split("banana", 1))  // ["banana"]
	fmt.Printf("%q\n", a.Split("banana", 2))  // ["b" "nana"]

	zp := regexp.MustCompile(`z+`)
	fmt.Printf("%q\n", zp.Split("pizza", -1)) // ["pi" "a"]
	fmt.Printf("%q\n", zp.Split("pizza", 0))  // [] (nil slice)
	fmt.Printf("%q\n", zp.Split("pizza", 1))  // ["pizza"]
	fmt.Printf("%q\n", zp.Split("pizza", 2))  // ["pi" "a"]





}