package gofromto

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:generate go run ./conversions/generate.go

// Measure holds the information on the amount and unit of a measure and allows conversions
type Measure struct {
	Amount    float64
	Unit      Unit
	Name      string
	Imprecise bool
}

// NewMeasure creates a new Measure
func NewMeasure(amount float64, u Unit, name string) Measure {
	return Measure{amount, u, name, false}
}

// measureMatcher a regular expression to parse a string into components suitable for turning into a Measure
var measureMatcher = regexp.MustCompile(`^(\d+(?:\.\d+)?)\s*(?:(\d+)/(\d+))?\s*(\w+)\s*(.*)$`)

// ParseMeasure parse a string and return a Measure
func ParseMeasure(measureString string) (Measure, error) {
	matched := measureMatcher.FindStringSubmatch(strings.TrimSpace(measureString))
	var measure Measure

	if len(matched) < 6 {
		return measure, errors.New("measure in unknown format")
	}
	amount, err := strconv.ParseFloat(matched[1], 64)

	if err != nil {
		return measure, errors.New("unable to parse amount")
	}

	numerator, err1 := strconv.ParseFloat(matched[2], 64)
	denominator, err2 := strconv.ParseFloat(matched[3], 64)

	if err1 == nil && err2 == nil && denominator > 0 {
		amount = amount + numerator/denominator
	}

	for unit, symbol := range UnitSymbol {
		if symbol == matched[4] {
			return Measure{amount, unit, matched[5], false}, nil
		}
	}

	for unit, name := range UnitName {
		if name == matched[4] {
			return Measure{amount, unit, matched[5], false}, nil
		}
	}

	return measure, fmt.Errorf("unknown unit %v", matched[4])
}

// To converts the current measure to the given unit and returns a new Measure with the result
func (m *Measure) To(targetUnit Unit) (Measure, error) {
	if m.Unit == targetUnit {
		return *m, nil
	}

	if toMap, prs := ConversionMap[m.Unit]; prs {
		if toConversion, prs := toMap[targetUnit]; prs {
			return Measure{toConversion.Conversion(m.Amount), targetUnit, m.Name, m.Imprecise || !toConversion.Precise}, nil
		}
	}

	var measure Measure

	return measure, errors.New("no conversion found")
}

// AmountString returns a human-readable string of the amount only
func (m *Measure) AmountString() string {
	numerator, denominator := FloatToFraction(m.Amount)
	if numerator > 0 {
		if m.Amount < 1 {
			return fmt.Sprintf("%d/%d", numerator, denominator)
		}
		return fmt.Sprintf("%.0f %d/%d", m.Amount, numerator, denominator)
	}
	return fmt.Sprintf("%.0f", m.Amount)
}

// String returns a human-readable string of the measure
func (m *Measure) String() string {
	return strings.TrimSpace(m.AmountString() + m.Unit.Symbol() + " " + m.Name)
}
