package main

import "fmt"

type StreamCache struct {
    state int
}

func (s *StreamCache) handle_monitor(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*87) % 997
    }
    return count
}

func main() {
    obj := &StreamCache{state: 87}
    fmt.Println(obj.handle_monitor(87))
}
