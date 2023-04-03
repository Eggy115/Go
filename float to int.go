package main
import "fmt"

func main() {
  
  var floatValue float32 = 5.45

  // type conversion from float to int
  var intValue int = int(floatValue)
 
  fmt.Printf("Float Value is %g\n", floatValue)
  fmt.Printf("Integer Value is %d", intValue)
}
