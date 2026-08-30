package main

import "fmt"

type FastContext struct {
    state int
}

func (s *FastContext) render_buffer(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*43) % 997
    }
    return total
}

func main() {
    obj := &FastContext{state: 43}
    fmt.Println(obj.render_buffer(43))
}
