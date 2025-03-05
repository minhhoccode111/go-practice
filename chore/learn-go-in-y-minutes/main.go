// single line comment

//go:build prod || dev || test

/*
a build tag is a line comment starting with //go:build and can be executed by go
build -tags="foo bar" command
build tags are placed before the package clause near or at the top of the file
followed by a blank line or other line comments
*/

/*
multi-line comment
*/

// a package clause starts every source file
// main is a special name declaring an executable rather than a library
package main

// import declaration declares library packages referenced in this file
import (
	"fmt"              // a package in the go standard library
	"io"               // implements some i/o utility functions
	m "math"           // math library with local alias m
	"net/http"         // yes, a web server
	_ "net/http/pprof" // profiling library imported only for side effects
	"os"               // os functions like working with the file system
	"strconv"          // string conversions
)

// a function definition. main is special. it is the entry point for the
// executable program. love it or hate it, go uses brace brackets
func main() {
	// println outputs a line to stdout
	// it comes from the package fmt
	fmt.Println("Hello, World")

	// call another function within this package
	beyondHello()
}

// functions have parameters in parentheses
// if there are no parameters, empty parentheses are still required
func beyondHello() {
	var x int // variable declaration, variables must be declared before use
	x = 3     // variable assignment

	// "short" declarations use := to infer the type, declare, and assign
	y := 4
	sum, prod := learnMultiple(x, y)          // function return 2 values
	fmt.Println("sum: ", sum, "prod: ", prod) // simple output
	learnTypes()                              // < y minutes, learn more

}

/*
<- multiline comment
functions can have parameters and (multiple!) return values
here x, y are the arguments and sum, prod is the signature (what's returned)
note that x and sum receive the type `int`
*/
func learnMultiple(x, y int) (sum, prod int) {
	return x + y, x * y // return two values
}

// some built-in types and literals
func learnTypes() {
	// short declaration usually gives you what you want
	str1 := "Learn go!" // string type
	str2 := `A "raw" string literal
	can include line breaks.` // same string type

	// non-ASCII literal. go source is UTF-8
	g := "Σ" // rune type, an alias for int32, holds a unicode code point

	f := 3.14159 // float64, an IEEE-754 64-bit floating point number
	c := 3 + 4i  // complex128, represented internally with two float64's

	// var syntax with initializers
	var u uint = 7 // unsigned, but implementation dependent size as with int
	var pi float32 = 22. / 7

	// conversion syntax with a short declaration
	n := byte('\n') // byte is an alias for uint8

	// arrays have size fixed at compile time
	var a4 [4]int                    // an array of 4 ints, initialized to all 0
	a5 := [...]int{3, 1, 5, 10, 100} // an array initialized with a fixed size of
	// five elements, with values 3, 1, 5, 10, 100

	// arrays have values semantics
	a4_cpy := a4                    // a4_cpy is a copy of a4, two separate instances
	a4_cpy[0] = 25                  // only a4_cpy is changes, a4 stays the same
	fmt.Println(a4_cpy[0] == a4[0]) // false

	// slices have dynamic size. arrays and slices each have advantages but use
	// cases for slices are much more common
	s3 := []int{4, 5, 9}    // compare to a5. no ellipsis here
	s4 := make([]int, 4)    // allocates slice of 4 ints, initialized to all 0
	var d2 [][]float64      // declaration only, nothing allocated here
	bs := []byte("a slice") // type comversion syntax

	// slices (as well as maps and channels) have reference semantics
	s3_cpy := s3                    // both variables point to the same instance
	s3_cpy[0] = 0                   // which means both are updated
	fmt.Println(s3_cpy[0] == s3[0]) // true

	// because they are dynamic, slices can be appended to on-demand
	// to append elements to a slice, the built-in append() function is used
	// first argument is a slice to which we are appending. commonly, the slice
	// variable is updated in place, as in example below
	s := []int{1, 2, 3}
	s = append(s, 4, 5, 6) // added 3 elements, slice not has length of 6
	fmt.Println(s)         // updated slice is not [1 2 3 4 5 6]

	// to append another slice, instead of list of atomic elements we can pass a
	//  reference to a slice of a slice literal like this, with a trailing
	// ellipsis, meaning take a slice and unpack its elements, appending them to
	// slice s
	s = append(s, []int{7, 8, 9}...) // second argument is a slice literal
	fmt.Println(s)                   // updated slice is not [1 2 3 4 5 6 7 8 9]

	p, q := learnMemory() // declares p, q to be type pointer to int
	fmt.Println(*p, *q)   // * follows a pointer, this prints two ints

	// maps are a dynamically growable associative array type, like the hash or
	// dictionary types of some other languages
	m := map[string]int{"three": 3, "four": 4}
	m["one"] = 1
	// looking up a missing key returns the zero value, which is 0 in this case,
	// since it's a map[string]int
	fmt.Println(m["key not present"]) // 0
	// check if a key is present in the make like this:
	if _, ok := m["one"]; ok {
		// do something
	}

	// unused variables are an error in Go
	// the underscore lets you "use" a variable but discard its value
	_, _, _, _, _, _, _, _, _, _ = str1, str2, g, f, u, pi, n, a5, s4, bs
	// usually you use it to ignore one of the return values of a function
	// for example, in a quick and dirty script you might ignore the error value
	// returned from os.Create, and expect that the file will always be created
	file, _ := os.Create("output.txt")
	fmt.Fprint(file, "This is how you write to a file, by the way")
	file.Close()

	// output of course counts as using a variable
	fmt.Println(s, c, a4, s3, d2, m)
	learnFlowControl() // back in the flow
}

// it is possible, unlike in many other languages for functions in go to have
// named return values
// assigning a name to the type begin returned in the function declaration line
// allows us to easily return from multiple points in a function as well as to
// only use the return keyword, without anything further
func learnNamedReturn(x, y int) (z int) {
	z = x * y
	return // z is implicit here, because we named it earlier
}

// go is fully garbage collected, it has pointers but no pointer arithmetic
// you can make a mistake with a nil poiner, but not by incrementing a pointer
// unlike in c/cpp taking and returning an address of a local variable is also
// safe
func learnMemory() (p, q *int) {
	// named return values p and q have type pointer to int
	p = new(int) // built-in function new allocates memory
	// the allocated int slice is initialized to 0, p is no longer nil
	s := make([]int, 20) // allocate 20 ints as a single block of memory
	s[3] = 7             // assign one of them
	r := -2              // declare another local variable
	return &s[3], &r     // & takes the address of an object
}

// use the aliased math library (see import, above)
func expensiveComputation() float64 {
	return m.Exp(10)
}

func learnFlowControl() {
	// if statements require brace brackets, and do not require parentheses
	if true {
		fmt.Println("told ya")
	}
	// formatting is standardized by the command line command "go fmt"
	if false {
		// pout
	} else {
		// gloat
	}
	// use switch in preference to chained if statements
	x := 42.0
	switch x {
	case 0:
	case 1, 2: // can have multiple matches on one case
	case 42:
		// cases don't "fall through"
		/*
		   there is a `fallthrough` keyword however, see:
		   https://go.dev/wiki/Switch#fall-through
		*/
	case 43:
	// unreached
	default:
		// default case is optional
	}

	// type switch allow switching on the type of something instead of value
	var data interface{}
	data = ""
	switch c := data.(type) {
	case string:
		fmt.Println(c, "is a string")
	case int64:
		fmt.Printf("%d is an int64\n", c)
	default:
		// all other cases
	}

	// like if, for doesn't use parens either
	// variables declared in for and if are local to their scope
	for x := 0; x < 3; x++ { // ++ is a statement
		fmt.Println("iteration", x)
	}
	// x == 42 here

	// for is the only loop statement in Go, but it has alternate forms
	for { // infinite loop
		break    // just kidding
		continue //  unreached
	}

	// you can use range to iterate over an array, a slice, a string, a map, or
	// a channel
	// range returns one (channel) or two values (array, slice, string and map)
	for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
		// for each pair in the map, print key and value
		fmt.Printf("key=%s, value=%d\n", key, value)
	}
	// if you only need the value, use the underscore as the key
	for _, name := range []string{"Bob", "Bill", "Joe"} {
		fmt.Printf("Hello, %s\n", name)
	}

	// as with for, := in an if statement means to declare and assign y first,
	// then test y > x
	if y := expensiveComputation(); y > x {
		x = y
	}
	// function literals are closures
	xBig := func() bool {
		return x > 10000 // references x declared above switch statement
	}
	x = 99999
	fmt.Println("xBig: ", xBig()) // true
	x = 1.3e3                     // this makes x == 1300
	fmt.Println("xBig: ", xBig()) // false now

	// what's more is function literals may be defined and called inline, acting
	// as an argument to function, as long as:
	// a) function literal is called immediately (),
	// b) result type matches expected type of argument
	fmt.Println("Add + double two numbers: ",
		func(a, b int) int {
			return (a + b) * 2
		}(10, 2),
	) // called with args 10 and 2
	// => Add + double two numbers: 24

	// when you need it, you'll love it
	goto love
love:
	learnFunctionFactory() // func return func is fun(3)(3)
	learnDefer()           // a quick detour to an important keyword
	learnInterfaces()      // good stuff coming up

}

func learnFunctionFactory() {
	// next two are equivalent, with second being more practical
	fmt.Println(sentenceFactory("summer")("A beautiful", "day!"))

	d := sentenceFactory("summer")
	fmt.Println(d("A beautiful", "day!"))
	fmt.Println(d("A lazy", "afternoon!"))
}

func sentenceFactory(s string) func(before, after string) string {
	return func(before, after string) string {
		return fmt.Sprintf("%s %s %s", before, s, after) // new string
	}
}

func learnDefer() (ok bool) {
	// a defer statement pushes a function call onto a list. the list of saved
	// calls is executed after the surrounding function returns
	defer fmt.Println("deferred statements execute in reverse (LIFO) order (stack).")
	defer fmt.Println("\nThis line is begin printed first because")
	// defer is commonly used to close a file, so the function closing the file
	// stays near the function opening the file
	return true
}

// define Stringer as an interface type with one method, String
type Stringer interface {
	String() string
}

// define pair as a struct with two fields, ints named x and y
type pair struct {
	x, y int
}

// define a method on type pair. pair now implements Stringer because pair has
// defined all the method in the interface
func (p pair) String() string { // p is called the "receiver"
	// Sprintf is another public function in package fmt
	// dot syntax references fields of p
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func learnInterfaces() {
	// brace syntax is a "struct literal". it evaluates to an initialized struct
	// the := syntax declares and initializes p to this struct
	p := pair{3, 4}
	fmt.Println(p.String()) // call String method of p, of type pair
	var i Stringer          // declare i of interface type Stringer
	i = p                   // valid because pair implements Stringer
	// call String method of i, of type Stringer. output same as above
	fmt.Println(i.String())

	// functions in the fmt package call the String method to ask an object for
	// a printable representation of itself
	fmt.Println(p) // output same as above. println calls String method
	fmt.Println(i) // output same as above

	learnVariadicParams("great", "learning", "here!")
}

// functions can have variadic parameters
func learnVariadicParams(myStrings ...any) { // any is an alias for interface{}
	// iterate each value of the variadic
	// the underscore here is ignoring the index argument of the array
	for _, param := range myStrings {
		fmt.Println("param: ", param)
	}

	// pass variadic value as a variadic parameter
	fmt.Println("params: ", fmt.Sprintln(myStrings...))

	learnErrorHandling()
}

func learnErrorHandling() {
	// ", ok" idiom used to tell if something worked or not
	m := map[int]string{3: "three", 4: "four"}
	if x, ok := m[1]; !ok { // ok will be false because 1 is not in the map
		fmt.Println("no one there")
	} else {
		fmt.Print(x) // x would be the value, if it were in the map
	}
	// an error value commuinicates not just "ok" but more about the problem
	if _, err := strconv.Atoi("non-int"); err != nil { // `_` discard value
		// prints 'strconv.ParseInt: parsing "non-int": invalid syntax'
		fmt.Println(err)
	}
	// we'll revisit interfaces a little latter. meanwhile
	learnConcurrency()
}

// c is a channel, a concurrency-safe communication object
func inc(i int, c chan int) {
	c <- i + 1 // <- is the "send" operator when a channel appears on the left
}

// we'll use inc to increment some numbers concurrently
func learnConcurrency() {
	// same make function used earlier to make a slice. make allocates and
	// initializes slices, maps, and channels
	c := make(chan int)
	// start three concurrent goroutines. numbers will be incremented
	// concurrently, perhaps in parallel if the machine is capable and properly
	// configured. all three send to the same channel
	go inc(0, c) // go is a statement that starts a new goroutine
	go inc(10, c)
	go inc(-805, c)
	// read three results from the channel and print them out
	// there is no telling in what order the results will arrive
	fmt.Println(<-c, <-c, <-c) // channel on right, <- is "receive" operator

	cs := make(chan string)       // another channel, this one handles strings
	ccs := make(chan chan string) // a channel of string channels
	go func() { c <- 84 }()       // start a new goroutine just to send a value
	go func() { cs <- "wordy" }() // again, for cs this time
	// select has syntax like a switch statement but each case involves a
	// channel operation. It selects a case at random out of the cases that are
	// ready to communicate
	select {
	case i := <-c: // the value received can be assigned to a variable
		fmt.Printf("it's a %T", i)
	case <-cs: // or the value received can be discarded
		fmt.Println("it's a string")
	case <-ccs: // empty channel, not ready for communicate
		fmt.Println("didn't happen.")
	}
	// at this point a value was taken from either c or cs. one of the two
	// goroutines started above has completed, the other will remain blocked

	learnWebProgramming() // go does it. you want to do it too
}

// a single function from package http starts a web server
func learnWebProgramming() {
	// first parameter of ListenAndServe is TCP address to listen to
	// second parameter is an interface, specifically http.Handler

	go func() {
		err := http.ListenAndServe(":8080", pair{})
		fmt.Println(err) // don't ignore errors
	}()

	requestServer()
}

// make pair an http.Handler by implementing its only method, ServeHTTP
func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// serve data with a method of http.responseWriter
	w.Write([]byte("You learned Go in Y minutes"))
}

func requestServer() {
	resp, err := http.Get("http://localhost:8080")
	fmt.Println(err)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Printf("\nWebserver said: `%s`", string(body))
}
