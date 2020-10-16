package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type sieve [math.MaxInt32]bool

var factorizationStart, primeCalculationStart, allStart time.Time
var factorizationTime, primeCalculationTime time.Duration

/*
 * Reads from STDIN.
 * example: `go build && echo 42 | index`
 */
func main() {

	fmt.Println("-------------------------")
	input, err := readIntFromUser("Which number should be decomposed as prime factor [pro tip: -1 = 2^(64)-1]: ")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Closing Program.")
		fmt.Println("-------------------------")
		return
	}

	allStart = time.Now()

	var parsedInput uint64

	if input == "-1" {
		parsedInput = uint64(math.MaxUint64) - 1
	} else {
		var err error
		parsedInput, err = convertToUInt64(input)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Closing Program.")
			fmt.Println("-------------------------")
			return
		}
	}

	fmt.Println("---")
	fmt.Println("The input:", parsedInput, "has", countBits(parsedInput), "bits")
	fmt.Println("---")
	PrimeFactorize(parsedInput)
	fmt.Println("---")

	fmt.Println("All calculations took:", time.Since(allStart))
	fmt.Println("Building partial sieve took:", primeCalculationTime)
	fmt.Println("Factorization took:", factorizationTime)
	fmt.Println("-------------------------")
}

func PrimeFactorize(n uint64) {
	primeCalculationStart = time.Now()
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

	primeCalculationTime = time.Since(primeCalculationStart)

	/*
	 * Uncomment to check to get visual output of all the prime numbers.
	 */
	// if countBits(n) <= 10 {
	// 	for i, val := range sieve {
	// 		if uint64(i) > n {
	// 			break
	// 		}
	// 		if val {
	// 			fmt.Print(i, " ")
	// 		}
	// 	}
	// 	fmt.Println("")
	// }

	f := factorize(n, sieve)
	if len(f) == 0 {
		fmt.Println(n, "is a prime number")
	} else {
		fmt.Println("Prime factor decomposition of", n, "is:", f)
	}

}

func factorize(n uint64, sp *sieve) []uint64 {
	factorizationStart = time.Now()
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
	factorizationTime = time.Since(factorizationStart)
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
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return i, err
}

