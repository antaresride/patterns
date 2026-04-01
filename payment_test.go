package main

import (
	"testing"
)

func TestPayee(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}

}

func Repeat(character string) string {
	var repeated string
	for i := 0; i < 4; i++ {
		repeated += character
	}
	return repeated
}
