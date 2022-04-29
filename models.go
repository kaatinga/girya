package girya

import (
	"strconv"
)

// Some symbols below can be changes depending on a language preference.
var (
	TonnesSymbol    = "t"
	KilogramsSymbol = "kg"
)

// Weight is a ready-to use type to represent mass in SI base units.
type Weight int64

func (weight Weight) String() string {
	var sign = ""
	if weight < 0 {
		weight = 0 - weight
		sign = "-"
	}
	if weight < 1000 {
		return sign + strconv.FormatInt(int64(weight), 10) + " " + KilogramsSymbol
	}
	if weight > 50000 {
		weight = weight / 1000 * 1000
	} else {
		weight = weight / 10 * 10
	}
	result := strconv.FormatFloat(float64(weight)/1000, 'f', -1, 64)

	var index int
	for key, value := range result {
		if value == '.' {
			index = key
		}
	}
	// Return as is as fractional part is not present.
	if index == 0 {
		return sign + result + " " + TonnesSymbol
	}

	// Looking for the place to cut.
	cutHere := len(result) - index + 1
	switch cutHere {
	case 0, 1, 2:
	case 3:
		cutHere = index + 2
	default:
		cutHere = index + 3
	}

	return sign + result[:cutHere] + " " + TonnesSymbol
}
