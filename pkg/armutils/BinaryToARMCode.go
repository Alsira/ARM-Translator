package armutils

import (
	"strings"
	"math"
)

/* This converts the 32 binary to the */
func BinaryToARMCode(binary_str string) string {

	/* Binary string is stored where the MSB is the first index in the string */
	/* Therefore I will flip the string in reverse */
	
	bstr_len := len(binary_str)

	for i, j := 0, bstr_len - 1; i < int32(math.Ceil(bstr_len / 2.0f)); i, j = i + 1, j - 1 {
		binary_str[i], binary_str[j] = binary_str[j], binary_str[i]
	}

	/* Check if branch instruction */
	if (binary_str[27] == '1' && binary_str[26] == '0' && binary_str[25] == '1') {
	
		decodeBranchInstruction(binary_str)	
		
		/* Check if ARM data processor instruction */
	} else if (binary_str[27] == '0' && binary_str[26] == '0') {

		decodeDataInstruction(binary_str)

	} else if (binary_str[27] == '0' && binary_str[26] == '0' && binary_str[25] == '0') {

		decodeBlockMoveInstruction(binary_str)

	} else if (binary_str[27] == '1' && binary_str[26] == '0' && binary_str[25] == '1') {

		decodeLSInstruction(binary_str)

	}
	



}