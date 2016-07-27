package main

import (
	"fmt"
	"testing"
)

func TestReplaceString(t *testing.T) {
	find := "q264"
	replace := "HXXX"
	s := "121234123l4j123l4j123l4j234q264lskdjfsldfjsdlfjkq264"
	data := []byte(s)
	replaceString(find, replace, &data)
	fmt.Printf("before: %s\n", s)
	fmt.Printf("after:  %s\n", string(data))
}
