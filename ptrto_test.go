package ptrto

import (
	"reflect"
	"testing"
)

const (
	testBool    bool    = true
	testByte    byte    = 0x01
	testFloat32 float32 = 1.23
	testFloat64 float64 = 1.23
	testInt     int     = 1
	testInt8    int8    = 2
	testInt16   int16   = 3
	testInt32   int32   = 4
	testInt64   int64   = 5
	testString  string  = "this is a test"
	testUint    uint    = 1
	testUint8   uint8   = 2
	testUint16  uint16  = 3
	testUint32  uint32  = 4
	testUint64  uint64  = 5
	testUintptr uintptr = 1
)

// Some types cannot be `const`
var (
	testStruct struct{}
)

func TestBool(t *testing.T) {
	ptr := Bool(testBool)

	if reflect.Ptr != reflect.TypeOf(ptr).Kind() {
		t.Errorf("Cannot return pointer to bool type")
	}
}

func TestByte(t *testing.T) {
	ptr := Byte(testByte)

	if reflect.Ptr != reflect.TypeOf(ptr).Kind() {
		t.Errorf("Cannot return pointer to byte type")
	}
}

func TestFloat32(t *testing.T) {
	ptr := Float32(testFloat32)

	if reflect.Ptr != reflect.TypeOf(ptr).Kind() {
		t.Errorf("Cannot return pointer to float32 type")
	}
}

func TestFloat64(t *testing.T) {
	ptr := Float64(testFloat64)

	if reflect.Ptr != reflect.TypeOf(ptr).Kind() {
		t.Errorf("Cannot return pointer to float64 type")
	}
}

func TestInt(t *testing.T) {
	ptr := Int(testInt)

	if reflect.Ptr != reflect.TypeOf(ptr).Kind() {
		t.Errorf("Cannot return pointer to int type")
	}
}

func TestInt8(t *testing.T) {
	ptr := Int8(testInt8)

	if reflect.Ptr != reflect.TypeOf(ptr).Kind() {
		t.Errorf("Cannot return pointer to int8 type")
	}
}

func TestInt16(t *testing.T) {
	ptr := Int16(testInt16)

	if reflect.Ptr != reflect.TypeOf(ptr).Kind() {
		t.Errorf("Cannot return pointer to int16 type")
	}
}

func TestInt32(t *testing.T) {
	ptr := Int32(testInt32)

	if reflect.Ptr != reflect.TypeOf(ptr).Kind() {
		t.Errorf("Cannot return pointer to int32 type")
	}
}

func TestInt64(t *testing.T) {
	ptr := Int64(testInt64)

	if reflect.Ptr != reflect.TypeOf(ptr).Kind() {
		t.Errorf("Cannot return pointer to int64 type")
	}
}

func TestString(t *testing.T) {
	ptr := String(testString)

	if reflect.Ptr != reflect.TypeOf(ptr).Kind() {
		t.Errorf("Cannot return pointer to string type")
	}
}

func TestStruct(t *testing.T) {
	ptr := Struct(testStruct)

	if reflect.Ptr != reflect.TypeOf(ptr).Kind() {
		t.Errorf("Cannot return pointer to struct type")
	}
}

func TestUint(t *testing.T) {
	ptr := Uint(testUint)

	if reflect.Ptr != reflect.TypeOf(ptr).Kind() {
		t.Errorf("Cannot return pointer to uint type")
	}
}

func TestUint8(t *testing.T) {
	ptr := Uint8(testUint8)

	if reflect.Ptr != reflect.TypeOf(ptr).Kind() {
		t.Errorf("Cannot return pointer to uint8 type")
	}
}

func TestUint16(t *testing.T) {
	ptr := Uint16(testUint16)

	if reflect.Ptr != reflect.TypeOf(ptr).Kind() {
		t.Errorf("Cannot return pointer to uint16 type")
	}
}

func TestUint32(t *testing.T) {
	ptr := Uint32(testUint32)

	if reflect.Ptr != reflect.TypeOf(ptr).Kind() {
		t.Errorf("Cannot return pointer to uint32 type")
	}
}

func TestUint64(t *testing.T) {
	ptr := Uint64(testUint64)

	if reflect.Ptr != reflect.TypeOf(ptr).Kind() {
		t.Errorf("Cannot return pointer to uint64 type")
	}
}

func TestUintptr(t *testing.T) {
	ptr := Uintptr(testUintptr)

	if reflect.Ptr != reflect.TypeOf(ptr).Kind() {
		t.Errorf("Cannot return pointer to uintptr type")
	}
}
