package main

import (
	"testing"
)

/*
command:
go test -v
go test -cover
*/
var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-valid-data", 100.0, 0.0, 0.0, true},
}

func TestDivision(t *testing.T) {
	for _, test := range tests {
		got, err := divide(test.dividend, test.divisor)
		if test.isErr {
			if err == nil {
				t.Error("Expected error, got nil")
			}
		} else {
			if err != nil {
				t.Error("Did not expect error, got nil", err.Error())
			}
		}

		if got != test.expected {
			t.Error("Expected ", test.expected, " got ", got)
		}
	}
}

func TestDivide(t *testing.T) {
	_, err := divide(10.0, 1.0)
	if err != nil {
		t.Error("Got an error when we should not have gotten an error")
	}
}

func TestBadDivide(t *testing.T) {
	_, err := divide(10.0, 0.0)
	if err == nil {
		t.Error("Dig not get an error when we should have gotten an error")
	}
}
