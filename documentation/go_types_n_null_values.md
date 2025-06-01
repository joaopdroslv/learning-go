# Go Types & Null Values

## Basic Types

- `int` => A number **without** decimal places (e.g., `-5`, `10`, `33`)
- `float64` => A number **with** decimal places (e.g., `-5.5`, `10.123`, `13.9`)
- `string` => A text value (created via double quotes or backticks: `"Hello World!"`, ``Hi everyone``)
- `bool` => `true` or `false`

But there also are some noteworthy "niche" basic types wich u'll tipically no need
that often but wich u should still know about:

- `uint` => An unsigned integer, which means a strictly non-negative integer (e.g., `0, 10, 20, 33, 77`)
- `int32` => A 32-bit signed integer, which is an integer with a specific range from `-2.147.483.648` to `2.147.483.647`
- `rune` => An alias for `int32`, represents a **Unicode code point** (i.e., a single character), and is used when dealing with Unicode characters (e.g., 'a', 'ñ', '世')
- `uint32` => A 32-bit unsigned integer, an integer that can represents values from `0` to `4.294.967.295`
- `int64` => A 64-bit signed integer, an integer that can represents a larger range of values, specifically from `-9.223.372.036.854.775.808` to `9.223.372.036.854.775.807`
- There also are more types like `int8` or `uint8` which work in the same way ( => integer with smaller number of range)

## Null Values

All Go value types come with a so-called **null value** which is the value stored in a variable if no other values is explicitly set.

For example, the following `int` variable would have a default value of `0` (because `0` is the null value of `int`, `int32` etc):

```go
var age int  // age is 0
```

Here's a list of the null values for thje different types:

- `int` => `0`
- `float64` => `0.0`
- `string` => `""` (i.e., an empty string)
- `bool` => `false`