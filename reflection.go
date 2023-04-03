package main
import "fmt"

type Person struct {
    Name string
    Age  int
}

func printFields(v interface{}) {
    value := reflect.ValueOf(v)
    for i := 0; i < value.NumField(); i++ {
        field := value.Field(i)
        fmt.Printf("%s = %v\n", reflect.TypeOf(field).Name(), field.Interface())
    }
}

func main() {
    p := Person{Name: "John", Age: 30}
    printFields(p)
}
