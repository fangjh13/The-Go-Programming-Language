package main

import "fmt"

func unique(s []string) []string {
    i := 0
    for _, w := range s {
        if s[i] == w {
            continue
        }
        i++
        s[i] = w
    }
    return s[:i+1]
}

func main() {
    s := unique([]string{"a", "a", "b", "c", "c", "c", "d", "d", "e"})
    fmt.Println(s)
}
