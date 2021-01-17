package main

import (
    "fmt"
    "unicode"
    "unicode/utf8"
    "bytes"
)

func removeDupSpace(s string) string {
    pre := false
    result := &bytes.Buffer{}
    for cursor := 0; cursor < len(s); {
        r, size := utf8.DecodeRuneInString(s[cursor:])
        if unicode.IsSpace(r) {
            if !pre {
                result.WriteByte(' ')
            }
            pre = true
        } else {
            pre = false
            result.WriteRune(r)
        }
        cursor += size
    }
    return result.String()
}

func main() {
    s := "abc\t     世界"
    fmt.Println(removeDupSpace(s))
}
