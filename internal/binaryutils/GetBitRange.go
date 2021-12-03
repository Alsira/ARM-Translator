package binaryutils

// Returns the bits in the given ranges, the rest are zeroed out
func GetBitRange(binary, lower, higher int) int {

	var returnNum int = 0

	// Here we'll loop through and add each bit
	for index := lower; index <= higher; index++ {

		returnNum += GetBit(binary, index) << index

	}

	return returnNum

}