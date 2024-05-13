package funcs

func Math(input string) []int { // convert to ascii values
	asciiValues := make([]int, 0)

	for _, ch := range input {
		asciiValue := int(ch)
		getascii := (asciiValue-32)*9 + 2
		asciiValues = append(asciiValues, getascii)
	}
	return asciiValues
}
