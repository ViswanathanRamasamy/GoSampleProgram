/*
Lesson4: sample program of Golang for the beginers.
Source Code Location: https://github.com/ViswanathanRamasamy/GoSampleProgram.git
*/
// Program to write the content to the file and read the file.
//defer keyword ensure that all the operations are completed before executing the defer command
// only the create file we are clsoing the file
//ioutil.Readstring no close required.

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func errorCheck(err error) {
	if err != nil {
		panic(err) //panic error can have only one message argument
	}
}
func openandWrite(name string) {
	var content string = "Hellow World"
	file, error := os.Create(name)
	errorCheck(error)
	len, error := io.WriteString(file, content)
	errorCheck(error)
	fmt.Println("print length written", len)
	defer file.Close()
}

func ReadFile(name string) {
	data, error := ioutil.ReadFile(name) //will read the file and put it in array
	errorCheck(error)
	fmt.Println(string(data)) //convert the array to string
}

func main() {
	var filename string = "./viswa.txt"
	openandWrite(filename)
	defer ReadFile(filename)

}

/*
 net/http => get
 io => parse the response body
 encoding/json => parse the response in json
 note error: = in multiple places

 response, error := http.Get()

 "%T", response ==> http response

 defer response.Body.Close()
 rawByte, error := ioutil.ReadAll(response.Body)
 string(rawByte)

    function going to return the slice of user struct
   func jsonreader(content string) [] Tour  {
       var tour Tour;
	   tours :=make([]Tour, 10,10);
	   decoder := json.NewDecoder(strings.NewReader(content))
	   to check whether there is any error
	   _,err := decoder.Token()
	   for decoder.More() {
		err := decoder.decoder(tour)
         tours.append(tours,tour)
	   }
	   return tours
   }

   type Tour struct {
	name , Add string
   }

   http.Get() return the pointer
*/
