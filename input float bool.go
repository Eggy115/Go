package main
import "fmt"

func main() {
  var temperature float32
  var sunny bool

  // take float input
  fmt.Println("Enter the current temperature:")
  fmt.Scanf("%g", &temperature)

  // take boolean input
  fmt.Println("Is the day sunny?")
  fmt.Scanf("%t", &sunny)

  fmt.Printf("Current temperature: %g\nIs the day Sunny? %t", temperature, sunny)

}
