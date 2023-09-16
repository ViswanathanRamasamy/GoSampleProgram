/*
Lesson6: sample program of Golang for the beginers.
Source Code Location: https://github.com/ViswanathanRamasamy/GoSampleProgram.git
*/
================================================================================
defer function for closing the file:
package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("myfile.txt") // no r or w
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close() // Ensure the file is closed when the function exits of main
    // Read from the file
	// .....
}
======================================================
unlocing the mutex:
package main

import (
    "fmt"
    "sync"
)

func main() {
    var mu sync.Mutex
    mu.Lock()
    defer mu.Unlock() // Ensure the mutex is unlocked when the function exits
    // Critical section code
	//.....
}
=====================================================
Timing Function Execution:
package main

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()
    defer func() {
        fmt.Println("Time taken:", time.Since(start))
    }()
    // Code to measure execution time
    for i := 0; i < 1000000; i++ {
        // Some operation
    }
}
===============================
panic error:
package main

import (
    "fmt"
)

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    // Code that may panic
    panic("This will cause a panic")
}
==============================
defer with function argument:
package main

import (
    "fmt"
)

func main() {
    x := 10
    defer fmt.Println("The value of x:", x)
    x = 20
    // Output will be "The value of x: 10"
}
