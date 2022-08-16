package main

import "testing"

var tests = []struct {
	name     string
	divident float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-data", 100.0, 10.0, 10.0, false},
	{"invalid-data", 100.0, 0.0, 0, true},
}

func TestDivison(t *testing.T) {
	for _, tt := range tests {
		got, err := devide(tt.divident, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected error but did not get one")
			}
		} else {
			if err != nil {
				t.Error("expected error but did get one", err.Error())
			}
		}

		if got != tt.expected {
			t.Errorf("expected %f but got %f", tt.expected, got)
		}
	}
}

func TestDevide(t *testing.T) {
	_, err := devide(10.0, 1.0)

	if err != nil {
		t.Error("Got error")
	}

}

func TestBadDevide(t *testing.T) {
	_, err := devide(10.0, 0)

	if err == nil {
		t.Error("Did not get error")
	}

}
