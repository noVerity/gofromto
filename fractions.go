package gofromto

import "math"

// FRACTION_ERROR defines how precise the conversion from floating point numbers to fractions is
const FRACTION_ERROR = 0.01

// FloatToFraction takes a float and returns the fraction part as (numerator, denominator) tuple
// Special cases to consider are (0, 1) where we wouldn't want to show a fraction and (1,1) where we'd need to round
// up the initial number when rendering it
func FloatToFraction(number float64) (int, int) {
	n := math.Floor(number)
	rest := number - n
	if rest < FRACTION_ERROR {
		return 0, 1
	}
	if 1-FRACTION_ERROR < rest {
		return 1, 1
	}

	lowerN := 0.0
	lowerD := 1.0

	upperN := 1.0
	upperD := 1.0

	for {
		// Keep changing the bounds to make smaller and smaller steps until we reach a suitable fraction
		middleN := lowerN + upperN
		middleD := lowerD + upperD
		if middleD*(rest+FRACTION_ERROR) < middleN {
			upperN = middleN
			upperD = middleD
		} else if middleN < (rest-FRACTION_ERROR)*middleD {
			lowerN = middleN
			lowerD = middleD
		} else {
			return int(middleN), int(middleD)
		}
	}
	return 0, 1
}
