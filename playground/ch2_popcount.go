package main

import "fmt"

var pc [256]byte



// PopCount returns the population count (number of set bits) of x.
func PopCountTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

func main() {
    fmt.Println(byte(333))
    fmt.Println(PopCountTable(8999))
}
