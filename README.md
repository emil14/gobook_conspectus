# GOPL

Notes and exercises from "The Go Programming Language" by Alan A. A. Donovan &amp; Brian W. Kernighan

# The origins of Go 

- Golang was developed because of (multiplicative) complexivity of Google software.
- It was influenced by Algol, Pascal, Modula, Oberon, Squeak, C and others.
- "CSP" is a thing. Go is an implementation of formal system for parallel/cuncurrent computing.

# Chapter 1

## 1.1 Hello world

- `main` package is special, it's an executable, not a library
- `main` func in `main` package is special too, it's the program's entry point
- Program consist of defenitions of 4 types - `const, type, var` and `func`. Order does not matter much
- Function body is a set of instructions
- `go fmt` is `gofmt` to all files in the directory
- `goimports` is like gofmt for but imports...


## 1.2 CLI args

"Most of programs process some input to generate some output"

- Args are accessible by `os.Args` slice
- `s[m:n]` notation includes first (0) element and excludes the last (n)
- `m` and `n` have default values `0` and `len(s)`
- `s[m:n]` where `0 <= m <= n <= len(s)` is `n - m`
-  Every package needs top-level comments, especialy `main`
- If variable doesn't initialized explicitly, it's initialized implicitly with a *zero value*
- Every arithmetic and logic operator does have a "assign" form (like `+=` for `+`)
- Go has one `for` loop that could be used in several different ways (including `range`)

Go has 4 types of variable declaration:

1. short declaration `x := 42`, can't be used outside of function
2. (implicit) default value initialization `var x int`
3. `var x = 42` - rarely used, most to declare multiple variables
4. `var x int = 42` - dont do this

Most of the time use `1` when default value matters and `2` if it doesn't

Also `strings.Join()` is usefull if you need to... join strings.

