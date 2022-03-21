package girya

import (
	"fmt"
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
	return fmt.Sprintf("%.2f %s", float64(w)/1000, TonnesSymbol)
}
