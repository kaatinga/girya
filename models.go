package girya

import (
	"strconv"
)

// Some symbols below can be changes depending on a language preference.
var (
	TonnesPostfix    = " t"
	KilogramsPostfix = " kg"
)

// Weight is a ready-to use type to represent mass in SI base units.
type Weight int64

func (weight Weight) String() string {
	value := int64(weight)
	var sign string
	if value < 0 {
		sign = "-"
		value = -value
	}

	if value < 1000 {
		return sign + strconv.FormatInt(value, 10) + KilogramsPostfix
	}
	if value > 50000 {
		return sign + strconv.FormatInt(value/1000, 10) + TonnesPostfix
	}

	centi := value / 10
	intPart := centi / 100
	frac := centi % 100

	buf := make([]byte, 0, len(sign)+8+len(TonnesPostfix))
	buf = append(buf, sign...)
	buf = strconv.AppendInt(buf, intPart, 10)
	if frac != 0 {
		buf = append(buf, '.', byte('0'+frac/10))
		if frac%10 != 0 {
			buf = append(buf, byte('0'+frac%10))
		}
	}
	buf = append(buf, TonnesPostfix...)
	return string(buf)
}
