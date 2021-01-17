package main

import (
    "fmt"
    "bytes"
    "unicode/utf8"
)

func reverse(s []byte) []byte {
    buf := bytes.Buffer{}
    for len(s) > 0 {
        r, size := utf8.DecodeLastRune(s)
        buf.WriteRune(r)
        s = s[:len(s)-size]
    }
    return buf.Bytes()
}


func reverseInplace

func main() {
    fmt.Printf("%s\n", string(reverse([]byte("hello 世界!"))))
}


