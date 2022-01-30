package gofromto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnit(t *testing.T) {
	for unit := range ConversionMap {
		_, prs := UnitName[unit]
		assert.Truef(t, prs, "Unit (%v) used in the conversion should have a name", unit)
		_, prs = UnitSymbol[unit]
		assert.Truef(t, prs, "Unit (%v) used in the conversion should have a symbol", unit)
	}
}

func TestUnit_Symbol(t *testing.T) {
	for unit, symbol := range UnitSymbol {
		conversions, prs := ConversionMap[unit]
		assert.Truef(t, prs, "Unit with symbol (%v) should have at least one conversion", symbol)
		assert.True(t, len(conversions) > 0, "Unit with symbol (%v) should have at least one conversion", symbol)
		assert.NotEmpty(t, unit.Symbol())
	}
}

func TestUnit_String(t *testing.T) {
	for unit, name := range UnitName {
		conversions, prs := ConversionMap[unit]
		assert.Truef(t, prs, "Unit with name (%v) should have at least one conversion", name)
		assert.True(t, len(conversions) > 0, "Unit with name (%v) should have at least one conversion", name)
	}
}
