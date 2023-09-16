/*
Lesson5: sample program of Golang for the beginers.
Source Code Location: https://github.com/ViswanathanRamasamy/GoSampleProgram.git
*/
================================================================================
/* Note
To pass the channel to a function

ch1 chan int or ch1 chan<-int
for data:= range chan1 //will break only when the channel is closed
*/

In Go, a goroutine is a lightweight, independently executing thread of control within a Go program. Goroutines are a fundamental feature of Go's concurrency model,
 and they are used for concurrent and parallel execution of code. Goroutines are managed by the Go runtime, and they enable you to perform tasks concurrently without the complexity of managing threads manually.
concurrency is achieved using the go routines:

package main

import (
    "fmt"
    "time"
)

func outsideGoroutine() {
    for i := 0; i < 3; i++ {
        fmt.Println("Outside Goroutine:", i)
        time.Sleep(time.Millisecond * 500)
    }
}

func main() {
    // Creating a goroutine outside the main function
    go outsideGoroutine()

    for i := 0; i < 3; i++ {
        fmt.Println("Inside Main:", i)
        time.Sleep(time.Millisecond * 500)
    }

    // Sleep for a while to allow the goroutine to complete
    time.Sleep(time.Second)

    fmt.Println("Main function completed.")
}
=====================================
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    for i := 0; i < 5; i++ {
        fmt.Println("Hello from the goroutine")
        time.Sleep(time.Millisecond * 500)
    }
}

func main() {
    // Create and start a goroutine
    go sayHello()

    // Main function continues to execute concurrently with the goroutine
    for i := 0; i < 5; i++ {
        fmt.Println("Hello from the main function")
        time.Sleep(time.Millisecond * 500)
    }

    // Allow some time for goroutine to finish
    time.Sleep(time.Second)
}

Hello from the main function
Hello from the goroutine
Hello from the main function
Hello from the goroutine
Hello from the main function
Hello from the goroutine
Hello from the main function
Hello from the goroutine
Hello from the main function
Hello from the goroutine
==========================================
go routine inside the main:
one routine calling another routine function:
package main

import (
    "fmt"
    "sync"
)

func sayHello() {
    fmt.Println("Hello from the goroutine")
}

func main() {
    // Create a wait group to wait for the goroutine to finish
    var wg sync.WaitGroup

    // Add 1 to the wait group to track one goroutine
    wg.Add(1)

    // Start a goroutine
    go func() {
        defer wg.Done() // Mark the goroutine as done when it finishes
        sayHello()
    }()

    // Main function continues to execute concurrently with the goroutine
    fmt.Println("Hello from the main function")

    // Wait for the goroutine to finish using the wait group
    wg.Wait()
}
Hello from the main function
Hello from the goroutine
=====================================================================================
//passing wg and channel to the routine?

go prallelism
Go take care of the program to run different cores.
sync package provides synchronization primitives that help manage concurrent access to shared data, coordinate the execution of goroutines,
 and avoid race conditions. 
package main

import (
	"fmt"
	"sync"
)

func squareWorker(wg *sync.WaitGroup, num int, results chan<- int) {
	defer wg.Done()
	result := num * num
	results <- result
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	results := make(chan int, len(nums))
	var wg sync.WaitGroup

	// Start a worker for each number
	for _, num := range nums {
		wg.Add(1)
		go squareWorker(&wg, num, results)
	}

	// Wait for all workers to finish
	wg.Wait()

	// Close the results channel after all workers are done
	close(results)

	// Collect and print the results
	for result := range results {
		fmt.Printf("Square: %d\n", result)
	}
}

==========================
//init is the special function that will get executed before the main function is getting executed.

package main

import (
    "fmt"
)

func init() {
    fmt.Println("This is the init function of the main package.")
}

func main() {
    fmt.Println("This is the main function.")
}

output:
This is the init function of the main package
This is the main function
=================================================
package main

import (
    "fmt"
    "sync"
)

func someFunction(wg *sync.WaitGroup) {
   
    fmt.Println("This function runs concurrently with main.")
	 defer wg.Done() // Mark the function as completed when done
}

func main() {
    var wg sync.WaitGroup

    // Create a WaitGroup to wait for the concurrent function to finish
    wg.Add(1) //one concurrent 

    // Start the concurrent function
    go someFunction(&wg)

    // Continue with main function's logic

    fmt.Println("This is the main function.")

    // Wait for the concurrent function to finish
    wg.Wait()

    fmt.Println("This is executed after the concurrent function.")
}

This is the main function.
This function runs concurrently with main.
This is executed after the concurrent function.

=======================================================
When you call wg.Add(n), you are incrementing an internal counter within the WaitGroup, indicating that you expect n goroutines to be started concurrently.
 When each of those goroutines is done with its work, it calls wg.Done() to 
decrement the counter. The wg.Wait() method in the main function will block until the counter reaches zero, indicating that all the goroutines have completed.
package main

// response.body.close()

import (
    "fmt"
    "net/http"
    "sync"
)

func downloadURL(url string, wg *sync.WaitGroup) {
    defer wg.Done()

    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("Error downloading %s: %v\n", url, err)
        return
    }
    defer resp.Body.Close()

    fmt.Printf("Downloaded %s, Status: %s\n", url, resp.Status)
}

func main() {
    var wg sync.WaitGroup

    urls := []string{
        "https://www.example.com",
        "https://www.google.com",
        "https://www.github.com",
    }

    for _, url := range urls {
        wg.Add(1)
        go downloadURL(url, &wg)
    }

    wg.Wait()
    fmt.Println("All downloads are complete.")
}

Downloaded https://www.example.com, Status: 200 OK
Downloaded https://www.google.com, Status: 200 OK
Downloaded https://www.github.com, Status: 200 OK
All downloads are complete.

===============================
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Worker %d started\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d completed\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    wg.Wait()
    fmt.Println("All workers have completed.")
}


Worker 1 started
Worker 2 started
Worker 3 started
Worker 4 started
Worker 5 started
Worker 4 completed
Worker 5 completed
Worker 1 completed
Worker 3 completed
Worker 2 completed
All workers have completed.
===============================================
//It is not necessary to close the channel before read?
//one dat will be read only by one receiver.

package main

import (
    "fmt"
)

func main() {
    // Create a buffered channel with a capacity of 2
    ch := make(chan int, 2)

    // Send two elements to the channel
    ch <- 1
    ch <- 2

    // Attempt to send a third element (will block until space is available)
    // ch <- 3

    // Close the channel to indicate that no more data will be sent
    close(ch)

    // Receive data from the channel
    for data := range ch {
        fmt.Println("Received:", data)
    }
}
===========================================================
Closing the Channel Immediately:
//dat can be read from the channel using :=range channel

If the channel is closed immediately after being created in the sender goroutine, the receiver will not have a chance to receive any data. 
In this case, the receiver will detect that the channel is already closed and exit without receiving any data.

Here's an example:


import (
    "fmt"
)

func main() {
    // Create a buffered channel with a capacity of 2 and close it immediately
    ch := make(chan int, 2)
    close(ch)

    // Attempt to receive data from the closed channel
    for data := range ch {
        fmt.Println("Received:", data)
    }

    fmt.Println("Receiver done")
}
==============================================\
go  dont have default arguments:
package main

import (
    "fmt"
)

// Option is a struct for specifying optional parameters
type Option struct {
    Param1 int
    Param2 string
}

// defaultOption returns a default Option value
func defaultOption() Option {
    return Option{
        Param1: 42,
        Param2: "default",
    }
}

// MyFunction takes an Option struct with optional parameters
func MyFunction(option Option) {
    fmt.Printf("Param1: %d, Param2: %s\n", option.Param1, option.Param2)
}

func main() {
    // Calling MyFunction with default values
    MyFunction(defaultOption())

    // Calling MyFunction with custom values
    customOption := Option{
        Param1: 100,
        Param2: "custom",
    }
    MyFunction(customOption)
}

=============================================
variadic parameter:
package main

import (
    "fmt"
)

func greet(message string, names ...string) {
    if len(names) == 0 {
        names = append(names, "Guest")
    }

    for _, name := range names {
        fmt.Printf("%s, %s!\n", message, name)
    }
}

func main() {
    greet("Hello")           // Output: Hello, Guest!
    greet("Hi", "Alice")     // Output: Hi, Alice!
    greet("Hola", "Bob", "Charlie") // Output: Hola, Bob! Hola, Charlie!
}

======================================
//How to read all data from the channel ?
In Go, a channel is a built-in concurrency primitive that allows communication and synchronization between goroutines.
Channels are used for sending and receiving data between goroutines, enabling safe concurrent access to shared data.
package main

import (
    "fmt"
    "time"
)

func main() {
    // Create a channel for sending and receiving integers
    ch := make(chan int)

    // Start a goroutine to send data to the channel
    go func() {
        for i := 1; i <= 5; i++ {
            ch <- i // Send integers from 1 to 5 to the channel
        }
        close(ch) // Close the channel to indicate that no more data will be sent
    }()

    // Start a goroutine to receive data from the channel
    go func() {
        for {
            value, ok := <-ch // Receive a value from the channel
            if !ok {
                break // Channel is closed, exit the loop
            }
            fmt.Println("Received:", value)
        }
    }()

    // Allow some time for goroutines to finish
    time.Sleep(time.Second)

    fmt.Println("Main goroutine has completed.")
}
Received: 1
Received: 2
Received: 3
Received: 4
Received: 5
Main goroutine has completed.

====================
2 sender and 3 receiver: Please not eone data send by the sender cannot be received by more than one receiver.

package main

import (
	"fmt"
	"sync"
)

func sender(ch chan int, wg *sync.WaitGroup, senderID int) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		data := i + senderID*10 // Unique data for each sender
		ch <- data              // Send data to the channel
	}
}

func receiver(ch chan int, wg *sync.WaitGroup, receiverID int) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("Receiver %d received: %d\n", receiverID, data)
	}
	fmt.Printf("Receiver %d done\n", receiverID)
}

func main() {
	var senderWg sync.WaitGroup
	var receiverWg sync.WaitGroup
	ch := make(chan int)

	// Start two sender goroutines
	for i := 1; i <= 2; i++ {
		senderWg.Add(1)
		go sender(ch, &senderWg, i)
	}

	// Start three receiver goroutines
	for i := 1; i <= 3; i++ {
		receiverWg.Add(1)
		go receiver(ch, &receiverWg, i)
	}

	// Wait for all senders to finish
	senderWg.Wait()

	// Close the channel to signal that no more data will be sent
	close(ch)

	// Wait for all receivers to finish
	receiverWg.Wait()
}

Receiver 1 received: 11
Receiver 2 received: 21
Receiver 3 received: 31
Receiver 3 received: 32
Receiver 1 received: 12
Receiver 2 received: 22
Receiver 1 received: 13
Receiver 3 received: 33
Receiver 2 received: 23
Receiver 1 received: 14
Receiver 2 received: 24
Receiver 3 received: 34
Receiver 3 received: 35
Receiver 2 received: 25
Receiver 1 received: 15
Receiver 2 received: 26
Receiver 1 done
Receiver 2 done
Receiver 3 done

================================================
package main

import (
    "fmt"
    "time"
)

func main() {
    // Create two channels
    ch1 := make(chan int)
    ch2 := make(chan string)

    // Start a goroutine to send data on ch1
    go func() {
        for i := 1; i <= 3; i++ {
            ch1 <- i
            time.Sleep(time.Millisecond * 200)
        }
        close(ch1)
    }()

    // Start a goroutine to send data on ch2
    go func() {
        for i := 1; i <= 3; i++ {
            ch2 <- fmt.Sprintf("Message %d", i)
            time.Sleep(time.Millisecond * 300)
        }
        close(ch2)
    }()

    // Use select to receive data from both channels
    for {
        select {
        case data, ok := <-ch1:
            if !ok {
                fmt.Println("ch1 closed.")
                ch1 = nil // Set to nil to exclude it from select
            } else {
                fmt.Println("Received from ch1:", data)
            }
        case data, ok := <-ch2:
            if !ok {
                fmt.Println("ch2 closed.")
                ch2 = nil // Set to nil to exclude it from select
            } else {
                fmt.Println("Received from ch2:", data)
            }
        }

        if ch1 == nil && ch2 == nil {
            break // Both channels are closed, exit the loop
        }
    }

    fmt.Println("Main goroutine done.")
}
Received from ch1: 1
Received from ch2: Message 1
Received from ch1: 2
Received from ch1: 3
Received from ch2: Message 2
Received from ch2: Message 3
ch1 closed.
ch2 closed.
Main goroutine done.
=============================================
interface:
package main

import (
    "fmt"
)

// Define an interface
type Shape interface {
    Area() float64
}

// Define a struct for a rectangle
type Rectangle struct {
    Width  float64
    Height float64
}

// Implement the Area method for the Rectangle type
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Define a struct for a circle
type Circle struct {
    Radius float64
}

// Implement the Area method for the Circle type
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    // Create instances of Rectangle and Circle
    rect := Rectangle{Width: 4, Height: 3}
    circle := Circle{Radius: 2}

    // Use the Area method through the Shape interface
    shapes := []Shape{rect, circle}

    for _, shape := range shapes {
        fmt.Printf("Area: %.2f\n", shape.Area())
    }
}\
Area: 12.00
Area: 12.56
==========================================
package main

import (
    "fmt"
)

type Writer interface {
    Write(data string)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data string) {
    fmt.Println(data)
}

func writeToChannel(ch chan string, message string) {
    ch <- message
}

func main() {
    writer := ConsoleWriter{}

    ch := make(chan string)

    go writeToChannel(ch, "Hello, Interface!")
    go writeToChannel(ch, "Direct function call")

    for {
        data, ok := <-ch
        if !ok {
            break // Channel is closed, exit the loop
        }
        writer.Write(data)
    }

    close(ch)
}
Hello, Interface!
Direct function call
dead lock since the receiver will not come out for the for loop as the channel is not closed.
================================================
//it is not necessary to have the receiver for writing the data to the channel. but when the channel has reached the capacity it will block.

package main

import (
	"fmt"
)

type Writer interface { // this interface is not requires for this program
	Write(data string)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data string) {
	fmt.Println(data)
}

func writeToChannel(ch chan string, message string) {
	ch <- message
}

func main() {
	writer := ConsoleWriter{}

	ch := make(chan string)

	go writeToChannel(ch, "Hello, Interface!")
	go writeToChannel(ch, "Direct function call")

	// Create a goroutine to receive and print data from the channel
	go func() {
		for {
			data, ok := <-ch
			if !ok {
				fmt.Println("Channel is closed")
				break
			}
			writer.Write(data)
		}
	}()

	// Sleep for a moment to allow time for data to be sent and received
	//time.Sleep(5 * time.Second)
	// Close the channel when you're done
	close(ch)

	// Sleep to give the receiving goroutine time to print data
	// before the program exits
	//time.Sleep(1 * time.Second)
}
Hello, Interface!
Direct function call
====================================

package main

import (
	"fmt"
	"sync"
)

type Writer interface {
	Write(data string)
}

type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data string) {
	fmt.Println(data)
}

func writeToChannel(ch chan string, message string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- message
}

func main() {
	writer := ConsoleWriter{}

	ch := make(chan string)

	var wg sync.WaitGroup

	wg.Add(2) // Number of goroutines to wait for

	go writeToChannel(ch, "Hello, Interface!", &wg)
	go writeToChannel(ch, "Direct function call", &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for data := range ch { //This will termonate only when the channel is closed
		writer.Write(data)
	}
}

Hello, Interface!
Direct function call

==================================================
defer :LIFO
package main

import "fmt"

func main() {
    fmt.Println("Before defer")
    defer fmt.Println("Deferred 1")
    fmt.Println("Between defer 1 and defer 2")
    defer fmt.Println("Deferred 2")
    fmt.Println("After defer")
}

Before defer
Between defer 1 and defer 2
After defer
Deferred 2
Deferred 1
======================================================

=====================================================
package main

import (
	"fmt"
	"log"
)

func someFunction(value int) (int, error) {
	if value < 0 {
		panic("Value cannot be negative") // Added panic statement
	}
	return value * 2, nil
}

func main() {
	// Example of regular error handling
	result, err := someFunction(5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Trigger a panic by calling someFunction with a negative value. This function will be called when the program is ending
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered:", r)
		}
	}()

	someFunction(-3) // This line will trigger a panic
	fmt.Println("This line will not be reached.")
}

Result: 10
Recovered: Value cannot be negative
==========================================
attactching structire to a function:
In Go, you can define a method on a struct by attaching a function to the struct type. 
This allows you to call the function on instances of the struct, similar to object-oriented programming. The syntax for defining such a method is func
 (s SomeStruct) methodName(argumentType) returnType.

package main

import (
    "fmt"
)

// Define a simple struct
type Rectangle struct {
    Width  float64
    Height float64
}

// Define a method called Area on the Rectangle struct
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    // Create an instance of the Rectangle struct
    rect := Rectangle{
        Width:  5.0,
        Height: 3.0,
    }

    // Call the Area method on the Rectangle struct
    area := rect.Area()

    // Print the result
    fmt.Printf("Area of the rectangle: %.2f square units\n", area)
}
Area of the rectangle: 15.00 square units
============================
Error logging:
log.Fatal to log a fatal error message. This function logs the message and then calls os.Exit(1) to exit the program.
//err.(*MyError).Message, //err
//err.(*MyError).Code
// note here you convert the MyError to the error.
//not multipleWrite it is MultiWrite

package main

import (
	"fmt"
	"log"
	"os"
	"io"
)

// Custom error type
type MyError struct {
	Message string
	Code    int
}

func (e *MyError) Error() string {
	return e.Message
}

// Function that may return a custom error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &MyError{Message: "Division by zero", Code: 1001}
	}
	return a / b, nil
}

func main() {
	// Set up a log file
	logFile, err := os.Create("error.log")
	if err != nil {
		log.Fatal("Error creating log file:", err ) //this will print error message only in the screen
	}
	defer logFile.Close()

	// Set the log output to both stdout and the log file
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	// Example of error handling and logging
	result, err := divide(10.0, 2.0)
	if err != nil {
		log.Printf("Error: %v (Code: %d)", err.(*MyError).Message, err.(*MyError).Code)
	} else {
		fmt.Println("Result:", result)
	}

	// Another example with an error
	result, err = divide(8.0, 0.0)
	if err != nil {
		log.Printf("Error: %v (Code: %d)", err, err.(*MyError).Code)
	} else {
		fmt.Println("Result:", result)
	}
}
Result: 5
Error: Division by zero (Code: 1001)
========================================