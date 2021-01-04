package main

import (
    "bytes"
    "fmt"
    "os"
)

func main() {
    for i := 1; i < len(os.Args); i++ {
        fmt.Printf("%s\n", comma(os.Args[i]))
    }
}

func comma(s string) string {
    buf := &bytes.Buffer{}
    x := len(s) % 3
    // 如果长度是3的倍数
    if x == 0 {
       x = 3
    }
    // 写第一个逗号前的字符
    buf.WriteString(s[:x])
    for i := x; i < len(s); i+=3 {
        buf.WriteByte(',')
        buf.WriteString(s[i:i+3])
    }
    return buf.String()
}
