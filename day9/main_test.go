package main

import "testing"

func TestRunPart1(t *testing.T) {
	histories := [][]int{
		{0, 3, 6, 9, 12, 15},
		{1, 3, 6, 10, 15, 21},
		{10, 13, 16, 21, 30, 45},
	}
	expected := 114
	output := RunPart1(histories)
	if output != expected {
		t.Errorf("Expected output to be: %d, received %d", expected, output)
	}
}
