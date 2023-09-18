/*
Lesson2: sample program of Golang for the beginers.
Source Code Location: https://github.com/ViswanathanRamasamy/GoSampleProgram.git
Author Name: Viswanathan Ramasamy
Email id: rviswawt@gmail.com
*/
//////////////////////////////////////////////
A slice is a dynamic data structure in Go that represents a portion of an array. Unlike arrays, slices can change in size, and they 
don't have a fixed length.

Difference between the regular string and raw string?
raw string doesn't interpreet the \n\t 
package main

import "fmt"

func main() {
    rawString := `This is a regular string.\nIt interprets escape sequences like \\n and \\t.`
    fmt.Println(rawString)
}

This is a regular string.\nIt interprets escape sequences like \\n and \\t
==========================
func main() {
   emptySlice := make([]int, 0, 5) // Creates an empty slice with a length of 0 and capacity of 5
   //emptySlice[0] = 100 will crash
   emptySlice =append(emptySlice,100)
   emptySlice[0] =200 //changing 100 to 200
   fmt.Printf("%v",emptySlice) //[200]
   
}
===================================
sort the numbers:

func main() {
    numbers := []int{5, 1, 4, 2, 3}
    
    for i:=0; i < len(numbers); i++ {
        for j:=i+1; j < len(numbers); j++ {
            if numbers[i] > numbers[j] {
                numbers[i],numbers[j] = numbers[j], numbers[i]
            }
            
        }
    }
   
   fmt.Print(numbers)
}
======================
find whether the key is present 
if the key is found delete the key

package main

import "fmt"

func main() {
    // Create a map with some key-value pairs
    studentGrades := map[string]int{
        "Alice":  90,
        "Bob":    85,
        "Charlie": 78,
    }

    // Check if a key exists in the map
    name := "Bob"
    if _, exists := studentGrades[name]; exists {
        fmt.Printf("%s's grade exists: %d\n", name, studentGrades[name])
    } else {
        fmt.Printf("%s's grade does not exist\n", name)
    }

    // Delete a key from the map if it exists
    keyToDelete := "Charlie"
    if _, exists := studentGrades[keyToDelete]; exists {
        delete(studentGrades, keyToDelete)
        fmt.Printf("%s's grade deleted. Updated map: %v\n", keyToDelete, studentGrades)
    } else {
        fmt.Printf("%s's grade does not exist in the map. No deletion.\n", keyToDelete)
    }
}
================================
Block scope:
func main() {
    const x = 10
    fmt.Println(x) // Accessible here
}

func anotherFunc() {
    // fmt.Println(x) // Not accessible here, x is out of scope
}
==========
Package Scope: Constants declared at the package level (outside of any function) have package-level scope, and they can be accessed throughout
 the entire package.

package mypackage

const Pi = 3.14159

func main() {
    fmt.Println(Pi) // Accessible here
}
===================================
Go Init:
In Go, the "init" function is a special function that is automatically called by the Go runtime when a package is initialized. 
It is called before the main function and can be used to perform initialization tasks for the package.

The "init" function does not take any arguments and does not return a value. It is typically used to set initial values for package-level variables,
 establish connections to external resources such as databases, or perform any other initialization tasks that need to be performed before
the main function is called.
//there can only be one main function in the entire program.

package main

import (
    "fmt"
)

func init() {
    fmt.Println("First init function")
}

func init() {
    fmt.Println("Second init function")
}

func main() {
    fmt.Println("Main function")
}
First init function
Second init function
Main function
================================================
struct:
A struct or a structure of Golang is a user-defined variety that helps the group or 
combines items of various types into a single type, each real-world entity that holds some characteristics can be represented as a structure
==================
static declaration:
const MyConstant = 42
var GlobalVariable = "Hello, World!"

type MyStruct struct {
    Field1 int
    Field2 string
}
====================
dynamic declaration:
func main() {
    // Dynamic allocation of a slice
    mySlice := make([]int, 5)

    // Dynamic allocation of a new instance of a custom type
    obj := new(MyStruct)

    // Variables created inside a function are also dynamically allocated
    localVar := "Dynamic"
}
====================
Go routine:
In the Go programming language, a goroutine is a lightweight thread of execution. Goroutines are used to perform tasks concurrently,
 and they are multiplexed onto a small number of OS threads, so they are very efficient.

Goroutines are different from traditional threads in several ways. They are multiplexed onto real threads, so there is not a
 one-to-one correspondence between goroutines and OS threads. This means that you can have many goroutines running 
concurrently on a small number of OS threads. Additionally, goroutines are very lightweight, so it is not expensive to create and manage them.
===========================================
Data type:
bool: a boolean value (true or false) 
int, int8, int16, int32, int64: signed integers of various sizes 
uint, uint8, uint16, uint32, uint64: unsigned integers of various sizes 
float32, float64: floating-point numbers 
complex64, complex128: complex numbers 
string: a string of Unicode characters 
byte: an alias for uint8 
rune: an alias for int32 
=============================
make and const will not have variables
var myMap map[string]int

myMap := make(map[string]int)
// OR
myMap := map[string]int{}

var myMap map[string]int // This is nil
===================================
package main

import (
	"fmt"
)

func mapsEqual(map1, map2 map[string]int) bool {
	// Check if the lengths of the maps are the same
	if len(map1) != len(map2) {
		return false
	}

	// Iterate over map1 and compare key-value pairs with map2
	for key, value1 := range map1 {
		value2, ok := map2[key]
		if !ok || value1 != value2 {
			return false
		}
	}

	return true
}

func main() {
	map1 := map[string]int{"apple": 5, "banana": 3, "cherry": 7}
	map2 := map[string]int{"apple": 5, "banana": 3, "cherry": 7}
	map3 := map[string]int{"apple": 5, "banana": 2, "cherry": 7}

	fmt.Println("map1 == map2:", mapsEqual(map1, map2)) // true
	fmt.Println("map1 == map3:", mapsEqual(map1, map3)) // false
}
============================================