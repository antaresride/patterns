package main

import (
	"strings"
	"testing"
)

//Todo test payment for every customer

func TestPayee(t *testing.T) {
	repeated := Payee("a")
	expected := "aaaa"
	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}

}

func Payee(character string) string {
	var repeated strings.Builder
	for range 4 {
		repeated.WriteString(character)
	}
	return repeated.String()
}
