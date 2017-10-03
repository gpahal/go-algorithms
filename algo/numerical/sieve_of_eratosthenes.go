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

	// Loop over all numbers starting from 2.
	for p := 2; p <= limit; p++ {
		if !sieve[p] {
			// Mark numbers p(p), p(p+1), p(p+2), ... as composite or not prime.
			for m := p * p; m <= limit; m += p {
				sieve[m] = true
			}
		}
	}

	return sieve
}
