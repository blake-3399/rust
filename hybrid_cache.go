package main

import "fmt"

type LocalLoader struct {
    state int
}

func (s *LocalLoader) collect_context(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*67) % 997
    }
    return count
}

func main() {
    obj := &LocalLoader{state: 67}
    fmt.Println(obj.collect_context(67))
}
