# Topic 5: Number-Theoretic Algorithms - Practice Instructions

## Overview
Number theory algorithms form the basis of modern cryptography. We'll implement fundamental algorithms and RSA encryption.

## Algorithms to Implement

### 1. Euclid's Algorithm (GCD)

**Basic GCD Algorithm:**
```
function gcd(a, b):
    while b ≠ 0:
        t = b
        b = a mod b
        a = t
    return a
```

**Extended Euclid (from slides):**
- Returns (d, x, y) such that ax + by = d = gcd(a,b)
- Used to compute modular inverses

**Detailed Steps:**
```
function extended_gcd(a, b):
    if b = 0:
        return (a, 1, 0)
    else:
        (d', x', y') = extended_gcd(b, a mod b)
        (d, x, y) = (d', y', x' - ⌊a/b⌋ * y')
        return (d, x, y)
```

**Modular Inverse:**
- a⁻¹ mod m exists iff gcd(a, m) = 1
- Compute using extended Euclid: ax + my = 1 ⇒ x mod m is inverse

### 2. Modular Exponentiation

**Square-and-Multiply Algorithm (from slides):**
- Compute a^b mod n efficiently
- Process bits of exponent from most significant to least

**Detailed Steps:**
```
function mod_exp(a, b, n):
    result = 1
    while b > 0:
        if b is odd:
            result = (result * a) mod n
        a = (a * a) mod n
        b = b >> 1  // right shift (divide by 2)
    return result
```

**Example:** Compute 3¹³ mod 7
```
b=13 (1101 in binary)
result=1
b odd? result=3, a=3²=9≡2 mod7, b=6
b even, a=2²=4 mod7, b=3
b odd? result=3*4=12≡5 mod7, a=4²=16≡2 mod7, b=1
b odd? result=5*2=10≡3 mod7, a=2²=4 mod7, b=0
Result: 3
```

### 3. RSA Encryption/Decryption

**Key Generation (from slides):**
1. Choose two large primes p and q (≈1024 bits each)
2. Compute n = p × q
3. Compute φ(n) = (p-1)(q-1)  // Euler's totient
4. Choose e such that:
   - 1 < e < φ(n)
   - gcd(e, φ(n)) = 1
   - Common choices: 3, 17, 65537
5. Compute d = e⁻¹ mod φ(n) using extended Euclid
6. Public key: (n, e)
7. Private key: (n, d)

**Encryption:**
- Message m must be 0 ≤ m < n
- Ciphertext c = m^e mod n

**Decryption:**
- Plaintext m = c^d mod n

**Why it works (Euler's theorem):**
- m^ed ≡ m^(kφ(n)+1) ≡ m mod n (for gcd(m,n)=1)
- Works for all m due to Chinese Remainder Theorem

**Implementation Details:**
- Use big integers (Go's `math/big` package)
- Padding required for security (PKCS#1)
- Textbook RSA is insecure for real use

### 4. Pollard's Rho Factorization

**Problem:** Factor composite integer n = p × q

**Algorithm (from slides):**
1. Choose random function f(x) = x² + c mod n
   (c ≠ 0, -2; typically c = 1)
2. Initialize: x = 2, y = 2, d = 1
3. While d = 1:
   - x = f(x) mod n
   - y = f(f(y)) mod n  // y moves twice as fast
   - d = gcd(|x-y|, n)
4. If d = n, failure (try different c or x₀)
   Else, d is a non-trivial factor

**Why it works (Birthday Paradox):**
- Sequence x₀, x₁, x₂, ... eventually repeats mod p
- Floyd's cycle detection finds collision
- |x-y| divisible by p but not necessarily by n

**Example from Slides:**
n = 455459
After several iterations: finds factor 743
n = 743 × 613

**Time Complexity:** O(√p) where p is smallest prime factor

### 5. Additional Algorithms

**Fermat's Factorization (for n = pq with close factors):**
- Try to express n as difference of squares: n = a² - b²
- Then n = (a-b)(a+b)
- Start with a = ⌈√n⌉, increment until a² - n is perfect square

**Pollard's p-1 Method:**
- Works when p-1 is smooth (all prime factors small)
- Compute a^(k!) mod n for increasing k
- gcd(a^(k!) - 1, n) may reveal p

### Number Theory Concepts

**Euler's Totient Function φ(n):**
- φ(p) = p-1 for prime p
- φ(p^k) = p^(k-1)(p-1)
- φ(ab) = φ(a)φ(b) if gcd(a,b)=1
- Formula: φ(n) = n ∏(1 - 1/p) over prime factors of n

**Chinese Remainder Theorem:**
- Solve x ≡ a mod m, x ≡ b mod n
- Solution exists if gcd(m,n) = 1
- Used in RSA decryption optimization

**Miller-Rabin Primality Test:**
- Probabilistic test
- Basis for generating RSA primes

### Implementation Tips

1. **Big Integers:** Use `math/big` for large numbers
2. **Modular Arithmetic:** Always reduce mod n after operations
3. **Efficiency:** Precompute values when possible
4. **Randomness:** Use cryptographically secure random numbers

### Security Considerations

1. **RSA Security:** Based on difficulty of factoring n = pq
2. **Key Size:** Minimum 2048 bits recommended today
3. **Padding:** Required to prevent attacks
4. **Side Channels:** Timing attacks can reveal private key

### Testing Strategy

1. **Small Examples:** Verify with hand calculations
2. **Property Testing:**
   - gcd(a,b) should divide both a and b
   - a × a⁻¹ ≡ 1 mod m
   - RSA: decrypt(encrypt(m)) = m
3. **Edge Cases:** Primes, powers of 2, large numbers
4. **Random Tests:** Test with random inputs

### Common Pitfalls

1. **Integer Overflow:** Use big integers for cryptography
2. **Modular Inverse:** Check gcd(a,m)=1 before computing
3. **RSA Message Size:** m must be < n
4. **Pollard's Rho:** May need multiple attempts with different parameters