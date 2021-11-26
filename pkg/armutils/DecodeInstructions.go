package armutils

getConditionalCode(nibble string) {

	conditions := make(map[string]string)

	conditions["0000"] = "EQ"
	conditions["0001"] = "NE"
	conditions["0010"] = "CS"
	conditions["0011"] = "CC"
	conditions["0100"] = "MI"
	conditions["0101"] = "PL"
	conditions["0110"] = "VS"
	conditions["0111"] = "VC"
	conditions["1000"] = "HI"
	conditions["1001"] = "LS"
	conditions["1010"] = "GE"
	conditions["1011"] = "LT"
	conditions["1100"] = "GT"
	conditions["1101"] = "LE"
	conditions["1110"] = "AL"
	conditions["1111"] = "NV"

	return conditions[nibble]

}

/* Not too sure how to do this, I don't think I could */
/* Decodes Branch Instruction */
decodeBranchInstruction(binary string) string {

	var instruction string


}

decodeDataInstruction(binary string) string {

	

}