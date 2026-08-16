# Reflect - Struct tag parser

Build a function that takes any struct and returns a slice of strings
containing only the names of fields that have a `db:"-"` tag, regardless of the
struct's type or package origin.

## Reference

- [laws-of-reflection](https://go.dev/blog/laws-of-reflection)
- [reflect](https://pkg.go.dev/reflect)
