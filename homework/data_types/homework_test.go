package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func ToLittleEndian(n uint32) uint32 {
	b0 := (n & 0x000000FF) << 24 // Младший байт  переносим на место старшего
	b1 := (n & 0x0000FF00) << 8  // Второй байт  сдвигаем на позицию третьего
	b2 := (n & 0x00FF0000) >> 8  // Третий байт  сдвигаем на позицию второго
	b3 := (n & 0xFF000000) >> 24 // Старший байт  переносим на место младшего

	return b0 | b1 | b2 | b3
}

func TestСonversion(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
