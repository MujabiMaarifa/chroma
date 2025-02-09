---
title: "Go Functions for Basic Arithmetic Operations"
date: 2023-10-04
author: [Your Name]
tags: [Go, Arithmetic, Functions]
---

# Go Functions for Basic Arithmetic Operations

This document provides an overview of two simple Go functions for performing basic arithmetic operations: addition and subtraction.

## Table of Contents

- [Addition Function](#addition-function)
- [Subtraction Function](#subtraction-function)
- [Examples](#examples)

## Addition Function

The `add` function takes two integers, `x` and `y`, and returns their sum.

### Function Signature

```go
func add(x int, y int) int
```

### Parameters

- `x` (int): The first integer to be added.
- `y` (int): The second integer to be added.

### Return Value

- `int`: The sum of `x` and `y`.

### Example

```go
result := add(5, 3)
fmt.Println(result) // Output: 8
```

## Subtraction Function

The `sub` function takes two integers, `x` and `y`, and returns the difference between `x` and `y`.

### Function Signature

```go
func sub(x int, y int) int
```

### Parameters

- `x` (int): The integer from which to subtract.
- `y` (int): The integer to subtract from `x`.

### Return Value

- `int`: The difference between `x` and `y`.

### Example

```go
result := sub(10, 4)
fmt.Println(result) // Output: 6
```

## Examples

### Example 1: Adding Two Positive Integers

```go
result := add(7, 3)
fmt.Println(result) // Output: 10
```

### Example 2: Subtracting One Integer from Another

```go
result := sub(15, 8)
fmt.Println(result) // Output: 7
```

### Example 3: Adding a Positive and a Negative Integer

```go
result := add(10, -3)
fmt.Println(result) // Output: 7
```

### Example 4: Subtracting a Positive Integer from a Negative Integer

```go
result := sub(-5, 4)
fmt.Println(result) // Output: -9
```

These examples demonstrate how to use the `add` and `sub` functions with various integer inputs.