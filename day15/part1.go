package main

func RunPart1(inputs []string) int {
	total := 0

	for _, input := range inputs {
		total += Hash(input)
	}
	return total
}
