package pointer

// SetIfNotNil sets the value of a pointer to a value if the value is not nil
func SetIfNotNil[T any](pointer *T, fn func(value *T)) {
	if pointer != nil {
		fn(pointer)
	}
}

// Int returns a pointer to an int
func Int(v int) *int {
	return &v
}

// Int64 returns a pointer to an int64
func Int64(v int64) *int64 {
	return &v
}

// Uint returns a pointer to a uint
func Uint(v uint) *uint {
	return &v
}

// Uint64 returns a pointer to a uint64
func Uint64(v uint64) *uint64 {
	return &v
}

// Float32 returns a pointer to a float32
func Float32(v float32) *float32 {
	return &v
}

// Float64 returns a pointer to a float64
func Float64(v float64) *float64 {
	return &v
}

// String returns a pointer to a string
func String(v string) *string {
	return &v
}

// Bool returns a pointer to a bool
func Bool(v bool) *bool {
	return &v
}

// Int8 returns a pointer to an int8
func Int8(v int8) *int8 {
	return &v
}

// Int16 returns a pointer to an int16
func Int16(v int16) *int16 {
	return &v
}

// Int32 returns a pointer to an int32
func Int32(v int32) *int32 {
	return &v
}

// Uint8 returns a pointer to a uint8
func Uint8(v uint8) *uint8 {
	return &v
}

// Uint16 returns a pointer to a uint16
func Uint16(v uint16) *uint16 {
	return &v
}

// Uint32 returns a pointer to a uint32
func Uint32(v uint32) *uint32 {
	return &v
}

// Complex64 returns a pointer to a complex64
func Complex64(v complex64) *complex64 {
	return &v
}

// Complex128 returns a pointer to a complex128
func Complex128(v complex128) *complex128 {
	return &v
}

// Of returns a pointer to a value
func Of[T any](v T) *T {
	return &v
}
