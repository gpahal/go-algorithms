package numerical

// SieveOfEratosthenes finds all prime numbers up to any given limit. It returns a bool slice of length `limit` and its
// nth index is true if n is NOT a prime and false if it is. 0 and 1 are not primes.
//
// NOTE: limit should be > 0 and <= 2 ^ 24 (around 16.7 million), otherwise nil is returned.
func SieveOfEratosthenes(limit int) []bool {
	// NOTE: 1 << n = 2 ^ n (as long as 2 ^ n is within the limits of what an int value can hold)
	if limit <= 0 || limit > (1<<24) {
		return nil
	}

	// sieve is the slice that stores primality values. By default, all numbers are prime ie. false.
	sieve := make([]bool, limit+1)

	// Marks 0 and 1 as NOT prime.
	sieve[0] = true
	sieve[1] = true

	// Mark multiples of 2 as composite or not prime.
	for m := 4; m <= limit; m += 2 {
		sieve[m] = true
	}

	// Loop over all odd numbers starting from 3. 2 is being treated as a special case as it is the only prime that is
	// even. This allows us to skip all even numbers in the loop.
	for p := 3; p <= limit; p += 2 {
		if !sieve[p] {
			// Mark numbers p(p), p(p+2), p(p+4), ... as composite or not prime. NOTE: p+1, p+3 are even and hence
			// p(p+1), p(p+3), ... have already been marked as they are multiples of 2.
			for m := p * p; m <= limit; m += 2 * p {
				sieve[m] = true
			}
		}
	}

	return sieve
}
