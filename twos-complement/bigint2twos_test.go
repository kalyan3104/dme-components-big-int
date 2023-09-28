package twoscomplement

import (
	"bytes"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToBytesOf(t *testing.T) {
	assertToBytesOk(t, "0", []byte{})
	assertToBytesOk(t, "1", []byte{0x01})
	assertToBytesOk(t, "-1", []byte{0xFF})
	assertToBytesOk(t, "-2", []byte{0xFE})
	assertToBytesOk(t, "255", []byte{0x00, 0xFF})
	assertToBytesOk(t, "256", []byte{0x01, 0x00})
	assertToBytesOk(t, "-255", []byte{0xFF, 0x01})
	assertToBytesOk(t, "-256", []byte{0xFF, 0x00})
	assertToBytesOk(t, "-257", []byte{0xFE, 0xFF})
}

func assertToBytesOk(t *testing.T, input string, expected []byte) {
	inputBi := big.NewInt(0)
	_ = inputBi.UnmarshalText([]byte(input))

	result := ToBytes(inputBi)
	assert.True(t, bytes.Equal(result, expected), "ToBytes returned wrong result. Want: %v. Have: %v", expected, result)
}

func TestToBytesOfLength(t *testing.T) {
	assertToBytesOfLengthOk(t, "0", 0, []byte{})
	assertToBytesOfLengthOk(t, "0", 1, []byte{0x00})
	assertToBytesOfLengthOk(t, "1", 1, []byte{0x01})
	assertToBytesOfLengthOk(t, "-1", 1, []byte{0xFF})
	assertToBytesOfLengthOk(t, "0", 3, []byte{0x00, 0x00, 0x00})
	assertToBytesOfLengthOk(t, "1", 3, []byte{0x00, 0x00, 0x01})
	assertToBytesOfLengthOk(t, "-1", 3, []byte{0xFF, 0xFF, 0xFF})
}

func assertToBytesOfLengthOk(t *testing.T, input string, length int, expected []byte) {
	inputBi := big.NewInt(0)
	_ = inputBi.UnmarshalText([]byte(input))

	result := ToBytesOfLength(inputBi, length)
	assert.True(t, bytes.Equal(result, expected), "ToBytesOfLength returned wrong result")
}
