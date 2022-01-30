package gofromto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMeasure(t *testing.T) {
	m, err := ParseMeasure("2 tbsp Sugar")

	assert.Nil(t, err, "should successfully parse measure string")
	assert.Equal(t, 2.0, m.Amount, "amount should match")
	assert.Equal(t, Tablespoon, m.Unit, "unit should match")
	assert.Equal(t, "Sugar", m.Name, "unit should match")

	m, err = ParseMeasure("15 1/2L Honey")

	assert.Nil(t, err, "should successfully parse measure string")
	assert.Equal(t, 15.5, m.Amount, "amount should match")
	assert.Equal(t, Liter, m.Unit, "unit should match")
	assert.Equal(t, "Honey", m.Name, "unit should match")
}

func TestMeasure_To(t *testing.T) {
	m := Measure{4, Celsius, "", false}
	m2, _ := m.To(Fahrenheit)

	assert.Equal(t, Fahrenheit, m2.Unit)
	assert.InDelta(t, 39.2, m2.Amount, 0.1)

	m = Measure{32, Fahrenheit, "", false}
	m2, _ = m.To(Kelvin)

	assert.Equal(t, Kelvin, m2.Unit)
	assert.InDelta(t, 273.15, m2.Amount, 0.1)

	for from, toMap := range ConversionMap {
		for to := range toMap {
			// Every conversion should have the reverse conversion
			_, prs := ConversionMap[to][from]
			assert.Truef(t, prs, "There should be a conversion from %v to %v", to, from)
		}
	}
}

func TestMeasure_AmountString(t *testing.T) {
	m := Measure{4.333, Celsius, "", false}

	assert.Equal(t, "4 1/3", m.AmountString())
}

func TestMeasure_String(t *testing.T) {
	m := Measure{1984.5, Fahrenheit, "Books", false}

	assert.Equal(t, "1984 1/2°F Books", m.String())

	m = Measure{0.5, Gram, "Salt", false}

	assert.Equal(t, "1/2g Salt", m.String())
}
