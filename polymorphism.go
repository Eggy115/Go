type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func main() {
    shapes := []Shape{Rectangle{Width: 3, Height: 4}, Circle{Radius: 5}}
    for _, s := range shapes {
        fmt.Println(s.Area())
    }
}
