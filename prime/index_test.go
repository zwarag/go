package main

import (
	"math"
	"testing"
)

func benchmarkPrime(n uint64, b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeFactorize(n)
	}
}

func BenchmarkPrime4(b *testing.B) {
	benchmarkPrime(uint64(math.Pow(2, 4))-1, b)
}

func BenchmarkPrime8(b *testing.B) {
	benchmarkPrime(uint64(math.Pow(2, 8))-1, b)
}

func BenchmarkPrime16(b *testing.B) {
	benchmarkPrime(uint64(math.Pow(2, 16))-1, b)
}

func BenchmarkPrime32(b *testing.B) {
	benchmarkPrime(uint64(math.Pow(2, 32))-1, b)
}

func BenchmarkPrime64(b *testing.B) {
	benchmarkPrime(uint64(math.MaxUint64) - 1, b)
}

func BenchmarkPrimeLargePrime(b *testing.B) {
	benchmarkPrime(uint64(18446744073709551557), b)
}
