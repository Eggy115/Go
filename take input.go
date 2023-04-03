package main
import "fmt"

func main() {
  var name string

  // takes input value for name
  fmt.Print("Enter your name: ")
  fmt.Scan(&name)

  fmt.Printf("Name: %s", name)

}
