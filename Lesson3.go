/*
Lesson3: sample program of Golang for the beginers.
Source Code Location: https://github.com/ViswanathanRamasamy/GoSampleProgram.git
*/
/*
Go implements:
go application is linked with the runtime package hence the size will be larger
in the case of java. we install the JVM in the client machine hence the runtime application size is less
This is not the case in Go

doent support
a. type interface ie reuse the paraent class function
b. no implict type conversion
c. no try catch except

Go installation path check
MSI is the binary installer for the windows.
in windows: cmd-> type path
in linux: export $PATH

To get the version of the go
Type: go version in command

Go file name should be of small letter and end with extension.go

To run  the go application

go run .\hello.go
go run hello.go
go run .

In the project there will be only one main function

To build the application for windows

hello folder contain main.go filename
hello folder >
go build . => will build the application for the current oS
GOOS="windows" go build
GOOS="linux" go build
=============================
data collection in go
arrays, slice, maps and struct
==============================
language organization:
function, interface and channel
================
data management:
pointer
====================
output:
we can specify the %f only in printf
=======================
variable declaration:
explict and implict:
explict:
var str string = "hello"
implict:
var str = "hello"
str := "hello"
=====================
Go support the user defined data types
if statement , for and switch statement doesn't require paranthesis ()
you can also initialize the variable in if

if x:=24; x>100 {
	fmt.Println(" ")
}
// import has to be appeared in the begining not at the middle of the file
// no break statemment in the case for Go
//similar to if switch can also take initialization
switch x:=24; x{
}
*/

package main // function is going to be the part of the main
import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort" //sor the elements
	"strconv"
	"strings"
	"time"
) // import the required package
/*
It is not necessary to include packages like fmt and sort in same line.
import "fmt"
import "sort"
//math.round(3.14) =>3

//elements stored in the map are unordered. ie it will be based on the hash value
output order is not guranteed

struct in go is custom type have both properties and methods

function definition and the struct declaration can be after the call/usage.

array vs slice
array:
var array[3] string
array = [3] string {}

slice:
var array[] string
array = [] string {}

append can also be used in slice to remove the last or first element
myslice = append(myslice[1:len(myslice)]) // remove the first element
myslice = append(myslice[:len(myslice)-1]) // remove the last element

myslice of data type int
sort.Ints(myslice) // will sort the element in the slice based on the int value in myslice
//myslice1 is string
sort.Strings(myslice1) // will sort the element in the slice based on the string value in myslice

difference between the new and make:
new will return the address but the memory is anot initialized, but make initialized
garbage collector will deallocate th memory
m := new (map[int]int)
m[0] = 1 //will crash
=============================
order of the map elements is not guranteed
==================================
var ptr int*;
p = nil;

vart i int
p = &i;
*p =10;

ptr1 := &i
==========================
f =:3.18
printf("%.1f", f) //3.2

println(f) //3.189999999999999999999

panic("This is error ") //when this line executes the program will print the error and stop the execution
//math.pi => 3.14
====================

for i:=1;i<10;i++ { //I variable scope is only with in this for loop

}
for loop can be used like while loop in c

varl:=1

	for varl <10 {
		varl++
	}

	goto label can be used in the for loop

Go has golabel:

goto labelEnd

labelEnd:

	fmt.Println()
*/
func main1() {
	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}

	for v, k := range myslice2 {
		fmt.Println(v, k) // 0 go ` Slices 2 Are 3 Powerful
	}

	state := make(map[string]string)
	state["s"] = "singapore"
	state["M"] = "Malaysia"
	state["I"] = "India"

	for k, v := range state {
		fmt.Println(k, v) // M Malaysia I India s singapore
	}

	for k := range state {
		fmt.Println(k) // S M I
	}

	i := 0
	myslice3 := make([]string, len(state), len(state))
	for k := range state {
		myslice3[i] = k
		i++
	}

	sort.Strings(myslice3)
	fmt.Println(myslice3) //I M S

	for i := range myslice3 {
		fmt.Println(myslice3[i])        // I M S
		fmt.Println(state[myslice3[i]]) //India Malaysia singapore
	}

}

/*
 type Person Struct //capital P indicates it is exported
 //access the invidual element like variable.<proprtyname>
*/
//Global
type Person struct //P is expoerted
{
	//The exported fields (Name and Age) are accessible from outside the
	//package, and the unexported field (salary) is only accessible within the package where it's defined.
	Name   string //Name field is exported
	Age    int
	salary int //Salary is not exported hence small letter
}

func main2() {
	var p Person = Person{"Viswa", 32, 15000}
	p1 := Person{"Viswa", 12, 15000}
	fmt.Println(p)        //Viswa 32 15000
	fmt.Println(p1)       // Viswa 12 15000
	fmt.Printf("%+v", p1) //{Name:Viswa Age:12 salary:15000}
}

func main3() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the text: ")
	inputText, _ := reader.ReadString('\n') //_ indicate to neglect the error obbject
	fmt.Println(inputText)
	fmt.Print("Enter the number: ")
	inputNumber, _ := reader.ReadString('\n') //flaot number is also read as string

	// Trim spaces and convert inputNumber to a float64
	afloat, err := strconv.ParseFloat(strings.TrimSpace(inputNumber), 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		// Round and print the result
		fmt.Printf("Rounded number: %.2f\n", afloat)
	}
}
func main4() {
	t := time.Now()
	fmt.Println(t)                    //2023-09-10 11:14:17.7436557 +0530 IST m=+0.003443701
	fmt.Println(t.Format(time.ANSIC)) //Sun Sep 10 11:15:25 2023

	t = time.Date(2009, time.November, 10, 22, 10, 0, 0, time.UTC) //
	fmt.Println(t)                                                 //2009-11-10 22:10:00 +0000 UTC
	fmt.Println(t.Format(time.ANSIC))                              //Tue Nov 10 22:10:00 2009
}

// import has to be appeared in the begining not at the middle of the file
// variable declared inside the switch case will be local to the switch
func main5() {
	rand.Seed(time.Now().Unix())
	down := rand.Intn(7) //randon number between 0 to 6
	fmt.Println(down)

	switch a := 10; a {
	//s:=1000; error
	case 10:
		s := 1000
		fmt.Println("number is 10", s)
		fallthrough
	case 11:
		fmt.Println("number is 11")
		fallthrough
	default:
		fmt.Println("number is 12")
	}
}

/*
number is 10
number is 11
number is 12
*/

type Student struct {
	Name    string
	Age     int
	message string
}

// // if you want the changes made in the receiver to function to impact the main
// func (d *Student) changeMessage() //here reference is made
// d is receiver.. copy is made
func (d Student) changeMessage() {
	// d.message = "hi hi hi hi" // or
	d.message = fmt.Sprintf("%v %v %v", "hi", "hi", "hi")
	fmt.Println(d.message) //hi hi hi
}

func main() {
	var obj Student = Student{"Viswa", 12, "hi"} //note { } not ()
	obj.changeMessage()
	fmt.Println(obj.message) //hi  //changes made inside the function will not imapac
	obj.message = "good"
}

/*

*/