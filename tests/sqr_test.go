package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, but got  %d", expected, result)
	}
}

func TestIsEven(t *testing.T) {
	/// test odd
	result := IsEven(51)
	expected := false
	if result != expected {
		t.Errorf("Expected %v for odd number, but got %v", expected, result)
	}

	/// even
	result2 := IsEven(50)
	expected2 := true
	if result2 != expected2 {
		t.Errorf("Expected %v for even number, but got %v", expected, result)
	}

}
