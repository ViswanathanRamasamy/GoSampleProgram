/*
String Program: sample program of Golang for the beginers.
Source Code Location: https://github.com/ViswanathanRamasamy/GoSampleProgram.git
Author Name: Viswanathan Ramasamy
Email id: rviswawt@gmail.com
*/
//====================================================================
//Get the user inout
//Note else {  should start at the } else {
//Scanln 

func main() {
    var inputString int //int to get the integer and string to get the string`
    numberItem,_err := fmt.Scanln(&inputString)
    
    if _err == nil {
        fmt.Println(inputString, numberItem )
    } else {
        fmt.Println(_err)
    }
}
//====================================================================
================================================================================
//we cannot pass the non constan to an array
//in that case you can use slice:
length := 100
var outputInt []int // Declare an empty integer slice
outputInt = make([]int, length+1) // Initialize the slice with a specified length
fmt.Println("Slice length:", len(outputInt))
=================================
//Note for loop
// for i:=0, j:=length // not allowed. It has to for i,j:=0, length
//for i++,j++ or for i,j = i++, j++ are not allowed it has to be i,j=i+1,j+1 
/* Reverese the string */
// there is char data type in go Lang
// we cannot assign the char to the string using index like str[i] = 'a'
package main
import "fmt"
func main() {
    inputstring := "hi how are you ?"
    reveredstring := reversestr(inputstring)
    fmt.Println("reversed string", reveredstring)
}

func reversestr(ainputstring string) (string) {
    
    length := len(ainputstring) -1;
    
    ainput := [] rune (ainputstring);
    
    for i, j:=0, length ; i <=length/2; i,j= i+1, j-1 {
        ainput[i],ainput[j] = ainput[j],ainput[i]
    }
    
    return string(ainput)
}
==================================================
function to print the invidual character:
func main() {
    input := "Hello, World!"
    for i:=0; i<len(input);i++ {
        fmt.Printf("%c" , input[i])
		//fmt.Println(input[i]) //will give the decimal value
    }
}
=================================================

func main() {
    input := "Hello, World!"
    output := ""
    for i:=len(input)-1; i>=0; i-- {
        output += string(input[i])
    }
    fmt.Print(output)
}
========================================
convert the string to upper case:
for loop to itertae over the character array

package main
import "fmt"

func upper(ainput string) (string) {
    
    output := []rune(ainput)
    
    for i,char := range(output) {
        if char > 'a' && char < 'z' {
            output[i] -= 32
        } 
    }
    return string(output)
}

func main() {
    input := "hello WORld"
    fmt.Print(upper(input))
}
=======================================================
not condition check
if !(ch >= 'a' && ch <= 'z') {
=============================================
we cannot change the value of the string at the particular index. it is immutable
//rune and byte different?
/* byte:

byte is an alias for uint8.
It is used to represent a single 8-bit byte, which can store values from 0 to 255.
It's commonly used when dealing with ASCII characters or binary data.

rune is an alias for int32.
It is used to represent a Unicode code point, which can range from 0 to 0x10FFFF (the valid Unicode range).
*/

func main() {
    input := "hello WORld"
    input [1] = 'v' //error
    fmt.Print(input)
}

func main() {
    input := "hello WORld"
    output := []byte(input)
    output[1] ='v'
    input = string(output)
    fmt.Print(input)
}
