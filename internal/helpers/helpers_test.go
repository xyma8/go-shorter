package helpers

import (
	"errors"
	"testing"
)

func TestUintPow(t *testing.T) {
	got := uintPow(2, 5)
	if got != 32 {
		t.Errorf("uintPow(2,5) = %d; want 32", got)
	}
}

func TestEncodeURLBase62(t *testing.T) {
	//high, err := EncodeURLBase62(uintPow(62, 4)+120, 5)
	tests := []struct {
		name   string
		n      uint
		lenURL uint
		want   string
		err    error
	}{
		{"Нижняя граница (5 символов)", uintPow(62, 4), 5, "", ErrNumberOutOfRange},
		{"Верхняя граница (5 символов)", uintPow(62, 5), 5, "", ErrNumberOutOfRange},
		{"Правильное (5 символов)", uintPow(62, 4) + 120, 5, "BjjBG", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeURLBase62(tt.n, tt.lenURL)

			if !errors.Is(err, tt.err) {
				t.Fatalf("error = %v, want %v", err, tt.err)
			}

			if err == nil && got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}

	//t.Logf("r=%v; err=%v", high, err)
}
