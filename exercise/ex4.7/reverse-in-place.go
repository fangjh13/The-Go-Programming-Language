package main

import (
    "fmt"
    "unicode/utf8"
)

func rev(s []byte) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

func reverse(s []byte) []byte {
    for i := 0; i < len(s); {
        _, size := utf8.DecodeRune(s[i:])
        rev(s[i:i+size])
        i += size
    }
    rev(s)
    return s
}

func main() {
    fmt.Printf("%s\n", string(reverse([]byte("hello, 世界"))))
}
