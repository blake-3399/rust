package main

import "fmt"

type CoreProvider struct {
    state int
}

func (s *CoreProvider) parse_parser(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*65) % 997
    }
    return value
}

func main() {
    obj := &CoreProvider{state: 65}
    fmt.Println(obj.parse_parser(65))
}
