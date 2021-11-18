package armutils

import (

)

/* This changes a base 16 string to a binary string */
func base16ToBinary(hex string) string {

	/* Maps hexadecimal to binary nibble */
	hex_map := make(map[uint8]string)
	hex_map['0'] = "0000"
	hex_map['1'] = "0001"
	hex_map['2'] = "0010"
	hex_map['3'] = "0011"
	hex_map['4'] = "0100"
	hex_map['5'] = "0101"
	hex_map['6'] = "0110"
	hex_map['7'] = "0111"
	hex_map['8'] = "1000"
	hex_map['9'] = "1001"
	hex_map['A'] = "1010"
	hex_map['B'] = "1011"
	hex_map['C'] = "1100"
	hex_map['D'] = "1101"
	hex_map['E'] = "1110"
	hex_map['F'] = "1111"

	/* Slice that holds the string */
	binary_str := make([]string, 4 * len(hex))

	/* Go through each position in the string and produce the nibble */
	for _, hex_digit := range hex {

		

	}



}

/* Takes a hex code string and returns the binary 32 bit machine code for the ARM instruction */
func HexToMachineCode(code string) int32 {

	



}

