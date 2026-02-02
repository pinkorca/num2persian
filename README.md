# num2persian

[![Go Reference](https://pkg.go.dev/badge/github.com/pinkorca/num2persian.svg)](https://pkg.go.dev/github.com/pinkorca/num2persian)
[![GitHub](https://img.shields.io/github/license/pinkorca/num2persian)](https://github.com/pinkorca/num2persian)

Go package for converting numbers to Persian (Farsi) text.

## Key Features

- Convert integers (`int`, `int64`) to Persian text
- Support for very large numbers via `big.Int` (up to دسیلیون/decillion)
- Floating-point conversion with auto or manual precision
- Ordinal numbers (اول، دوم، سوم، ...)
- Currency formatting (تومان/ریال)
- Zero dependencies

## Installation

```bash
go get github.com/pinkorca/num2persian
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/pinkorca/num2persian"
)

func main() {
    fmt.Println(num2persian.Convert(1234))
    // Output: هزار و دویست و سی و چهار
}
```

## Usage

**Basic:**

```go
num2persian.Convert(0)      // صفر
num2persian.Convert(21)     // بیست و یک
num2persian.Convert(-500)   // منفی پانصد
```

**Large numbers:**

```go
num2persian.Convert(1000000)       // یک میلیون
num2persian.Convert(1000000000)    // یک میلیارد

// Very large numbers
n := new(big.Int)
n.SetString("1000000000000000000000000000000000", 10)
num2persian.ConvertBigInt(n) // یک دسیلیون
```

**Floats:**

```go
// Auto-detect precision
num2persian.ConvertString("12.5")   // دوازده ممیز پنج
num2persian.ConvertString("3.14")   // سه ممیز چهارده
num2persian.ConvertString("12.50")  // دوازده ممیز پنجاه

// Manual precision
num2persian.ConvertFloat(12.5, 1)   // دوازده ممیز پنج
```

**Ordinals:**

```go
num2persian.ConvertOrdinal(1)   // اول
num2persian.ConvertOrdinal(2)   // دوم
num2persian.ConvertOrdinal(3)   // سوم
num2persian.ConvertOrdinal(21)  // بیست و یکم
```

**Currency:**

```go
num2persian.ToToman(1500000)   // یک میلیون و پانصد هزار تومان
num2persian.ToRial(15000000)   // پانزده میلیون ریال
```

## Supported Scales

| Scale | Persian | Value |
|-------|---------|-------|
| Thousand | هزار | 10³ |
| Million | میلیون | 10⁶ |
| Billion | میلیارد | 10⁹ |
| Trillion | تریلیون | 10¹² |
| Quadrillion | کوادریلیون | 10¹⁵ |
| Quintillion | کوینتیلیون | 10¹⁸ |
| Sextillion | سکستیلیون | 10²¹ |
| Septillion | سپتیلیون | 10²⁴ |
| Octillion | اکتیلیون | 10²⁷ |
| Nonillion | نونیلیون | 10³⁰ |
| Decillion | دسیلیون | 10³³ |

## License

This project is licensed under the [MIT License](LICENSE).

## Author

[PinkOrca](https://github.com/pinkorca)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
