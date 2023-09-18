/*
sample program of Golang for the beginers.
Source Code Location: https://github.com/ViswanathanRamasamy/GoSampleProgram.git
Author Name: Viswanathan Ramasamy
Email id: rviswawt@gmail.com
*/

Why choose Go for development?
- **Answer:** Go offers several advantages for development, including:
    - Strong performance and efficiency.
    - A simple and clean syntax that promotes readability and maintainability.
    - Concurrency support with goroutines and channels.
    - A rich standard library.
    - Cross-platform compatibility.
    - Strong focus on simplicity, safety, and scalability.
    - A large and active community
	- strong typing, faster compilation
	- general purpose high level language

//const dont require var
	
// To add comments
/* multiple line comments */
// Filename notation: file name should have .go
// To execute : go run ./<filename>.go
package main // function is going to be the part of the main
import (
	"fmt"
) // import the required package

func main1() { //{ cannot be the first line}
	fmt.Println("Hello, World") // ; is not required
}

// function can be in single line
func main2() { fmt.Println("Hello, World") }

//different type of variable are int, float32, string, bool
// 2 ways to initilize the variable direct and inferred. in inferred 2 ways
// note the go will initialize the variable to the default value if not initialized

// var a bool; var b string; var c int; //false "" and 0

// what is the difference between the var and inferrred?
// a. var can be used inside and outside the function with or with out initialization
// b. variable declaration and definition can be different line
// and b is not possible in inferred
//if we print the empty string "" in the console nothing will be printed.

func main3() {
	var student1 string = "John" //type is string
	var student2 = "Jane"        //type is inferred // cannot use := here
	x := 2                       //type is inferred
	student3 := "Jane"           // string inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println(student3)
	student1 = "viswa" //changing trhe content of the variable
	fmt.Println(student1)
}

// multiple variable declartion and initialization
// var a, b, c, d = 1, 3, 5, 7 or var a, b, c, d int = 1, 3, 5, 7
// a, b, c, d := 1, 3, 5, 7
// in multi variable initialization we need to initializa all the variable. cannot initiallize partially
func main4() {
	a, b, c, d := 1, 3, 5.0, "hellow"

	fmt.Println(a) //1
	fmt.Println(b) //3
	fmt.Println(c) //5.0
	fmt.Println(d) //hellow
}

// In scenerios if you like to partially initialize the partial variable
func main5() {

	var (
		a int
		b int    = 1
		c string = "hello"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// Go variable Rules
// No restruction on the length of the variable
// variable name cannot have space , start with digit and variables are case sensitive/
//Camel Case: myVariableName = "John"
//Pascal Case: MyVariableName = "John"
//Snake Case: my_variable_name = "John"
//const can be declared or out side the function . They Generally return in capital letter
//we cannot define the const  like const pi:= 3.14 or var const pi float32 =3.14

const PI = 3.14
const PI1 float32 = 3.14

const (
	A int = 1
	B     = 3.14
	C     = "Hi!"
)

func main6() {
	const PI2 = 3.14
	fmt.Println(PI)
	fmt.Println(PI1)
	fmt.Println(PI2)
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}

// otherway to print the out
// Print
// Println
// printf
// Print() inserts a space between the arguments if neither are strings:
// The Println() function is similar to Print() with the difference that
// a whitespace is added between the arguments, and a newline is added at the end:
func main7() {
	var i, j string = "Hello", "World"

	fmt.Print(i)
	fmt.Print(j)    //HelloWorld in sample line
	fmt.Print(i, j) //HelloWorld in sample line

	fmt.Print(i, "\n") //Hello
	fmt.Print(j, "\n") // Word

	fmt.Print(i, "\n", j, "\n") //Hello  world both in seperate line

	k, m := 1, 2
	fmt.Print(k)
	fmt.Print(m)    //12
	fmt.Print(k, m) //1 2

	fmt.Printf("\ni has value: %v and type: %T\n", i, i) // %v gives the value and %T will give you the type
	fmt.Printf("j has value: %v and type: %T", j, j)
}

func main8() {
	var i = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)   //15.5 //value
	fmt.Printf("%#v\n", i)  // go syntax //15.5
	fmt.Printf("%v%%\n", i) //15.5%
	fmt.Printf("%T\n", i)   //float64

	fmt.Printf("%v\n", txt)  //Hellow World!
	fmt.Printf("%#v\n", txt) //"Hello World!"
	fmt.Printf("%T\n", txt)  //string
}

func main9() {
	var i = true //No True
	var j = false

	fmt.Printf("%t\n", i) //True //yiu can also use %v for boot
	fmt.Printf("%t\n", j) //false //%t cannot be used for the other data types
}

func main10() {
	var txt = "Hello"

	fmt.Printf("%v\n", txt)      //Hello
	fmt.Printf("%s\n", txt)      //Hello // no %#s
	fmt.Printf("%q\n", txt)      //"Hello"
	fmt.Printf("%8s\n", txt)     //   Hello
	fmt.Printf("%-8shi \n", txt) //Hello   hi
	fmt.Printf("%x\n", txt)      //48656c6c6f
	fmt.Printf("% x\n", txt)     //48 65 6c 6c 6f

}

func main11() {
	var i = 15

	fmt.Printf("%b\n", i)   //1111
	fmt.Printf("%d\n", i)   //15
	fmt.Printf("%+d\n", i)  //+15
	fmt.Printf("%o\n", i)   //17
	fmt.Printf("%O\n", i)   //0o17
	fmt.Printf("%x\n", i)   // f
	fmt.Printf("%X\n", i)   //F
	fmt.Printf("%#x\n", i)  //0xF
	fmt.Printf("%4d\n", i)  //  15
	fmt.Printf("%-4d\n", i) //15  ..
	fmt.Printf("%04d\n", i) //0015
}

// we have 3 data types
func main12() {
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	fmt.Println("Boolean: ", a) //true
	fmt.Println("Integer: ", b) //5
	fmt.Println("Float:   ", c) //3.14
	fmt.Println("String:  ", d) //HI
}

//int , int8, int16, int 32, int 64. int can be 32 or 64 depends on the machine
// float 32 or float64
//array

func main() {

	var arr1 = [3]int{1, 2, 3}    // or var arr1 = [...]int{1,2,3}
	arr2 := [5]int{4, 5, 6, 7, 8} // or arr2 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)      //[1 2 3]
	fmt.Println(arr2)      //[4 5 6 7 8]
	fmt.Println(len(arr2)) // 5

	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Print(cars)        //[Volvo BMW Ford Mazda]
	fmt.Println(len(cars)) // 4

	fmt.Println(cars[1])
	cars[1] = "viswa"
	fmt.Println(cars[1])
	arr11 := [5]int{} // initialized [0 0 0 0 0]
	fmt.Println(arr11)

	var arr12 = [5]int{} // initialized [0 0 0 0 0] 
	//var arr12 [5]int{} //not allowed
	fmt.Println(arr12)
	// arr2 := [5]int{1,2} //partially initialized 12 0 0 0
	//arr1 := [5]int{1:10,2:40} // [0 10 40 0 0]
}

    func main10 () {
      //var numbers [5]int {} //not allowed
      var numbers [5]int // 0 0 0 0 0
	  numbers = [5]int{1, 2, 3, 4, 5}

      fmt.Println(numbers)
        
    }
============================
//print will what ever the content you passed.
//if format the specifier you need specify %f in the printf
// also not the typecasting of flaot requires both in the numerator and denominator.
//index value is neglected.

package main
import "fmt"

func main() {
    var array [6] int
    array = [6] int {1,2,33,4,7,5}
    
    sum := 0
    for _,data := range array {
        sum += data
    }
    
    fmt.Printf("%0.2f",float64(sum) / float64(len(array)))
}
=========================	
program to find the average of the element without using the lenpackage main

import "fmt"

func main() {
    // Declare an array of integers
    numbers := [5]int{10, 20, 30, 40, 50}

    // Declare variables to store the sum and length of the array
    var sum int
    var length int

    // Calculate the sum of the array elements and its length
    for _, num := range numbers { //num is local scope cannot be accessed outside the for loop
        sum += num
        length++
    }

    // Calculate the average
    average := float64(sum) / float64(length)

    // Print the array, sum, length, and average
    fmt.Println("Array:", numbers)
    fmt.Println("Sum:", sum)
    fmt.Println("Length:", length)
    fmt.Println("Average:", average)
}
=======================================
How to make the variable to access outside the for loop
package main
import "fmt"


func main() {
    var input [5] int
    input = [5] int {1,23,3,4,5}
    sum :=0
    i:=0
    data:=0
    for i,data = range input { 
        sum +=data
        i++;
    }
    
    fmt.Println(sum, i)
    fmt.Println(float64(sum)/float64(i)) //if i is in the scope of the for loop we will get the error i is undefined

}
==========================================
package main

import (
	"fmt"
	"strings"
)

func isPalindrome(str string) bool {
	// Remove spaces and convert to lowercase
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ToLower(str)

	// Check if the string is the same forwards and backwards
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	input := "A man a plan a canal Panama"
	if isPalindrome(input) {
		fmt.Println(input, "is a palindrome.")
	} else {
		fmt.Println(input, "is not a palindrome.")
	}
}
================================================
package main
import "fmt"

func main() {
   input := "MadaM"
   status := true
   length := len(input) -1
   
   for i,j:=0,length; i<=j; i,j =i+1,j-1 {
   
       if input [i] != input [j] {
           fmt.Println("not palindrome")
           status = false;
           break;
       }
   }
   
   if status == true {
       fmt.Println("palindrome")
   }
   
}

/*
To print the % in the output console use %% in the print
Printf has the format specifier not println and print
bool data types %t true false
*/
