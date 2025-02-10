// @Author: Imam Magribi
package pointer_test

import (
	"github.com/zingerone/base_go/utils/pointer"
	"reflect"
	"testing"
)

func TestBool(t *testing.T) {
	type args struct {
		v bool
	}
	tests := []struct {
		name string
		args args
		want *bool
	}{
		{
			name: "TestBoolTrue",
			args: args{v: true},
			want: pointer.Bool(true),
		},
		{
			name: "TestBoolFalse",
			args: args{v: false},
			want: pointer.Bool(false),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.Bool(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128(t *testing.T) {
	type args struct {
		v complex128
	}
	tests := []struct {
		name string
		args args
		want *complex128
	}{
		{
			name: "TestComplex128",
			args: args{v: complex(1, 2)},
			want: pointer.Complex128(complex(1, 2)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.Complex128(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64(t *testing.T) {
	type args struct {
		v complex64
	}
	tests := []struct {
		name string
		args args
		want *complex64
	}{
		{
			name: "TestComplex64",
			args: args{v: complex(1, 2)},
			want: pointer.Complex64(complex(1, 2)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.Complex64(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	type args struct {
		v float32
	}
	tests := []struct {
		name string
		args args
		want *float32
	}{
		{
			name: "TestFloat32",
			args: args{v: 1.23},
			want: pointer.Float32(1.23),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.Float32(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want *float64
	}{
		{
			name: "TestFloat64",
			args: args{v: 1.23},
			want: pointer.Float64(1.23),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.Float64(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want *int
	}{
		{
			name: "TestInt",
			args: args{v: 123},
			want: pointer.Int(123),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.Int(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	type args struct {
		v int16
	}
	tests := []struct {
		name string
		args args
		want *int16
	}{
		{
			name: "TestInt16",
			args: args{v: 123},
			want: pointer.Int16(123),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.Int16(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	type args struct {
		v int32
	}
	tests := []struct {
		name string
		args args
		want *int32
	}{
		{
			name: "TestInt32",
			args: args{v: 123},
			want: pointer.Int32(123),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.Int32(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	type args struct {
		v int64
	}
	tests := []struct {
		name string
		args args
		want *int64
	}{
		{
			name: "TestInt64",
			args: args{v: 123},
			want: pointer.Int64(123),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.Int64(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	type args struct {
		v int8
	}
	tests := []struct {
		name string
		args args
		want *int8
	}{
		{
			name: "TestInt8",
			args: args{v: 123},
			want: pointer.Int8(123),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.Int8(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOf(t *testing.T) {
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
			name: "TestOfInt",
			args: args[int]{v: 123},
			want: pointer.Int(123),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.Of(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Of() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSetIfNotNil(t *testing.T) {
	type args[T any] struct {
		pointer *T
		fn      func(value *T)
	}
	type testCase[T any] struct {
		name string
		args args[T]
	}
	tests := []testCase[int]{
		{
			name: "TestSetIfNotNil",
			args: args[int]{pointer.Int(123), func(value *int) {
				*value = 456

			},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pointer.SetIfNotNil(tt.args.pointer, tt.args.fn)
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want *string
	}{
		{
			name: "TestString",
			args: args{v: "test"},
			want: pointer.String("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.String(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint(t *testing.T) {
	type args struct {
		v uint
	}
	tests := []struct {
		name string
		args args
		want *uint
	}{
		{
			name: "TestUint",
			args: args{v: 123},
			want: pointer.Uint(123),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.Uint(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		name string
		args args
		want *uint16
	}{
		{
			name: "TestUint16",
			args: args{v: 123},
			want: pointer.Uint16(123),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.Uint16(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	type args struct {
		v uint32
	}
	tests := []struct {
		name string
		args args
		want *uint32
	}{
		{
			name: "TestUint32",
			args: args{v: 123},
			want: pointer.Uint32(123),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.Uint32(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name string
		args args
		want *uint64
	}{
		{
			name: "TestUint64",
			args: args{v: 123},
			want: pointer.Uint64(123),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.Uint64(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8(t *testing.T) {
	type args struct {
		v uint8
	}
	tests := []struct {
		name string
		args args
		want *uint8
	}{
		{
			name: "TestUint8",
			args: args{v: 123},
			want: pointer.Uint8(123),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pointer.Uint8(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uint8() = %v, want %v", got, tt.want)
			}
		})
	}
}
