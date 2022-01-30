# GoFromTo

GoFromTo is a small library to convert different from and to different units.

Created for practice and personal user, not guaranteed to be complete or correct.

## Installation

Get the package

````shell
    go get -u github.com/noVerity/gofromto
````

Import it in your code
```go
import "github.com/noVerity/gofromto"
```

## Quick start

### For a known unit

```go
// Create the current unit you have
measure := gofromto.NewMeasure(15, gofromto.Tablespoon, "Sugar")
// Convert it to a different unit
convertedMeasure := measure.To(gofromto.Milliliter)
// Get a string representation
fmt.Print(convertedMeasure.String())
```

### From a string

```go
// Create the current unit you have
measure := gofromto.ParseMeasure("15tbsp Sugar")
// Convert it to a different unit
convertedMeasure := measure.To(gofromto.Milliliter)
// Get a string representation
fmt.Print(convertedMeasure.String())
```

## Development

### Conversions

The library takes a list of known conversions in `conversion.go` and extrapolates any other
conversions that can be achieved via additional steps. When adding a new conversion add
them to the list in both directions `Unit A -> Unit B` and `Unit B -< Unit A` then run
`go generate` before committing any changes.