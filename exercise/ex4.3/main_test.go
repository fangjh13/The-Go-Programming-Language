package main

import (
    "testing"
    "reflect"
)

func TestReverse(t *testing.T) {
    tests := []struct{
        a, want [6]int
    }{
        {[6]int{1, 2, 3, 4, 5, 6}, [6]int{6, 5, 4, 3, 2, 1}},
    }
    for _, test := range tests {
        reverse(&test.a)
        if !reflect.DeepEqual(test.a, test.want) {
            t.Errorf("got %v, want %v", test.a, test.want)
        }
    }
}
