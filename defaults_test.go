package defaults

import (
	"bytes"
	"testing"
)

func TestOnlyAcceptStructPointers(t *testing.T) {
	err := Apply(struct{}{})
	switch err.(type) {
	case ErrNotAStructPointer:
		return
	default:
		t.Errorf("expected an ErrNotAStructPointer error, however got a %T with value '%v'", err, err)
	}
}

func TestDoNotOverrideNonZeroValues(t *testing.T) {
	test := struct {
		A string `default:"hello"`
	}{
		A: "byebye",
	}

	if test.A != "byebye" {
		t.Error("we expected the default value not to override the set value, but it did")
	}
}

func TestSetDefaults(t *testing.T) {
	test := struct {
		String    string  `default:"hello"`
		StringPtr *string `default:"hello"`

		// Technically an int32, but we have some additional parsing in place
		// for potential runes.
		Rune    rune  `default:"a"`
		RunePtr *rune `default:"a"`

		Bool    bool  `default:"true"`
		BoolPtr *bool `default:"true"`

		Int      int    `default:"1"`
		IntPtr   *int   `default:"1"`
		Int8     int8   `default:"1"`
		Int8Ptr  *int8  `default:"1"`
		Int16    int16  `default:"1"`
		Int16Ptr *int16 `default:"1"`
		Int32    int32  `default:"1"`
		Int32Ptr *int32 `default:"1"`
		Int64    int64  `default:"1"`
		Int64Ptr *int64 `default:"1"`

		Uint      uint    `default:"1"`
		UintPtr   *uint   `default:"1"`
		Uint8     uint8   `default:"1"`
		Uint8Ptr  *uint8  `default:"1"`
		Uint16    uint16  `default:"1"`
		Uint16Ptr *uint16 `default:"1"`
		Uint32    uint32  `default:"1"`
		Uint32Ptr *uint32 `default:"1"`
		Uint64    uint64  `default:"1"`
		Uint64Ptr *uint64 `default:"1"`

		Float32    float32  `default:"1.0"`
		Float32Ptr *float32 `default:"1.0"`
		Float64    float64  `default:"1.0"`
		Float64Ptr *float64 `default:"1.0"`

		ByteSlice    []byte  `default:"hello"`
		ByteSlicePtr *[]byte `default:"hello"`

		Embedded struct {
			String string `default:"hello"`
		}
		EmbeddedPtr *struct {
			String string `default:"hello"`
		}
	}{}

	if err := Apply(&test); err != nil {
		t.Errorf("could not parse struct tags: %v", err)
		return
	}

	var (
		s = "hello"
		r = 'a'
		b = true
		i = 1
	)

	err := func(t *testing.T, name string, expected, actual interface{}) {
		t.Errorf("%s: expected '%v', got '%v'", name, expected, actual)
	}

	// Strings
	if s != test.String {
		err(t, "string", s, test.String)
	}
	if s != *test.StringPtr {
		err(t, "string pointer", s, test.StringPtr)
	}

	// Runes
	if r != test.Rune {
		err(t, "rune", r, test.Rune)
	}
	if r != *test.RunePtr {
		err(t, "rune pointer", r, *test.RunePtr)
	}

	// Booleans
	if b != test.Bool {
		err(t, "bool", i, test.Bool)
	}
	if b != *test.BoolPtr {
		err(t, "bool pointer", i, *test.BoolPtr)
	}

	// Integers
	if i != test.Int {
		err(t, "int", i, test.Int)
	}
	if i != *test.IntPtr {
		err(t, "int pointer", i, *test.IntPtr)
	}
	if int8(i) != test.Int8 {
		err(t, "int8", int8(i), test.Int8)
	}
	if int8(i) != *test.Int8Ptr {
		err(t, "int8 pointer", int8(i), *test.Int8Ptr)
	}
	if int16(i) != test.Int16 {
		err(t, "int16", int16(i), test.Int16)
	}
	if int16(i) != *test.Int16Ptr {
		err(t, "int16 pointer", int16(i), *test.Int16Ptr)
	}
	if int32(i) != test.Int32 {
		err(t, "int32", int32(i), test.Int32)
	}
	if int32(i) != *test.Int32Ptr {
		err(t, "int32 pointer", int32(i), *test.Int32Ptr)
	}
	if int64(i) != test.Int64 {
		err(t, "int64", int64(i), test.Int64)
	}
	if int64(i) != *test.Int64Ptr {
		err(t, "int64 pointer", int64(i), *test.Int64Ptr)
	}

	// Unsigned integers
	if uint(i) != test.Uint {
		err(t, "uint", uint(i), test.Uint)
	}
	if uint(i) != *test.UintPtr {
		err(t, "uint pointer", uint(i), *test.UintPtr)
	}
	if uint8(i) != test.Uint8 {
		err(t, "uint8", uint8(i), test.Uint8)
	}
	if uint8(i) != *test.Uint8Ptr {
		err(t, "uint8 pointer", uint8(i), *test.Uint8Ptr)
	}
	if uint16(i) != test.Uint16 {
		err(t, "uint16", uint16(i), test.Int16)
	}
	if uint16(i) != *test.Uint16Ptr {
		err(t, "uint16 pointer", uint16(i), *test.Uint16Ptr)
	}
	if uint32(i) != test.Uint32 {
		err(t, "uint32", uint32(i), test.Uint32)
	}
	if uint32(i) != *test.Uint32Ptr {
		err(t, "uint32 pointer", uint32(i), *test.Uint32Ptr)
	}
	if uint64(i) != test.Uint64 {
		err(t, "uint64", uint64(i), test.Uint64)
	}
	if uint64(i) != *test.Uint64Ptr {
		err(t, "uint64 pointer", uint64(i), *test.Int64Ptr)
	}

	// Floats
	if float32(i) != test.Float32 {
		err(t, "float32", float32(i), test.Float32)
	}
	if float32(i) != *test.Float32Ptr {
		err(t, "float32 pointer", float32(i), *test.Float32Ptr)
	}
	if float64(i) != test.Float64 {
		err(t, "float64", float64(i), test.Float64)
	}
	if float64(i) != *test.Float64Ptr {
		err(t, "float64 pointer", float64(i), *test.Float64Ptr)
	}

	// Bytes
	if !bytes.Equal([]byte(s), test.ByteSlice) {
		err(t, "byte slice", []byte(s), test.ByteSlice)
	}
	if !bytes.Equal([]byte(s), *test.ByteSlicePtr) {
		err(t, "byte slice pointer", []byte(s), *test.ByteSlicePtr)
	}

	// Embedded structs
	if s != test.Embedded.String {
		err(t, "string in embedded struct", s, test.Embedded.String)
	}
	if s != test.EmbeddedPtr.String {
		err(t, "string in embedded struct", s, test.EmbeddedPtr.String)
	}
}
