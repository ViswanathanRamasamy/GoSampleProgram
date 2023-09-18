/*
String Program: sample program of Golang for the beginers.
Source Code Location: https://github.com/ViswanathanRamasamy/GoSampleProgram.git
Author Name: Viswanathan Ramasamy
Email id: rviswawt@gmail.com
*/
//====================================================================
//stack 
package main
import "fmt"
import "strconv"

func main() {
    input := 0
    _,err := fmt.Scanln(&input)
    
    if err !=nil {
        return
    }
    
    stack := make([]string, input)
    
    for i:=0; i<input; i++ {
        stack[i] ="Hello"+strconv.Itoa(i)
    }
    
    for len(stack) >0 {
        fmt.Println(stack[len(stack)-1])
        stack = stack[:len(stack)-1]
    }
}
5
Hello4
Hello3
Hello2
Hello1
Hello0
=======================================