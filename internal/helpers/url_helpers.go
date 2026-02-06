package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
)

var (
	ErrInvalidLenURL    = errors.New("Invalid lenURL")
	ErrNumberOutOfRange = errors.New("Number out of range")
)

func EncodeURLBase62(n uint, lenURL uint) (string, error) {
	alphabet := "jBZPywq1bM8LUX0ol7t5VgDzSIsi9mNxTeFv3rn6hKcOQRJpCd2aE4WfHkGAuY"

	if lenURL > 8 {
		return "", fmt.Errorf("%w: Must be <=8", ErrInvalidLenURL)
	}
	res := make([]byte, lenURL)

	if n >= uintPow(62, lenURL) || n <= uintPow(62, lenURL-1) {
		return "", fmt.Errorf("%w: Must be 62^%d < n < 62^%d", ErrNumberOutOfRange, lenURL-1, lenURL)
	}

	for i := 1; n > 0; i++ {
		mod := n % 62
		n = n / 62
		res[len(res)-i] = alphabet[mod]
	}

	return string(res), nil
}

func Feistel(n uint) (uint32, error) {
	SIZE := uintPow(62, 5) - uintPow(62, 4)
	n = n % SIZE
	L := uint16(n & 0xFFFF)
	R := uint16(n >> 16)
	fmt.Println(L, R)
	key := []byte("super-secret-key-32-bytes_do-not-change")
	rounds := 3

	for i := 0; i < rounds; i++ {
		h := hmac.New(sha256.New, key)
		h.Write([]byte{byte(R >> 8), byte(R), byte(i)})
		sum := h.Sum(nil)
		F := binary.BigEndian.Uint16(sum[:2])

		L, R = R, L^F
	}
	res := (uint32(R)<<16 | uint32(L)) + uint32(uintPow(62, 4))
	return res, nil
}

func EncodeBiject(n uint) (uint, error) {
	LOW := uintPow(62, 4)
	SIZE := uintPow(62, 5) - uintPow(62, 4)
	var a uint = 45555
	var b uint = 909
	res := ((a*n + b) % SIZE) + LOW

	return res, nil
}

func uintPow(base, exp uint) uint {
	res := uint(1)
	for i := uint(0); i < exp; i++ {
		res *= base
	}

	return res
}
