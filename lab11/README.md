# Lab 11: Number Theory Algorithms

## Exercise 1: Modular Arithmetic Basics

### Division Property
If `a` divides `b` (written `a|b`) and `a` divides `c`, then `a` divides `(b + c)`.

Example: 3|6 and 3|9, so 3|(6+9)=15

### Quotient and Remainder
For any integers `a, b` with `b > 0`, there exist unique `q, r` such that:
```
a = b × q + r, where 0 ≤ r < b
```
- `q` = quotient
- `r` = remainder

Examples:
- 34 ÷ 3: 34 = 3×11 + 1 (q=11, r=1)
- (-8) ÷ 5: -8 = 5×(-2) + 2 (q=-2, r=2)

## Exercise 2: Euclidean Algorithm

### Goal
Find GCD (Greatest Common Divisor) of two numbers.

### Basic Euclidean Algorithm
```
function gcd(a, b):
    while b ≠ 0:
        temp = b
        b = a mod b
        a = temp
    return a
```

Example: gcd(34, 21)
```
34 ÷ 21 = 1 remainder 13
21 ÷ 13 = 1 remainder 8
13 ÷ 8 = 1 remainder 5
8 ÷ 5 = 1 remainder 3
5 ÷ 3 = 1 remainder 2
3 ÷ 2 = 1 remainder 1
2 ÷ 1 = 2 remainder 0
GCD = 1
```

### Extended Euclidean Algorithm
Finds `x, y` such that: `a×x + b×y = gcd(a,b)`

Useful for finding modular inverses:
- `a` has inverse modulo `m` if `gcd(a,m) = 1`
- Inverse = `x mod m` from extended algorithm

Example: Find inverse of 3 modulo 11
```
gcd(11, 3) = 1
11 = 3×3 + 2
3 = 2×1 + 1
2 = 1×2 + 0

Working backwards:
1 = 3 - 2×1
  = 3 - (11 - 3×3)×1
  = 3 - 11 + 3×3
  = 4×3 - 11
So: 4×3 ≡ 1 (mod 11)
Inverse of 3 mod 11 is 4
```

## Exercise 3: RSA Cryptography

### Key Generation
1. Choose two prime numbers: `p` and `q`
2. Compute: `n = p × q`
3. Compute: `φ(n) = (p-1) × (q-1)`
4. Choose `e` such that: `1 < e < φ(n)` and `gcd(e, φ(n)) = 1`
5. Compute `d = e^(-1) mod φ(n)` (using extended Euclidean)

### Encryption
For message `m`: `c = m^e mod n`

### Decryption  
For ciphertext `c`: `m = c^d mod n`

### Example
```
p = 11, q = 17
n = 187
φ(n) = 160
Choose e = 3
Find d such that 3×d ≡ 1 mod 160
d = 107 (since 3×107 = 321 ≡ 1 mod 160)

Encrypt m = 15: c = 15^3 mod 187 = 3375 mod 187 = 9
Decrypt c = 9: m = 9^107 mod 187 = 15
```

## Exercise 4: Integer Factorization

### 1. Fermat's Factorization Method
If `n = a² - b²`, then `n = (a-b)(a+b)`

Algorithm:
1. Start with `a = ceil(sqrt(n))`
2. Check if `a² - n` is a perfect square `b²`
3. If yes, factors are `(a-b)` and `(a+b)`
4. Otherwise, increment `a` and repeat

Example: Factor 5769223
```
sqrt(5769223) ≈ 2401.9
Try a = 2402: 2402² - 5769223 = 5769604 - 5769223 = 381 (not square)
Try a = 2403: 2403² - 5769223 = 5774409 - 5769223 = 5186 (not square)
...
Eventually find factors
```

### 2. Pollard's p-1 Method
Works if `p-1` has only small prime factors.

Algorithm:
1. Choose bound `B`
2. Compute `M = product of primes^exponents ≤ B`
3. Compute `a = 2^M mod n`
4. Compute `d = gcd(a-1, n)`
5. If `1 < d < n`, `d` is a factor

### 3. Pollard's Rho Algorithm
Uses random walk to find factors.

Algorithm:
1. Choose random `x`, `c`
2. Define `f(x) = (x² + c) mod n`
3. Compute sequence: `x, f(x), f(f(x)), ...`
4. Look for cycle using Floyd's algorithm
5. Compute `gcd(|x-y|, n)` at each step
6. If `1 < gcd < n`, found factor

Example for small n:
```
n = 130733
x = 2, c = 1
f(x) = x² + 1 mod n
Track until gcd(|x-y|, n) > 1
```