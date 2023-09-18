/*
Factorial: sample program of Golang for the beginers.
Source Code Location: https://github.com/ViswanathanRamasamy/GoSampleProgram.git
Author Name: Viswanathan Ramasamy
Email id: rviswawt@gmail.com
*/
//====================================================================
// main should main {
// after the . first charater should be capital letter
package main
import "fmt"

func factorial(n int) int {
    if n<=1 {
        return 1
    } else {
        return n * factorial (n-1)
    }
}

func main() {
    var input int
    _,err := fmt.Scanln(&input)
    
    if err != nil {
        fmt.Println(err)
        return
    }
    
    fmt.Println(factorial(input))
}
================================================
//there is no static variable concelpt in go lang but we can make use of the package variable to do it.
package main
import "fmt"

var finalresult int

func fib(input int) int {
    if input <= 1 {
        return input
    } else {
    
        finalresult = (fib(input-1)+fib(input-2))
        return finalresult
    }
}

func main() {
    var input int
    _,err := fmt.Scanln(&input)
    
    if err != nil {
        fmt.Println(err)
        return
    }
    
    for i:=0 ; i <=input; i++ {
        fmt.Print(fib(i))
    }
}
/////////////////////////////////////
package main
import "fmt"


func fib(input int) int {
    if input <= 1 {
        return input
    } else {
    
        return(fib(input-1)+fib(input-2))
    }
}

func main() {
    var input int
    _,err := fmt.Scanln(&input)
    
    if err != nil {
        fmt.Println(err)
        return
    }
    
    for i:=0 ; i <=input; i++ {
        fmt.Print(fib(i))
    }
}
///////////////////////////////////
fibanancis 
package main
import "fmt"


func main() {
    var input int
    _,err := fmt.Scanln(&input)
    
    if err != nil {
        fmt.Println(err)
        return
    }
    
    a, b := 0,1
    for i:=0 ; i <=input; i++ {
        fmt.Print(a, " ")
        a,b = b,a+b
    }
}
==========================================