package pointer_test

import (
	"errors"
	"github.com/zingerone/base_go/utils/pointer"
	"reflect"
	"testing"
)

func TestSafeInt64(t *testing.T) {
	type args struct {
		v int64
	}
	tests := []struct {
		name string
		args args
		want *int64
	}{
		{
			name: "Test case 1",
			args: args{v: 42},
			want: func() *int64 { v := int64(42); return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeInt64(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeBool(t *testing.T) {
	type args struct {
		v bool
	}
	tests := []struct {
		name string
		args args
		want *bool
	}{
		{
			name: "Test case 1",
			args: args{v: true},
			want: func() *bool { v := true; return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeBool(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeByte(t *testing.T) {
	type args struct {
		v byte
	}
	tests := []struct {
		name string
		args args
		want *byte
	}{
		{
			name: "Test case 1",
			args: args{v: byte(42)},
			want: func() *byte { v := byte(42); return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeByte(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeByte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeChannel(t *testing.T) {
	type args[T any] struct {
		v chan T
	}
	ch := make(chan int)
	type testCase[T any] struct {
		name string
		args args[T]
		want *chan T
	}
	tests := []testCase[int]{
		{
			name: "Test case 1",
			args: args[int]{v: ch},
			want: func() *chan int { v := ch; return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeChannel(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeChannel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeComplex128(t *testing.T) {
	type args struct {
		v complex128
	}
	tests := []struct {
		name string
		args args
		want *complex128
	}{
		{
			name: "Test case 1",
			args: args{v: complex128(1 + 2i)},
			want: func() *complex128 { v := complex128(1 + 2i); return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeComplex128(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeComplex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeComplex64(t *testing.T) {
	type args struct {
		v complex64
	}
	tests := []struct {
		name string
		args args
		want *complex64
	}{
		{
			name: "Test case 1",
			args: args{v: complex64(1 + 2i)},
			want: func() *complex64 { v := complex64(1 + 2i); return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeComplex64(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeComplex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeError(t *testing.T) {
	type args struct {
		v error
	}
	tests := []struct {
		name string
		args args
		want *error
	}{
		{
			name: "Test case 1",
			args: args{v: errors.New("error")},
			want: func() *error { v := errors.New("error"); return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeError(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeFloat32(t *testing.T) {
	type args struct {
		v float32
	}
	tests := []struct {
		name string
		args args
		want *float32
	}{
		{
			name: "Test case 1",
			args: args{v: float32(42.0)},
			want: func() *float32 { v := float32(42.0); return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeFloat32(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeFloat64(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want *float64
	}{
		{
			name: "Test case 1",
			args: args{v: float64(42.0)},
			want: func() *float64 { v := float64(42.0); return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeFloat64(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeInt(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		{
			name: "Test case 1",
			args: args{v: 42},
			want: func() *int { v := 42; return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeInt(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeInt32(t *testing.T) {
	type args struct {
		v int32
	}
	tests := []struct {
		name string
		args args
		want *int32
	}{
		{
			name: "Test case 1",
			args: args{v: int32(42)},
			want: func() *int32 { v := int32(42); return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeInt32(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeInterface(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name string
		args args
		want *interface{}
	}{
		{
			name: "Test case 1",
			args: args{v: "test"},
			want: func() *interface{} { v := interface{}("test"); return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeInterface(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeMap(t *testing.T) {
	type args[K comparable, V any] struct {
		v map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want *map[K]V
	}
	tests := []testCase[string, int]{
		{
			name: "Test case 1",
			args: args[string, int]{v: map[string]int{"test": 42}},
			want: func() *map[string]int { v := map[string]int{"test": 42}; return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeMap(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeOf(t *testing.T) {
	type args[T any] struct {
		v T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want *T
	}
	tests := []testCase[int]{
		{
			name: "Test case 1",
			args: args[int]{v: 42},
			want: func() *int { v := 42; return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeOf(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafePointer(t *testing.T) {
	type args[T any] struct {
		v *T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want **T
	}
	tests := []testCase[int]{
		{
			name: "Test case 1",
			args: args[int]{v: func() *int { v := 42; return &v }()},
			want: func() **int { v := func() *int { v := 42; return &v }(); return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafePointer(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafePointer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeRune(t *testing.T) {
	type args struct {
		v rune
	}
	tests := []struct {
		name string
		args args
		want *rune
	}{
		{
			name: "Test case 1",
			args: args{v: rune('a')},
			want: func() *rune { v := rune('a'); return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeRune(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeRune() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeSlice(t *testing.T) {
	type args[T any] struct {
		v []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want *[]T
	}
	tests := []testCase[int]{
		{
			name: "Test case 1",
			args: args[int]{v: []int{1, 2, 3}},
			want: func() *[]int { v := []int{1, 2, 3}; return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeSlice(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeString(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "Test case 1",
			args: args{v: "test"},
			want: func() *string { v := "test"; return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeString(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeStruct(t *testing.T) {
	type args[T any] struct {
		v T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want *T
	}
	tests := []testCase[struct{ A int }]{
		{
			name: "Test case 1",
			args: args[struct{ A int }]{v: struct{ A int }{A: 42}},
			want: func() *struct{ A int } { v := struct{ A int }{A: 42}; return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeStruct(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeUint(t *testing.T) {
	type args struct {
		v uint
	}
	tests := []struct {
		name string
		args args
		want *uint
	}{
		{
			name: "Test case 1",
			args: args{v: uint(42)},
			want: func() *uint { v := uint(42); return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeUint(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeUint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeUint16(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want *uint16
	}{
		{
			name: "Test case 1",
			args: args{v: uint16(42)},
			want: func() *uint16 { v := uint16(42); return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeUint16(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeUint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeUint32(t *testing.T) {
	type args struct {
		v uint32
	}
	tests := []struct {
		name string
		args args
		want *uint32
	}{
		{
			name: "Test case 1",
			args: args{v: uint32(42)},
			want: func() *uint32 { v := uint32(42); return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeUint32(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeUint64(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name string
		args args
		want *uint64
	}{
		{
			name: "Test case 1",
			args: args{v: uint64(42)},
			want: func() *uint64 { v := uint64(42); return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeUint64(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSafeUint8(t *testing.T) {
	type args struct {
		v uint8
	}
	tests := []struct {
		name string
		args args
		want *uint8
	}{
		{
			name: "Test case 1",
			args: args{v: uint8(42)},
			want: func() *uint8 { v := uint8(42); return &v }(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.SafeUint8(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SafeUint8() = %v, want %v", got, tt.want)
			}
		})
	}
}
