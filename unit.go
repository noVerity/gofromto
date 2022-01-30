package gofromto

// Unit represents the units
type Unit int64

// String returns human-readable name of the unit
func (u Unit) String() string {
	return UnitName[u]
}

// Symbol returns human-readable symbol of the unit
func (u Unit) Symbol() string {
	return UnitSymbol[u]
}

const (
	Tablespoon Unit = iota
	Teaspoon
	Hint
	Drop
	Smidgen
	Pinch
	Dash
	Cup
	Milliliter
	Liter
	Gallon
	Quart
	Pint
	FluidOunce
	Ounce
	Gram
	Kilogram
	MetricTon
	Pound
	Fahrenheit
	Celsius
	Kelvin
	Rankine
	Millimeter
	Centimeter
	Meter
	Kilometer
	Mile
	Inch
	Feet
	Yard
)

// UnitName maps units to their names
var UnitName = map[Unit]string{
	Tablespoon: "Tablespoon",
	Teaspoon:   "Teaspoon",
	Hint:       "Hint",
	Drop:       "Drop",
	Smidgen:    "Smidgen",
	Pinch:      "Pinch",
	Dash:       "Dash",
	Cup:        "Cup",
	Milliliter: "Milliliter",
	Liter:      "Liter",
	Gallon:     "Gallon",
	Quart:      "Quart",
	Pint:       "Pint",
	FluidOunce: "Fluid Ounce",
	Ounce:      "Ounce",
	Gram:       "Gram",
	Kilogram:   "Kilogram",
	MetricTon:  "Ton (metric)",
	Pound:      "Pound",
	Fahrenheit: "Fahrenheit",
	Celsius:    "Celsius",
	Kelvin:     "Kelvin",
	Rankine:    "Rankine",
	Millimeter: "Millimeter",
	Centimeter: "Centimeter",
	Meter:      "Meter",
	Kilometer:  "Kilometer",
	Mile:       "Mile",
	Inch:       "Inch",
	Feet:       "Feet",
	Yard:       "Yard",
}

// UnitSymbol maps units to their symbols
var UnitSymbol = map[Unit]string{
	Tablespoon: "tbsp",
	Teaspoon:   "tsp",
	Hint:       "hint",
	Drop:       "dr.",
	Smidgen:    "smdg.",
	Pinch:      "pn.",
	Dash:       "ds.",
	Cup:        "cups",
	Milliliter: "mL",
	Liter:      "L",
	Gallon:     "gal",
	Quart:      "qt",
	Pint:       "pt",
	FluidOunce: "fl oz",
	Ounce:      "oz",
	Gram:       "g",
	Kilogram:   "kg",
	MetricTon:  "t",
	Pound:      "lbs",
	Fahrenheit: "°F",
	Celsius:    "°C",
	Kelvin:     "°K",
	Rankine:    "°R",
	Millimeter: "mm",
	Centimeter: "cm",
	Meter:      "m",
	Kilometer:  "km",
	Mile:       "mile",
	Inch:       "in",
	Feet:       "f",
	Yard:       "yrd",
}
