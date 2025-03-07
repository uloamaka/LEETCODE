func closestPrimes(left int, right int) []int {
    if right < 2 || left > right {
        return []int{-1, -1}
    }
    if left < 2 {
        left = 2
    }

    if left == 2 && right >= 3 {
        return []int{2, 3} 
    }
    if left <= 3 && right < 5 {
        if left <= 2 {
            return []int{2, 3}
        }
        return []int{-1, -1} 
    }

    sqrtLimit := int(1000) 
    if right > 1000000 {
        sqrtLimit = 31623 
    }
    smallPrimes := getSmallPrimes(sqrtLimit)

    if left%2 == 0 && left > 2 {
        left++
    }

    rangeSize := right - left + 1
    sieve := make([]bool, rangeSize)
    for _, prime := range smallPrimes {
        if prime*prime > right {
            break
        }
        start := left + ((prime - left%prime) % prime)
        if start == prime {
            start += prime
        }
        for j := start; j <= right; j += prime {
            sieve[j-left] = true
        }
    }

    primes := []int{}
    if left <= 2 && right >= 2 {
        primes = append(primes, 2)
    }
    for i := 0; i < rangeSize; i++ {
        num := left + i
        if num > right {
            break
        }
        if !sieve[i] && num >= left {
            primes = append(primes, num)
        }
    }

    if len(primes) < 2 {
        return []int{-1, -1}
    }
    minDiff := right - left
    ans := []int{primes[0], primes[1]}
    for i := 1; i < len(primes); i++ {
        diff := primes[i] - primes[i-1]
        if diff < minDiff {
            minDiff = diff
            ans[0] = primes[i-1]
            ans[1] = primes[i]
        }
    }
    return ans
}

func getSmallPrimes(limit int) []int {
    if limit < 2 {
        return []int{}
    }
    sieve := make([]bool, limit+1)
    for i := 2; i*i <= limit; i++ {
        if !sieve[i] {
            for j := i * i; j <= limit; j += i {
                sieve[j] = true
            }
        }
    }
    var primes []int
    for i := 2; i <= limit; i++ {
        if !sieve[i] {
            primes = append(primes, i)
        }
    }
    return primes
}