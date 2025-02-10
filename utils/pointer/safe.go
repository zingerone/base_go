// @Author: Imam Magribi
package pointer

// SafeInt64 returns a pointer to the given int64 value.
func SafeInt64(v int64) *int64 {
	return &v
}

// SafeInt returns a pointer to the given int value.
func SafeInt(v int) *int {
	return &v
}

// SafeInt32 returns a pointer to the given int32 value.
func SafeInt32(v int32) *int32 {
	return &v
}

// SafeFloat64 returns a pointer to the given float64 value.
func SafeFloat64(v float64) *float64 {
	return &v
}

// SafeFloat32 returns a pointer to the given float32 value.
func SafeFloat32(v float32) *float32 {
	return &v
}

// SafeString returns a pointer to the given string value.
func SafeString(v string) *string {
	return &v
}

// SafeBool returns a pointer to the given bool value.
func SafeBool(v bool) *bool {
	return &v
}

// SafeUint64 returns a pointer to the given uint64 value.
func SafeUint64(v uint64) *uint64 {
	return &v
}

// SafeUint32 returns a pointer to the given uint32 value.
func SafeUint32(v uint32) *uint32 {
	return &v
}

// SafeUint16 returns a pointer to the given uint16 value.
func SafeUint16(v uint16) *uint16 {
	return &v
}

// SafeUint8 returns a pointer to the given uint8 value.
func SafeUint8(v uint8) *uint8 {
	return &v
}

// SafeUint returns a pointer to the given uint value.
func SafeUint(v uint) *uint {
	return &v
}

// SafeByte returns a pointer to the given byte value.
func SafeByte(v byte) *byte {
	return &v
}

// SafeRune returns a pointer to the given rune value.
func SafeRune(v rune) *rune {
	return &v
}

// SafeComplex64 returns a pointer to the given complex64 value.
func SafeComplex64(v complex64) *complex64 {
	return &v
}

// SafeComplex128 returns a pointer to the given complex128 value.
func SafeComplex128(v complex128) *complex128 {
	return &v
}

// SafeError returns a pointer to the given error value.
func SafeError(v error) *error {
	return &v
}

// SafeInterface returns a pointer to the given interface value.
func SafeInterface(v interface{}) *interface{} {
	return &v
}

// SafeMap returns a pointer to the given map value.
func SafeMap[K comparable, V any](v map[K]V) *map[K]V {
	return &v
}

// SafeSlice returns a pointer to the given slice value.
func SafeSlice[T any](v []T) *[]T {
	return &v
}

// SafeStruct returns a pointer to the given struct value.
func SafeStruct[T any](v T) *T {
	return &v
}

// SafePointer returns a pointer to the given pointer value.
func SafePointer[T any](v *T) **T {
	return &v
}

// SafeChannel returns a pointer to the given channel value.
func SafeChannel[T any](v chan T) *chan T {
	return &v
}

// SafeOf returns a pointer to the given value of any type.
func SafeOf[T any](v T) *T {
	return &v
}
