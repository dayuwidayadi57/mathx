# mathx
Extended Math library for Go with Generics support.

## Features
- **Generics Support**: Supports `int`, `float64`, `int64`, and more without manual type casting.
- **Smart Rounding**: Precision rounding to a specific number of decimal places.
- **Array Helpers**: Built-in `Sum` and `Average` functions for numeric slices.
- **Clean Code**: Zero dependencies, lightweight, and high performance.

## Installation
```bash
go get [github.com/dayuwidayadi57/mathx](https://github.com/dayuwidayadi57/mathx)


## Usage
import "[github.com/dayuwidayadi57/mathx](https://github.com/dayuwidayadi57/mathx)"

// Automatic type inference (no float64 casting needed)
result := mathx.Max(10, 20)          // returns 20 (int)
avg := mathx.Average([]int{1, 2, 3}) // returns 2.0 (float64)
rounded := mathx.RoundTo(3.1415, 2)  // returns 3.14 (float64)
