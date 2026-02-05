package helpers

import "errors"

func EncodeBase62(n uint) (string, error) {
	alphabet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	res := []byte{0, 0, 0, 0, 0}

	if n > uintPow(62, 5) || n < uintPow(62, 4) {
		return "", errors.New("Number must be 62^4 < n < 62^5")
	}

	for i := 1; n > 0; i++ {
		mod := n % 62
		n = n / 62
		res[len(res)-i] = alphabet[mod]
	}
	return string(res), nil
}

func EncodeFeistel(n uint) (uint, error) {
	return 0, nil
}

func uintPow(base, exp uint) uint {
	res := uint(1)
	for i := uint(0); i < exp; i++ {
		res *= base
	}

	return res
}
