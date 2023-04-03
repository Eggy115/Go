// Program to take input from user using Scanf()

package main
import "fmt"

func main() {
  var name string
  var age int

  // take name and age input using format specifier
  fmt.Println("Enter your name and age:")
  fmt.Scanf("%s %d", &name, &age) 

  fmt.Printf("Name: %s\nAge: %d", name, age)
}
