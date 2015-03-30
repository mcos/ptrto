// Package ptrto provides functions to return pointers to Go's built in datatypes without having to instantiate
// new variables.
package ptrto

// Bool returns a pointer to a bool type
func Bool(b bool) *bool {
	return &b
}

// Byte returns a pointer to a byte type
func Byte(b byte) *byte {
	return &b
}

// Float32 returns a pointer to a float32 type
func Float32(f float32) *float32 {
	return &f
}

// Float64 returns a pointer to a float64 type
func Float64(f float64) *float64 {
	return &f
}

// Int returns a pointer to an int type
func Int(i int) *int {
	return &i
}

// Int8 returns a pointer to an int8 type
func Int8(i int8) *int8 {
	return &i
}

// Int16 returns a pointer to an int16 type
func Int16(i int16) *int16 {
	return &i
}

// Int32 returns a pointer to an int32 type
func Int32(i int32) *int32 {
	return &i
}

// Int64 returns a pointer to an int64 type
func Int64(i int64) *int64 {
	return &i
}

// Interface returns a pointer to an interface type
func Interface(v interface{}) *interface{} {
	return &v
}

// String returns a pointer to a string type
func String(s string) *string {
	return &s
}

// Struct returns a pointer to a struct type
func Struct(s struct{}) *struct{} {
	return &s
}

// Uint returns a pointer to a uint type
func Uint(u uint) *uint {
	return &u
}

// Uint8 returns a pointer to a uint8 type
func Uint8(u uint8) *uint8 {
	return &u
}

// Uint16 returns a pointer to a uint16 type
func Uint16(u uint16) *uint16 {
	return &u
}

// Uint32 returns a pointer to a uint32 type
func Uint32(u uint32) *uint32 {
	return &u
}

// Uint64 returns a pointer to a uint64 type
func Uint64(u uint64) *uint64 {
	return &u
}

// Uintptr returns a pointer to a uintptr type
func Uintptr(u uintptr) *uintptr {
	return &u
}
