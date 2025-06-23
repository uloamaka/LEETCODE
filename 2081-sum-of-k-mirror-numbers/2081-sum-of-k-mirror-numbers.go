func kMirror(k int, n int) int64 {
    var ans int
    factor := int(math.Pow(float64(n), 9))
    pal := generatePalindromes(factor)
    sort.Ints(pal)
    cnt := 0
    for _, p := range pal {
        fmt.Println("====", p)
        converted := strconv.FormatInt(int64(p), k)
        num, _ := strconv.Atoi(converted)
        if checkPalindrome(num) == true {
            ans += p 
            cnt++
            fmt.Println(num)
        }
        // fmt.Println(cnt)
        if cnt == n {
            break
        }
        
    }
    // strconv.FormatInt(pal, 2)
    return int64(ans)
}

func createPalindrome(input, base int, isOdd bool) int {
    palin := input
	n := input

	if isOdd {
		n /= base
	}

	for n > 0 {
		palin = palin*base + n%base
		n /= base
	}
	return palin
}

func generatePalindromes(n int) []int {
    var palindromes []int
	max := n

	// Generate both even and odd length palindromes
	for isOdd := false; ; isOdd = !isOdd {
		i := 1
		for {
			palin := createPalindrome(i, 10, isOdd)
			if palin > max {
				break
			}
			palindromes = append(palindromes, palin)
			i++
		}

		// After generating both even and odd, break the loop
		if isOdd {
			break
		}
	}

	return palindromes
}

func checkPalindrome(num int) bool{
   input_num := num
   var remainder int
   res := 0
   for num>0 {
      remainder = num % 10
      res = (res * 10) + remainder
      num = num / 10
   }
   if input_num == res {
      return true
   } else {
      return false
   }
}