# Factorization

This approach uses a mix of sieve of eratosthenes and dynamic programming to remember prime numbers that have already been found. 
So the factorization only checks for prime numbers instead of all possible numbers.  

First a partial sieve of eratosthenes is build to speed up the finding of all prime numbers needed to get the factorization of the inserted number.  
Then the factors will be calculated. 
To not overshoot, the program will constantly check if the product of the factors sums up to the searched number.

## Build

command: `go build index.go`

## Usage

### STDIN
command: `echo 42 | index`

```
-------------------------
Which number should be decomposed as prime factor [pro tip: -1 = 2^(64)-1]: ---
The input: 42 has 5 bits
---
Prime factor decomposition of 42 is: [2 3 7]
---
All calculations took: 3.615683696s
Building partial sieve took: 3.615554887s
Factorization took: 4.088µs
-------------------------
```

### User Input

The program asks you for a number when it starts.  
There is a shortcut. Enter `-1` to have `2^(64)-1` as the inserted number.

`index`  

```
-------------------------
Which number should be decomposed as prime factor [pro tip: -1 = 2^(64)-1]: 42---
The input: 42 has 5 bits
---
Prime factor decomposition of 42 is: [2 3 7]
---
All calculations took: 3.615683696s
Building partial sieve took: 3.615554887s
Factorization took: 4.088µs
-------------------------
```

## Benchmark

To benchmark a number with 2^64 bit range do the following.  
command: `index`
Then enter `-1` as your number.

The program has the following output:

```
-------------------------
Which number should be decomposed as prime factor [pro tip: -1 = 2^(64)-1]: -1
---
The input: 18446744073709551614 has 64 bits
---
Prime factor decomposition of 18446744073709551614 is: [2 7 7 73 127 337 92737 649657]
---
All calculations took: 11.256513583s
Building partial sieve took: 3.765266636s
Factorization took: 7.491156074s
-------------------------
```
