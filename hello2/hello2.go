/*
Declare a main package. In Go, code executed
 as an application must be in a main package.
*/
package main

/*
Import two packages: example.com/greetings 
and the fmt package. This gives your code 
access to functions in those packages. 
Importing example.com/greetings 
(the package contained in the module you created
 earlier) gives you access to the Hello
  function. You also import fmt, with 
  functions for handling input and output
   text (such as printing text to the console).
*/
import (
    "fmt"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Jarb")
    fmt.Println(message)
}