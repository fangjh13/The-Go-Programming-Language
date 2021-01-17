package main

import (
    "testing"
    "reflect"
)

func TestUnique(t *testing.T) {
    input := []string{"a", "a", "b", "b", "c"}
    want := []string{"a", "b", "c"}
    got := unique(input)
    if !reflect.DeepEqual(want, got) {
        t.Errorf("got %v, want %v", got, want)
    }
}
