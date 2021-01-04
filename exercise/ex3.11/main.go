package main

import (
    "bytes"
    "fmt"
    "os"
    "strings"
)

func main() {
    for i := 1; i < len(os.Args); i++ {
        fmt.Printf("%s\n", comma(os.Args[i]))
    }
}

/* 
func comma(s string) string {
    if len(s) < 3 {
        return s
    }
    sStart, sEnd := 0, len(s)
    buf := &bytes.Buffer{}
    // 处理正负号
    if s[0] == '+' || s[0] == '-' {
        buf.WriteByte(s[0])
        sStart = 1
    }
    // 处理小数点
    dotIndex := strings.Index(s, ".")
    if dotIndex != -1 {
        sEnd = dotIndex
    }
    // 小数点和符号中间的部分
    targetS := s[sStart:sEnd]
    targetL := len(targetS)
    j := targetL % 3
    if j == 0 {
        j = 3
    }
    buf.WriteString(targetS[:j])
    for i := j; i < targetL; i += 3 {
        buf.WriteByte(',')
        buf.WriteString(targetS[i: i+3])
    }
    // 追加小数部分
    buf.WriteString(s[sEnd:])
    return buf.String()
}
*/

func comma(s string) string {
    if len(s) < 3 {
        return s
    }
    sStart, sEnd := 0, len(s)
    buf := &bytes.Buffer{}
    // 处理正负号
    if s[0] == '+' || s[0] == '-' {
        buf.WriteByte(s[0])
        sStart = 1
    }
    // 处理小数点
    dotIndex := strings.Index(s, ".")
    if dotIndex != -1 {
        sEnd = dotIndex
    }
    // 小数点和符号中间的部分
    targetS := s[sStart:sEnd]
    targetL := len(targetS)
    j := targetL % 3
    if j > 0 {
        buf.Write([]byte(targetS[:j]))
        if targetL > j {
            buf.WriteByte(',')
        }
    }
    for i, c := range targetS[j:] {
        if i % 3 == 0 && i != 0 {
            buf.WriteByte(',')
        }
        buf.WriteRune(c)
    }
    // 追加小数部分
    if dotIndex != -1 {
        buf.WriteByte('.')
        buf.WriteString(s[dotIndex+1:])
    }
    return buf.String()
}
