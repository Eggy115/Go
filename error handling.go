func compute() (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic: %v", r)
        }
    }()
    if someCondition() {
        panic("something went wrong")
    }
    result = 42
    return result, nil
}

func main() {
    result, err := compute()
    if err != nil {
        log.Fatalf("Error: %v", err)
    }
    fmt.Println(result)
}
