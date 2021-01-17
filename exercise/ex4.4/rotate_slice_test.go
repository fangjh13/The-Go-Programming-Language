package rotate

import (
    "testing"
    "reflect"
)

func TestRotate(t *testing.T) {
    tests := []struct{
        a, want []int
    }{
        {[]int{1, 2, 3, 4}, []int{2, 3, 4, 1}},
    }
    for _, test := range tests {
        rotateSlice(test.a)
        if !reflect.DeepEqual(test.a, test.want) {
            t.Errorf("got %v, want %v", test.a, test.want)
        }
    }
}
