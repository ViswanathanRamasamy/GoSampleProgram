/*
Prime Number: sample program of Golang for the beginers.
Source Code Location: https://github.com/ViswanathanRamasamy/GoSampleProgram.git
Author Name: Viswanathan Ramasamy
Email id: rviswawt@gmail.com
*/
//====================================================================
//println will add the space between each arguments
//print will add space only between the arguments if they are strings
//unused variable and packages are error in golang
//

package main
import "fmt"

func main() {
    var input int
    var status bool
    itemCount,err := fmt.Scanln(&input)
    
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Print(itemCount , ".", "Given number ",input)
    
    for i:=2; i< input;i++ {
        if input%i == 0 {
            status = true
            break
        }
    }
    
    if status == false {
        fmt.Println(input, " Number is prime")
    } else {
         fmt.Println(input, " Number is not prime")
    }
}
===========================
To find the prime number using the math notation:
package main
import "fmt"
import "math"

func main() {
    var input int
    var status bool
    itemCount,err := fmt.Scanln(&input)
    
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Print(itemCount , ".", "Given number ",input, "\n")
    
    if input %2 !=0 || input %3 != 0 {
        for i:=5; i< int(math.Sqrt(float64(input)));i =i+6 {
            if input%i == 0 || input%(i+2) == 0 {
                status = true
                break
            }
        }
    }
    
    if status == false {
        fmt.Println(input, " Number is prime")
    } else {
         fmt.Println(input, "Number is not prime")
    }
}