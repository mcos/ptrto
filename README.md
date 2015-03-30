ptrto
=====

`ptrto` is a tiny library that exposes functions that return pointers to the `Go` builtin types (http://golang.org/pkg/builtin).

The following types are supported:
* bool
* byte
* float32
* float64
* int
* int8
* int16
* int32
* int64
* string
* uint
* uint8
* uint16
* uint32
* uint64
* uintptr

The motivation for this was tests for functions and structs with `interface{}` arguments and types. In reality it's not likely that you'll need to use pointers to these types.
