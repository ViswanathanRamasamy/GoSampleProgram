/*
Lesson2: sample program of Golang for the beginers.
Source Code Location: https://github.com/ViswanathanRamasamy/GoSampleProgram.git
*/

package main // function is going to be the part of the main
import (
	"fmt"
) // import the required package
/* Difference between the slice and array?
size of the array is fixed.
slice array can be increased.

len represent the current number of elements holding
capcity : maximum number of elements that the slice can hold with allocating the addition al memory
slice will not take value in []
*/

func main1() {
	number1 := []int{}              // empty slice
	number2 := []int{1, 2, 3, 4, 5} //slice with 5 elements
	fmt.Printf("%v\n", number1)     // []
	fmt.Printf("%v\n", number2)     // [1 2 3 4 5]

	fmt.Printf("%d\n", len(number1)) //0
	fmt.Printf("%d\n", len(number2)) //5

	fmt.Printf("%d\n", cap(number1)) //0
	fmt.Printf("%d\n", cap(number2)) //5

	number2 = append(number2, 20, 21) //[1 2 3 4 5 20 21]
	//if we are going to append 2 slice
	//myslice3 := append(myslice1, myslice2...)

	fmt.Printf("%v\n", (number2)) //5
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4] //since the slice can hold from the 2 element to the end of the array it capacity is 4
	//myslice := arr1[:4] // here we create the slice from the element 0 to 3

	fmt.Println(arr1, " \n")        //[10 11 12 13 14 15]
	fmt.Println(myslice, "\n ")     //[12 13]
	fmt.Println(len(arr1), "\n ")   //6
	fmt.Println(cap(arr1), "\n")    //6
	fmt.Println(len(myslice), "\n") //2
	fmt.Println(cap(myslice), "\n") //4

	//myslice1 := make([]int, 5, 10) //len is 5 and capcity is 10
	//myslice1 := make([]int, 5, 10) //len and capacity is 5

	//numbersCopy := make([]int, len(neededNumbers))
	//copy(numbersCopy, neededNumbers)
	//if the numberofcopy has more length , all the elements from the needed number will be copied to numbersCopy
	//remaining elements are zero. if neededNumbers is more length then only the  required number element will be copied to  numbersCopy

}

//Logical operator && || !
//BitWIse operator & | ^ <<  >>
//Arithamatic : var x =6+3; x += 90
//Comparision: > , >=, <,<=, == ,!=

/*
if ()
{

}
else
{

}

if ()
{

}
else if()
{

}
else
{

}

Here the day is number . if the daty is number you cannot use the string in case
switch day
{
    case 1: //or case (1):
	case 2: /or case (2):
    default:
}

Here the day is string
switch day
{
	case "1":
	case "2":
    default:
}

multi-case switch:
switch expression {
case x,y:
   // code block if expression is evaluated to x or y
case v,w:
   // code block if expression is evaluated to v or w
case z:
...
default:
   // code block if expression is not found in any cases
}
*/
//========================================================

func main3() {
	for i := 0; i <= 100; i += 10 { // Not we cannot use var or ( in the for loop
		fmt.Println(i) //0 10 20 30 40 50 60 70 80 90 100
	}
}

//========================================================

// for  loop has continue break
// nested for loop
// nested for loop
func main4() {
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}
}

//========================================================

// range keyword provides the idex and the value. range works for array, slice and map
func main() {
	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}
}

//========================================================

//if we want to omit the index

func main5() {
	fruits := [3]string{"apple", "orange", "banana"}
	for _, val := range fruits { // to omit the value for idx, _ := range fruits {
		fmt.Printf("%v\n", val)
	}
}

//========================================================

// calling a function with argument. if no argument then we can call like func familyName()
func familyName(fname string, age int) {
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func main6() {
	familyName("Liam", 3)
	familyName("Jenny", 14)
	familyName("Anja", 30)
}

//========================================================

// function return the result. in the below example the function is doing 2 return
// ine is int and string
// name return
func myFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!" //string concatentation
	return
}

// Here we are ignoring the 2nd return parameter. similarly we can also ignore the first return parameter
func main7() {
	a, _ := myFunction(5, "Hello")
	a1, b1 := myFunction(5, "Hello") //Here both the reurn arguments are used
	fmt.Println(a)
	fmt.Println(myFunction(5, "Hello")) //10 Hellow World
}

//========================================================

// uname return
func myFunction(x int, y int) (int, string) { // retun the int and string // if it is single return you can put func myFunction(x int, y int) int
	return x + y, "Hellow" //3, Hellow
}

func main8() {
	fmt.Println(myFunction(1, 2)) //3 Hellow
}

// ========================================================
func myFunction(x int, y int) (result int) {
	result = x + y
	return result //simply return since the result is specified in the function first line
}

func main() {
	fmt.Println(myFunction(1, 2))
}

// =============================
// recursive function for loop
func factorial_recursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorial_recursion(x-1)
	} else {
		y = 1
	}
	return
}

func main12() {
	fmt.Println(factorial_recursion(4))
}

// =====================================
// package main : no  ""
// function definition can be after the call
// structure variable //type strcuturename struct
// structure variable //var p structurename
type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	// Print Pers1 info by calling a function
	printPerson(pers1)

	// Print Pers2 info by calling a function
	printPerson(pers2)
}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}

/*
 var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"} // print the element %v a
  b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}  //prprint the elemnt %v b

  //Tuple using make
var a = make(map[KeyType]ValueType)
b := make(map[KeyType]ValueType)
//(a == nil) is true
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"

  print the invidual element a["year"]
   delete(a,"year") //delete the key
   val1, ok1 := a["brand"] if the brand is present in a ford will be copied to val1 and ok1 is true
   b := a //map are reference;
   any changes in b will impact and vice versa since the maps are copies as reference

   func main() {
  a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

  var b []string             // defining the order //b is nil silice
  b = append(b, "one", "two", "three", "four")

  for k, v := range a {        // loop with no order
    fmt.Printf("%v : %v, ", k, v)
  }

  fmt.Println()

  for _, element := range b {  // loop with the defined order
    fmt.Printf("%v : %v, ", element, a[element])
  }
}

two : 2, three : 3, four : 4, one : 1,
one : 1, two : 2, three : 3, four : 4,
=======================================
when the return is not required
package main
import ("fmt")

func myFunction(){
  fmt.Println("I just got executed!")
}

func main() {

myFunction()

}
*/
/*
myslice when using make need the length and the capacity. capacity is not must.
slice can be copied to other other slice variable using copy?
range wiull work on the array , slice, map. range value or key can be ignored using_.
when we specify the return variable in the function definition then return <return variable > is not must it can be return;
*/
