/*
Lesson4: sample program of Golang for the beginers.
Source Code Location: https://github.com/ViswanathanRamasamy/GoSampleProgram.git
Author Name: Viswanathan Ramasamy
Email id: rviswawt@gmail.com
*/
==========================================================================
JSON marshaling (or encoding) is the process of converting Go data structures (typically structs, slices, maps, and basic types) 
into a JSON-formatted string

package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	// Create an instance of the Person struct
	person := Person{
		Name:    "John Doe",
		Age:     30,
		Address: "123 Main St",
	}

	// Marshal the struct into a JSON string
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the JSON data as a string
	fmt.Println("JSON Data:", string(jsonData))

	// Unmarshal JSON data into a new struct
	var newPerson Person
	err = json.Unmarshal(jsonData, &newPerson)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Access the fields of the new struct
	fmt.Println("Name:", newPerson.Name)
	fmt.Println("Age:", newPerson.Age)
	fmt.Println("Address:", newPerson.Address)
}

JSON Data: {"name":"John Doe","age":30,"address":"123 Main St"}
if you remove the struct tag
type Person struct {
	Name    string 
	Age     int    
	Address string 
}
JSON Data: {"Name":"John Doe","Age":30,"Address":"123 Main St"}
===================================================

// Program to write the content to the file and read the file.
//defer keyword ensure that all the operations are completed before executing the defer command
// only the create file we are clsoing the file
//ioutil.Readstring no close required.

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func errorCheck(err error) {
	if err != nil {
		panic(err) //panic error can have only one message argument
	}
}
func openandWrite(name string) {
	var content string = "Hellow World"
	file, error := os.Create(name)
	errorCheck(error)
	len, error := io.WriteString(file, content)
	errorCheck(error)
	fmt.Println("print length written", len)
	defer file.Close()
}

func ReadFile(name string) {
	data, error := ioutil.ReadFile(name) //will read the file and put it in array
	errorCheck(error)
	fmt.Println(string(data)) //convert the array to string
}

func main() {
	var filename string = "./viswa.txt"
	openandWrite(filename)
	defer ReadFile(filename)

}

/*
 net/http => get
 io => parse the response body
 encoding/json => parse the response in json
 note error: = in multiple places

 response, error := http.Get()

 "%T", response ==> http response

 defer response.Body.Close()

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// Make an HTTP request to get JSON data
	response, err := http.Get("https://example.com/api/tours") // Replace with your API URL
	if err != nil {
		fmt.Println("HTTP request failed:", err)
		return
	}
	defer response.Body.Close()

	// Read the raw JSON data from the response
	rawBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Convert raw JSON bytes to a string
	content := string(rawBytes)

	// Parse JSON data and get the tours
	tours := jsonreader(content)

	// Print the parsed tours
	for _, tour := range tours {
		fmt.Printf("Name: %s\nAddress: %s\n\n", tour.Name, tour.Add)
	}
}

func jsonreader(content string) []Tour {
	var tour Tour
	var tours []Tour

	decoder := json.NewDecoder(strings.NewReader(content))

	// Check for the JSON start token
	_, err := decoder.Token()
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	// Loop through the JSON array and decode Tour structs
	for decoder.More() {
		err := decoder.Decode(&tour)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			continue // Skip this item and proceed with the next
		}
		tours = append(tours, tour)
	}

	return tours
}

type Tour struct {
	Name string `json:"name"`
	Add  string `json:"address"`
}


   http.Get() return the pointer
*/
============================
copy the content from one file to another file:
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Open the source file for reading
	sourceFileName := "source.txt"
	sourceFile, err := os.Open(sourceFileName)
	if err != nil {
		fmt.Println("Error opening source file:", err)
		return
	}
	defer sourceFile.Close()

	// Create or open the destination file for writing
	destinationFileName := "destination.txt"
	destinationFile, err := os.Create(destinationFileName)
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		return
	}
	defer destinationFile.Close()

	// Copy the content from the source file to the destination file
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		fmt.Println("Error copying content:", err)
		return
	}

	fmt.Println("File copied successfully!")
}
==================================================

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Create a scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	// Create a slice to store the integers
	var integers []int

	fmt.Println("Enter integers (press Enter without input to finish):")

	for {
		fmt.Print("Enter an integer: ")
		scanner.Scan()
		input := scanner.Text()

		// Check if the input is empty, and exit the loop if it is
		if input == "" {
			break
		}

		// Parse the input as an integer
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		// Add the integer to the slice
		integers = append(integers, num)
	}

	if scanner.Err() != nil {
		fmt.Println("An error occurred while reading input:", scanner.Err())
		return
	}

	// Print the collected integers
	fmt.Println("Collected integers:", integers)
}
========================================
package main

import (
	"fmt"
)

func main() {
	var integers []int

	fmt.Println("Enter integers (press Enter without input to finish):")

	for {
		var num int
		_, err := fmt.Scanln(&num)

		if err != nil {
			fmt.Println("An error occurred while reading input:", err)
			break
		}

		integers = append(integers, num)
	}

	fmt.Println("Collected integers:", integers)
}
=====================================================
\d : single digit
\d+ : number
here fraction will printed as the seperate number
\.\d+
\d{3,5}? matches between 3 and 5 digits but prefers the shorter match (lazy quantifier).
\\[A-Za-z]+ print the word
package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	inputString := "The price of the product is $50.99, and the discount is 10%."
	
	// Define a regular expression to match numbers
	re := regexp.MustCompile(`(\d+\.\d+)?`) //(`\d+\.\d+|\d+`) //
	
	// Find all matches in the input string
	matches := re.FindAllString(inputString, -1)
	
	// Convert the matched strings to numbers
	var numbers []float64
	for _, match := range matches {
		number, err := strconv.ParseFloat(match, 64)
		if err == nil {
			numbers = append(numbers, number)
		}
	}
	
	// Print the extracted numbers
	fmt.Println("Extracted numbers:", numbers)
}
=======================================
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Create a slice of integers
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

	// Use the sort.Ints() function to sort the slice in ascending order
	sort.Ints(numbers)

	// Print the sorted slice
	fmt.Println(numbers)
}
======================================
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Create a slice of strings
	words := []string{"apple", "banana", "cherry", "date"}

	// Use the sort.Strings() function to sort the slice in lexicographic order
	sort.Strings(words)

	// Print the sorted slice
	fmt.Println(words)
}
=======================================================
In Go, a closure is a function value that references variables from outside its body. Closures are a powerful and flexible feature of the language.
They allow functions to capture and remember their surrounding context, 
package main

import "fmt"
//2 return statement
func makeCounter() func() int {
	count := 0

	// This function is a closure because it captures the 'count' variable
	// from the surrounding context.
	return func() int {
		count++
		return count
	}
}

func main() {
	counter1 := makeCounter()
	counter2 := makeCounter()

	fmt.Println(counter1()) // 1
	fmt.Println(counter1()) // 2
	fmt.Println(counter2()) // 1 (Separate closure instance)
}
==============
package main

import "fmt"

func multiplyBy(factor int) func(int) int {
    // This closure takes an 'int' argument and returns a new function
    // that multiplies the argument by 'factor'.
    return func(x int) int {
        return x * factor
    }
}

func main() {
    // Create a closure that multiplies by 2
    double := multiplyBy(2)

    // Use the closure to double a number
    result := double(5)
    fmt.Println(result) // Output: 10
}
==============================
2 return statement
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 5; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
==
0 0
1 -2
3 -6
6 -12
10 -20
=======================
here is no traditional inheritance as you might find in object-oriented languages like Java or Python. 
Instead, Go provides a different mechanism called composition, which allows you to achieve similar functionality through embedding.
package main

import "fmt"

// Parent struct
type Person struct {
    FirstName string
    LastName  string
}

// Child struct that embeds the Parent struct
type Student struct {
    Person
    StudentID int
}

func main() {
    // Create a Student object
    student := Student{
        Person: Person{
            FirstName: "John",
            LastName:  "Doe", //note , here
        }, //note , here
        StudentID: 12345,
    }

    // Access fields from both parent and child structs
    fmt.Println("First Name:", student.FirstName)
    fmt.Println("Last Name:", student.LastName) //or student.Person.LastName
    fmt.Println("Student ID:", student.StudentID) 
}
==============================================
package main

import (
	"fmt"
)

func main() {
	var val interface{} = 42

	switch s:=val.(type) {
	case int:
		fmt.Println("It's an int!", s)
	case string:
		fmt.Println("It's a string!")
	default:
		fmt.Println("It's another type!")
	}
}
==================
s.(type) will work only for switch
 if s.(type) == int {
        fmt.Println("hi")
    }
====================
THIS IS VALID:    
	if value, ok:=s.(int); ok {
        fmt.Println(value)
    }
=====================
package main

import "fmt"

func main() {
    var val interface{} = 42

    if v, ok := val.(int); ok {
        fmt.Println("It's an int:", v)
    } else if v, ok := val.(string); ok {
        fmt.Println("It's a string:", v)
    } else {
        fmt.Println("It's another type!")
    }
}
===================================
In Go, an interface is a type that specifies a set of method signatures. It defines a contract that concrete types must adhere to
 if they want to be considered as implementing that interface.
 Interfaces enable polymorphism and decoupling in your Go programs, making your code more flexible and extensible.
package main

import (
    "fmt"
)

type Shape interface {
    Area() float64
}

type Square struct {
    SideLength float64
}

func (s Square) Area() float64 {
    return s.SideLength * s.SideLength
}

type Rectangle struct {
    Length float64
    Width  float64
}

func (r Rectangle) Area() float64 {
    return r.Length * r.Width
}

func main() {
    var shape Shape

    square := Square{SideLength: 5}
    rectangle := Rectangle{Length: 4, Width: 6} //order can be changed inside {Width:10, length:100 }

    // Assigning a Square to the Shape interface
    shape = square
    fmt.Printf("Area of square: %f\n", shape.Area())

    // Assigning a Rectangle to the Shape interface
    shape = rectangle
    fmt.Printf("Area of rectangle: %f\n", shape.Area())

    // Type assertion to get the underlying type
    if square, ok := shape.(Square); ok {
        fmt.Printf("It's a square! Side length: %f\n", square.SideLength)
    } else {
        fmt.Println("It's not a square.")
    }
}
==============================================================
program for stack:

stack program:
package main

import "fmt"

func main() {

// Create

var stack []string

// Push
stack = append(stack, "! ")
stack = append(stack, "world")

stack = append(stack, "Hello ")

for len(stack) > 0 {

        // Print top

        n := len(stack) - 1

        fmt.Print(stack[n])

        // Pop

        stack = stack[:n]

}

// Output: Hello world!

}
====================================