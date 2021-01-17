package shadiff

import (
    "crypto/sha256"
)

func popCount(b byte) int {
    count := 0
    for ; b != 0; count++ {
        b &= b - 1
    }
    return count
}


func bitDiff(a, b []byte) int {
    count := 0
    for i := 0; i < len(a) || i < len(b); i++ {
        switch {
        case i >= len(a):
            count += popCount(a[i])
        case i >= len(b):
            count += popCount(b[i])
        default:
            count += popCount(a[i] ^ b[i])
        }
    }
    return count
}

func ShaBitDiff(a, b []byte) int {
    shA := sha256.Sum256(a)
    shB := sha256.Sum256(b)
    return bitDiff(shA[:], shB[:])
}
