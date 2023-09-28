package twoscomplement

import "math/big"

var bigOne = big.NewInt(1)

// ToBytes returns a variable length two's complement byte array representation of the input.
// Does not alter input.
func ToBytes(bi *big.Int) []byte {
	var resultBytes []byte
	switch bi.Sign() {
	case -1:
		// compute 2's complement
		plus1 := big.NewInt(0)
		plus1 = plus1.Add(bi, big.NewInt(1)) // add 1
		resultBytes = plus1.Bytes()
		for i, b := range resultBytes {
			resultBytes[i] = ^b // negate every bit
		}
		if len(resultBytes) == 0 || resultBytes[0]>>7 != 1 {
			// if test bit is not 1,
			// add another byte in front
			// to disambiguate from a positive number
			resultBytes = append([]byte{0xFF}, resultBytes...)
		}
	case 0:
		return []byte{}
	case 1:
		resultBytes = bi.Bytes()
		if resultBytes[0]>>7 != 0 {
			// if test bit is not 0,
			// add another byte in front
			// to disambiguate from a negative number
			resultBytes = append([]byte{0x00}, resultBytes...)
		}
	}

	return resultBytes
}

// ToBytesOfLength returns a byte array representation, 2's complement if number is negative.
// Big endian.
func ToBytesOfLength(i *big.Int, bytesLength int) []byte {
	var resultBytes []byte
	switch i.Sign() {
	case -1:
		// compute 2's complement
		plus1 := big.NewInt(0)
		plus1 = plus1.Add(i, big.NewInt(1)) // add 1
		plus1Bytes := plus1.Bytes()
		offset := len(plus1Bytes) - bytesLength
		resultBytes = make([]byte, bytesLength)
		for i := 0; i < bytesLength; i++ {
			j := offset + i
			if j < 0 {
				resultBytes[i] = 255 // pad left with 11111111
			} else {
				resultBytes[i] = ^plus1Bytes[j] // also negate every bit
			}
		}
		break
	case 0:
		// just zeroes
		resultBytes = make([]byte, bytesLength)
		break
	case 1:
		originalBytes := i.Bytes()
		resultBytes = make([]byte, bytesLength)
		offset := len(originalBytes) - bytesLength
		for i := 0; i < bytesLength; i++ {
			j := offset + i
			if j < 0 {
				resultBytes[i] = 0 // pad left with 00000000
			} else {
				resultBytes[i] = originalBytes[j]
			}
		}
		break
	}

	return resultBytes
}
