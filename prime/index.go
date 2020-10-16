package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type sieve [math.MaxInt32]bool

func main() {

	// input, err := readIntFromUser("Which number should be decomposed as prime factor: ")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	pow := 64
	//number := math.Pow(float64(2), float64(pow))
	number := uint64(math.MaxUint64)

	n := uint64(number)
	fmt.Println(n, "has", countBits(n), "bits")

	var sieve = new(sieve)
	for i := range sieve {
		sieve[i] = true
	}
	sieve[0] = false
	sieve[1] = false

	sieve.removeMultiple(2)
	sieve.removeMultiple(3)
	sieve.removeMultiple(5)
	sieve.removeMultiple(7)
	sieve.removeMultiple(11)
	sieve.removeMultiple(13)

	top := uint32(math.Floor(math.Sqrt(float64(n))))
	// fmt.Println("Will check up to:", top)

	for i := uint32(3); i <= top; i += 2 {
		IsPrimeSqrtSieve(i, sieve)
	}

	if pow <= 10 {
		for i, val := range sieve {
			if uint64(i) > n {
				break
			}
			if val {
				fmt.Print(i, " ")
			}
		}
		fmt.Println("")
	}

	f := factorize(n, sieve)
	fmt.Println(f)

	// k := uint64(281474976710655)
	// fmt.Println(k, "has", countBits(k), "bits")
	// fmt.Println(IsPrimeSqrt(281474976710655))

}

func factorize(n uint64, sp *sieve) []uint64 {
	var retVal []uint64

	half := uint64(math.Floor(float64(n) / 2))
	for ; ; {
		if n%2 == 0 {
			retVal = append(retVal, 2)
			n = n / 2
		} else {
			break
		}
	}
	done := checkIfDone(n, retVal)
	if done {
		return retVal
	}
	for prime, isPrime := range sp {
		p := uint64(prime)
		if p > half {
			break
		}
		if !isPrime {
			continue
		}
		for ; ; {
			if n%p == 0 {
				retVal = append(retVal, p)
				n = n / p
			} else {
				break
			}
			done := checkIfDone(n, retVal)
			if done {
				return retVal
			}
		}
	}
	return retVal
}

func checkIfDone(n uint64, retVal []uint64) bool {
	prod := uint64(1)
	for _, l := range retVal {
		prod *= l
	}
	if prod == n {
		return true
	}
	return false
}

func countBits(v uint64) uint64 {
	return uint64(math.Floor(math.Log2(float64(v))))
}

func isPrime(v uint64) bool {
	top := uint32(math.Floor(math.Sqrt(float64(v))))
	for i := uint32(2); i <= top; i++ {
		if v%uint64(i) == 0 {
			return false
		}
	}
	return true
}

func IsPrimeSqrtSieve(value uint32, sp *sieve) bool {
	top := uint32(math.Floor(math.Sqrt(float64(value))))
	if value%2 == 0 {
		return false
	}
	for i := uint32(3); i <= top; i += 2 {
		if !sp[i] {
			continue
		}
		if value%i == 0 {
			if !sp[i] {
				sp.removeMultiple(i)
			}
			return false
		}
	}
	return value > 1
}

func (sp *sieve) removeMultiple(n uint32) {
	l := len(sp)
	for i := 2 * n; i < uint32(l); i += n {
		sp[i] = false
	}
}

func readIntFromUser(s string) (string, error) {
	fmt.Print(s)
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", errors.New("an error occurs while reading input")
	}

	input = strings.TrimSuffix(input, "\n")
	return input, nil
}

func convertToUInt64(s string) (uint64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return uint64(i), nil
}
