package rotate

func rotateSlice(s []int) {
    first := s[0]
    copy(s, s[1:])
    s[len(s)-1] = first
}
