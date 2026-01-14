# Topic 5: Number-Theoretic Algorithms - Practice Instructions

## Overview
Implement number theory algorithms including RSA and factorization.

## Algorithms to Implement

### 1. Euclid's Algorithm
**Steps:**
- GCD(a,b) = GCD(b, a mod b)
- Extended version returns coefficients for Bézout's identity

### 2. Modular Exponentiation
**Steps:**
- Square-and-multiply method
- Compute a^b mod N efficiently

### 3. RSA Encryption/Decryption
**Steps:**
1. Generate RSA keys (p, q primes, N = pq, choose e, compute d)
2. Encrypt: c = m^e mod N
3. Decrypt: m = c^d mod N

### 4. Pollard's Rho Factorization
**Steps:**
1. Use polynomial f(x) = x² + 1 mod N
2. Detect cycle using Floyd's algorithm
3. Compute GCD(|x - y|, N) to find factor