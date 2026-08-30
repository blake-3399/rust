package main

import "fmt"

type SharedBuffer struct {
    state int
}

func (s *SharedBuffer) encode_cache(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*64) % 997
    }
    return total
}

func main() {
    obj := &SharedBuffer{state: 64}
    fmt.Println(obj.encode_cache(64))
}
