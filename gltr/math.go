package glt

func MaxUint16(a, b uint16) uint16 {
	if a > b {
		return a
	} else {
		return b
	}
}

func clampUint16(lowerLimit, upperLimit, number uint16) uint16 {
	if number < lowerLimit {
		return lowerLimit
	}
	if number > upperLimit {
		return upperLimit
	}
	return number
}
