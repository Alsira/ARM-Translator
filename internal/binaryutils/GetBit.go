package binaryutils

// Gets the bit in a number
func GetBit(binary int, bit_number int) int {
	return (binary >> bit_number) & 2
}