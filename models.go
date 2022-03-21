package girya

import (
	"strconv"
)

// Some symbols below can be changes depending on a language preference.

var (
	TonnesSymbol    = "t"
	KilogramsSymbol = "kg"
)

// Weight is a ready-to use ~uint32 to represent mass in SI base units.
type Weight uint32

func (w Weight) String() string {
	if w < 1000 {
		return strconv.FormatUint(uint64(w), 10) + " " + KilogramsSymbol
	}
	result := strconv.FormatFloat(float64(w)/1000, 'f', -1, 64)

	var index int
	for key, value := range result {
		if value == '.' {
			index = key
		}
	}
	// Return as is as dot does not present.
	if index == 0 {
		return result + " " + TonnesSymbol
	}

	// Looking for the place to cut.
	cutHere := len(result) - index + 1
	switch cutHere {
	case 0, 1, 2:
	case 3:
		cutHere -= 1
	default:
		cutHere = index + 3
	}

	return result[:cutHere] + " " + TonnesSymbol
}
