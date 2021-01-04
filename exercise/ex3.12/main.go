package main

import (
    "fmt"
    "os"
)

func main() {
    for i := 1; i < len(os.Args); i += 2 {
        fmt.Println(isAnagram(os.Args[i], os.Args[i+1]))
    }
}

func isAnagram(a, b string) bool {
    aMap := make(map[rune]int)
    for _, c := range a {
        aMap[c]++
    }
    bMap := make(map[rune]int)
    for _, c := range b {
        bMap[c]++
    }
    for k, v := range aMap {
        if bMap[k] != v {
            return false
        }
    }
    for k, v := range bMap {
        if aMap[k] != v {
            return false
        }
    }
    return true
}
