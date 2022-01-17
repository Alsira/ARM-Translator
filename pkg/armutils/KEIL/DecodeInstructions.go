package armutils

import (
	"../internal/binaryutils"
)

func getConditionalCode(nibble int32) string {

	/* Don't bother, this is really long and will return */
	switch nibble {

	case 0000:
		return "EQ"
		break

	case 0001:
		return "NE"
		break

	case 0010:
		return "CS"
		break

	case 0011:
		return "CC"
		break

	case 0100:
		return "MI"
		break

	case 0101:
		return "PL"
		break

	case 0110:
		return "VS"
		break

	case 0111:
		return "VC"
		break

	case 1000:
		return "HI"
		break

	case 1001:
		return "LS"
		break

	case 1010:
		return "GE"
		break

	case 1011:
		return "LT"
		break

	case 1100:
		return "GT"
		break

	case 1101:
		return "LE"
		break

	case 1110:
		return "AL"
		break

	case 1111:
		return "NV"
		break

	default:
		return ""

	}

	return ""

}

// Returns the string for the name of the opcode
func getDataOpCode(binary int) string {

	// Get the nibble for the op-code
	nibble := GetBitRange(binary, 21, 24) >> 21

	switch nibble {

	case 0000:
		return "AND"
		break

	case 0001:
		return "EOR"
		break

	case 0010:
		return "SUB"
		break

	case 0011:
		return "RSB"
		break

	case 0100:
		return "ADD"
		break

	case 0101:
		return "ADC"
		break

	case 0110:
		return "SBC"
		break

	case 0111:
		return "RSC"
		break

	case 1000:
		return "TST"
		break

	case 1001:
		return "TEQ"
		break

	case 1010:
		return "CMP"
		break

	case 1011:
		return "CMN"
		break

	case 1100:
		return "ORR"
		break

	case 1101:
		return "MOV"
		break

	case 1110:
		return "BIC"
		break

	case 1111:
		return "MNV"
		break

	default:
		return ""
	}

	return ""

}

/* Not too sure how to do this, I don't think I could */
/* Decodes Branch Instruction */
func decodeBranchInstruction(binary string) string {

	return ""

}

func decodeDataInstruction(binary string) string {

}
