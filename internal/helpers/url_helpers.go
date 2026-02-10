package helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
)

var (
	ErrInvalidLenURL    = errors.New("Invalid lenURL")
	ErrNumberOutOfRange = errors.New("Number out of range")
)

const (
	LOW    = 62 * 62 * 62 * 62
	HIGH   = LOW * 62
	SIZE   = HIGH - LOW
	Rounds = 10
)

func EncodeURLBase62(n uint, lenURL uint) (string, error) {
	alphabet := "jBZPywq1bM8LUX0ol7t5VgDzSIsi9mNxTeFv3rn6hKcOQRJpCd2aE4WfHkGAuY"

	if lenURL > 8 {
		return "", fmt.Errorf("%w: Must be <=8", ErrInvalidLenURL)
	}
	res := make([]byte, lenURL)

	if n >= uintPow(62, lenURL) || n <= uintPow(62, lenURL-1) {
		return "", fmt.Errorf("%w: Must be 62^%d < N < 62^%d", ErrNumberOutOfRange, lenURL-1, lenURL)
	}

	for i := 1; n > 0; i++ {
		mod := n % 62
		n = n / 62
		res[len(res)-i] = alphabet[mod]
	}

	return string(res), nil
}

func prf(key []byte, data []byte) uint64 {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	sum := mac.Sum(nil)
	return binary.BigEndian.Uint64(sum[:8])
}

func ff1Permute(n uint64, key []byte) (uint64, error) {
	m := uint64(math.Sqrt(float64(SIZE)))
	A := n / m
	B := n % m

	for i := 0; i < Rounds; i++ {
		buf := make([]byte, 24)
		binary.BigEndian.PutUint64(buf[0:8], uint64(i))
		binary.BigEndian.PutUint64(buf[8:16], B)
		binary.BigEndian.PutUint64(buf[16:24], SIZE)

		y := prf(key, buf)
		A = (A + y) % (SIZE / m)

		A, B = B, A
	}

	return A*m + B, nil
}

func PermuteRange(n uint64, key []byte) (uint64, error) {
	if n >= SIZE {
		return 0, fmt.Errorf("%w: N must be 62^%d < (N + 62^%d) < 62^%d", ErrNumberOutOfRange, 4, 4, 5)
	}

	res, err := ff1Permute(n, key)
	if err != nil {
		return 0, nil
	}

	return res + LOW, nil
}

func uintPow(base, exp uint) uint {
	res := uint(1)
	for i := uint(0); i < exp; i++ {
		res *= base
	}

	return res
}
