package main

import (
    "errors"
    "fmt"
    "unicode/utf8"
)

// The key difference is that Reverse is now iterating over each rune in the string, rather than each byte.
// Note that this is just an example, and does not handle combining characters correctly.
func Reverse(s string) (string, error) {
    if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
    // fmt.Printf("input: %q\n", s)
    r := []rune(s)
    // fmt.Printf("runes: %q\n", r)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r), nil
}

func main() {
    input := "The quick brown fox jumped over the lazy dog"
    rev, revError := Reverse(input)
    doubleRev, doubleRevError := Reverse(rev)
    fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q, err: %v\n", rev, revError)
    fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevError)
}

