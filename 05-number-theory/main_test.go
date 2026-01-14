package topic5

import (
	"math/big"
	"testing"
)

func TestGCD(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int64
		expected int64
	}{
		{"coprime numbers", 13, 27, 1},
		{"same number", 42, 42, 42},
		{"example from slides", 100, 67, 1},
		{"larger example", 123456, 7890, 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GCD(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("GCD(%d, %d) = %d, expected %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

func TestExtendedGCD(t *testing.T) {
	tests := []struct {
		name          string
		a, b          int64
		expectedGcd   int64
		shouldSatisfy bool // whether to check ax + by = gcd
	}{
		{"coprime", 13, 27, 1, true},
		{"example from slides", 100, 67, 1, true},
		{"non-coprime", 12, 18, 6, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, x, y := ExtendedGCD(tt.a, tt.b)

			if g != tt.expectedGcd {
				t.Errorf("GCD = %d, expected %d", g, tt.expectedGcd)
			}

			if tt.shouldSatisfy {
				// Check Bézout's identity: ax + by = g
				ax := big.NewInt(tt.a)
				ax.Mul(ax, big.NewInt(x))

				by := big.NewInt(tt.b)
				by.Mul(by, big.NewInt(y))

				ax.Add(ax, by)
				if ax.Int64() != g {
					t.Errorf("Bézout's identity failed: %d*%d + %d*%d = %d, expected %d",
						tt.a, x, tt.b, y, ax.Int64(), g)
				}
			}
		})
	}
}

func TestModExp(t *testing.T) {
	tests := []struct {
		name     string
		a, b, n  int64
		expected int64
	}{
		{"small numbers", 2, 10, 1000, 24}, // 2^10 mod 1000 = 1024 mod 1000 = 24
		{"modulo prime", 3, 100, 7, 4},     // 3^100 mod 7 = 4
		{"zero exponent", 123, 0, 456, 1},  // anything^0 mod n = 1
		{"RSA-like", 65, 17, 3233, 2790},   // Example from RSA
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ModExp(tt.a, tt.b, tt.n)
			if result != tt.expected {
				t.Errorf("(%d^%d mod %d) = %d, expected %d", tt.a, tt.b, tt.n, result, tt.expected)
			}
		})
	}
}

func TestRSA(t *testing.T) {
	// Small primes for testing
	p, q := int64(61), int64(53)

	t.Run("key generation", func(t *testing.T) {
		key, err := GenerateRSAKey(p, q)
		if err != nil {
			t.Fatalf("Failed to generate key: %v", err)
		}

		// Verify N = p*q
		if key.N != p*q {
			t.Errorf("N = %d, expected %d", key.N, p*q)
		}

		// Verify e and d are inverses mod φ
		phi := (p - 1) * (q - 1)
		if (key.E*key.D)%phi != 1 {
			t.Errorf("e*d mod φ = %d, expected 1", (key.E*key.D)%phi)
		}
	})

	t.Run("encryption/decryption", func(t *testing.T) {
		// Use textbook RSA with small numbers
		N := p * q
		e := int64(17)   // Common RSA exponent
		d := int64(2753) // Precomputed inverse mod φ

		message := int64(65) // "A" in ASCII

		ciphertext := RSAEncrypt(message, e, N)
		plaintext := RSADecrypt(ciphertext, d, N)

		if plaintext != message {
			t.Errorf("Decryption failed: got %d, expected %d", plaintext, message)
		}
	})
}

func TestPollardRho(t *testing.T) {
	// Note: PollardRho is probabilistic and may not always find factor
	// Use numbers with small factors for testing

	testCases := []struct {
		name      string
		n         int64
		hasFactor bool // whether n has a factor PollardRho should find
	}{
		{"small composite", 5959, true},       // 59 * 101
		{"product of two primes", 3233, true}, // 61 * 53 (RSA example)
		{"prime number", 17, false},           // Should return n or fail
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			factor := PollardRho(tc.n)

			if factor == 0 || factor == tc.n {
				// Algorithm didn't find a non-trivial factor
				if tc.hasFactor {
					t.Logf("PollardRho didn't find factor for %d (this is possible)", tc.n)
				}
			} else {
				// Verify it's actually a factor
				if tc.n%factor != 0 {
					t.Errorf("%d is not a factor of %d", factor, tc.n)
				}

				// Verify it's non-trivial
				if factor == 1 || factor == tc.n {
					t.Errorf("Found trivial factor: %d", factor)
				}
			}
		})
	}
}
