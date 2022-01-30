package gofromto

// Conversion holds the rules of how to convert from one unit to another
type Conversion struct {
	From       Unit
	To         Unit
	Conversion func(amount float64) float64
	Precise    bool
}

// ConversionList a list of all defined conversions. There is a full mapping in ConversionMap that is generated from
// this list
var ConversionList = []Conversion{
	// Volume
	{Cup, Tablespoon, func(amount float64) float64 { return amount * 15.7725491 }, true},
	{Cup, Teaspoon, func(amount float64) float64 { return amount * 47.3176473 }, true},
	{Cup, Milliliter, func(amount float64) float64 { return amount * 236.5882365 }, true},
	{Milliliter, Cup, func(amount float64) float64 { return amount * 0.0042267528 }, true},
	{Tablespoon, Milliliter, func(amount float64) float64 { return amount * 14.7867647813 }, true},
	{Milliliter, Liter, func(amount float64) float64 { return amount * 0.001 }, true},
	{Liter, Milliliter, func(amount float64) float64 { return amount * 1000 }, true},
	{Liter, Gallon, func(amount float64) float64 { return amount * 0.2641720524 }, true},
	{Gallon, Liter, func(amount float64) float64 { return amount * 3.785411784 }, true},
	{Liter, Quart, func(amount float64) float64 { return amount * 1.0566882094 }, true},
	{Quart, Liter, func(amount float64) float64 { return amount * 0.946352946 }, true},
	{Liter, Pint, func(amount float64) float64 { return amount * 2.1133764189 }, true},
	{Pint, Liter, func(amount float64) float64 { return amount * 0.473176473 }, true},
	{Teaspoon, Milliliter, func(amount float64) float64 { return amount * 4.9289215938 }, true},
	{Cup, FluidOunce, func(amount float64) float64 { return amount * 8 }, true},
	{FluidOunce, Cup, func(amount float64) float64 { return amount / 8 }, true},
	// Weight
	{Ounce, Gram, func(amount float64) float64 { return amount * 28.349523125 }, true},
	{Pound, Ounce, func(amount float64) float64 { return amount * 16 }, true},
	{Pound, Gram, func(amount float64) float64 { return amount * 453.59237 }, true},
	{Gram, Pound, func(amount float64) float64 { return amount * 0.0022046226 }, true},
	{Gram, Kilogram, func(amount float64) float64 { return amount * 0.001 }, true},
	{Kilogram, Gram, func(amount float64) float64 { return amount * 1000 }, true},
	{Kilogram, MetricTon, func(amount float64) float64 { return amount * 0.001 }, true},
	{MetricTon, Kilogram, func(amount float64) float64 { return amount * 1000 }, true},
	// Temperature
	{Fahrenheit, Celsius, func(amount float64) float64 { return (amount - 32) * 5 / 9 }, true},
	{Celsius, Fahrenheit, func(amount float64) float64 { return (amount * 9 / 5) + 32 }, true},
	{Celsius, Kelvin, func(amount float64) float64 { return amount + 273.15 }, true},
	{Kelvin, Celsius, func(amount float64) float64 { return amount - 273.15 }, true},
	{Celsius, Rankine, func(amount float64) float64 { return (amount + 273.15) * 9 / 5 }, true},
	{Rankine, Celsius, func(amount float64) float64 { return (amount - 491.67) * 5 / 9 }, true},
	// Distance
	{Centimeter, Inch, func(amount float64) float64 { return amount * 0.3937007874 }, true},
	{Inch, Centimeter, func(amount float64) float64 { return amount * 2.54 }, true},
	{Millimeter, Centimeter, func(amount float64) float64 { return amount * 0.1 }, true},
	{Centimeter, Millimeter, func(amount float64) float64 { return amount * 10 }, true},
	{Centimeter, Meter, func(amount float64) float64 { return amount * 0.01 }, true},
	{Meter, Centimeter, func(amount float64) float64 { return amount * 100 }, true},
	{Meter, Kilometer, func(amount float64) float64 { return amount * 0.001 }, true},
	{Kilometer, Meter, func(amount float64) float64 { return amount * 1000 }, true},
	{Meter, Mile, func(amount float64) float64 { return amount * 0.0006213712 }, true},
	{Mile, Meter, func(amount float64) float64 { return amount * 1609.344 }, true},
	{Inch, Feet, func(amount float64) float64 { return amount * 0.0833333333 }, true},
	{Feet, Inch, func(amount float64) float64 { return amount * 12 }, true},
	{Inch, Yard, func(amount float64) float64 { return amount * 0.0277777778 }, true},
	{Yard, Inch, func(amount float64) float64 { return amount * 36 }, true},
}
