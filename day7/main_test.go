package main

import "testing"

func TestPart1(t *testing.T) {
	inputs := []deal{
		{hand: "32T3K", bid: 765},
		{
			hand: "T55J5",
			bid:  684,
		},
		{
			hand: "KK677",
			bid:  28,
		},
		{
			hand: "KTJJT",
			bid:  220,
		},
		{
			hand: "QQQJA",
			bid:  483,
		},
	}
	expected := 6440
	output := RunPart1(inputs)
	if expected != output {
		t.Errorf("Expected total to be: %d, received: %d", expected, output)
	}
}
