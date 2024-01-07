package main

func Hash(input string) int {
	currentVal := 0
	for _, char := range input {
		charVal := int(char)
		currentVal += charVal
		currentVal *= 17
		currentVal %= 256
	}
	return currentVal
}
