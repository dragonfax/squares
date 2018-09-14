package squares

func clamp(lowerLimit, upperLimit, number float64) float64 {
	if number < lowerLimit {
		return lowerLimit
	}
	if number > upperLimit {
		return upperLimit
	}
	return number
}
