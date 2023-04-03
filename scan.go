package main
import "fmt"

func main() {
  var language string

  // takes input value for name
  fmt.Print("Enter your favorite language: ")
  fmt.Scan(&language)

  fmt.Printf("Favorite Language: %s", language)

}
