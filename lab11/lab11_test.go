package main

import (
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{34, 21, 1},
		{48, 18, 6},
		{17, 5, 1},
		{0, 5, 5},
		{5, 0, 5},
	}

	for _, test := range tests {
		result := gcd(test.a, test.b)
		if result != test.expected {
			t.Errorf("gcd(%d, %d) = %d, expected %d",
				test.a, test.b, result, test.expected)
		}
	}
}

func TestModInverse(t *testing.T) {
	tests := []struct {
		a, m     int
		expected int // -1 if no inverse
	}{
		{3, 11, 4}, // 3*4 = 12 ≡ 1 mod 11
		{2, 4, -1}, // No inverse, gcd(2,4)=2
		{7, 13, 2}, // 7*2 = 14 ≡ 1 mod 13
	}

	for _, test := range tests {
		result := modInverse(test.a, test.m)
		if result != test.expected {
			t.Errorf("modInverse(%d, %d) = %d, expected %d",
				test.a, test.m, result, test.expected)
		}
	}
}

func TestRSA(t *testing.T) {
	// Simple RSA example
	p, q := 11, 17
	n := p * q               // 187
	phi := (p - 1) * (q - 1) // 160
	e := 3
	d := modInverse(e, phi) // Should be 107

	message := 15
	cipher := rsaEncrypt(message, e, n)
	decrypted := rsaDecrypt(cipher, d, n)

	if decrypted != message {
		t.Errorf("RSA failed: message=%d, decrypted=%d", message, decrypted)
	}
}
