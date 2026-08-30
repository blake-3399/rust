package main

import "fmt"

type StreamLoader struct {
    state int
}

func (s *StreamLoader) encode_gateway(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*75) % 997
    }
    return result
}

func main() {
    obj := &StreamLoader{state: 75}
    fmt.Println(obj.encode_gateway(75))
}
