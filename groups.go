package gofromto

type UnitGroup int64

const (
	Metric UnitGroup = iota
	Imperial
)

var MetricGroupings = [][]Unit{
	{Milliliter, Liter},
	{Gram, Kilogram, MetricTon},
	{Millimeter, Centimeter, Meter, Kilometer},
}

var ImperialGroupings = [][]Unit{
	{Inch, Feet, Yard, Mile},
	{FluidOunce, Pint, Quart, Gallon},
	{Ounce, Pound},
}

var Groupings = map[UnitGroup][][]Unit{
	Metric:   MetricGroupings,
	Imperial: ImperialGroupings,
}
