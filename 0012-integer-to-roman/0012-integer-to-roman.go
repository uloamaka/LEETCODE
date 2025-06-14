func intToRoman(num int) string {
    var th,h,t,u int
    var ans strings.Builder
    u = num%10; 
    t = (num/10)%10; 
    h = (num/100)%10; 
    th = num/1000;
    
    m := map[int]string{
        1 : "I",
        5 : "V",
        10 : "X",
        50 : "L",
        100 : "C",
        500 : "D",
        1000 : "M",
    }
    if th != 0 {
        val := m[1000]
        for i := 0; i < th; i++ {
            ans.WriteString(val)
        }
    }

    if h != 0 {
		ans.WriteString(convertDigit(h, 100, m))
	}

	if t != 0 {
		ans.WriteString(convertDigit(t, 10, m))
	}

	if u != 0 {
		ans.WriteString(convertDigit(u, 1, m))
	}

	return ans.String()
}

func convertDigit(digit, place int, m map[int]string) string {
	var res strings.Builder
	symbols := map[int]string{
		place:     m[place],
		place * 5: m[place*5],
		place * 10: m[place*10],
	}

	switch {
	case digit == 9:
		res.WriteString(symbols[place] + symbols[place*10])
	case digit >= 5:
		res.WriteString(symbols[place*5])
		for i := 0; i < digit-5; i++ {
			res.WriteString(symbols[place])
		}
	case digit == 4:
		res.WriteString(symbols[place] + symbols[place*5])
	default:
		for i := 0; i < digit; i++ {
			res.WriteString(symbols[place])
		}
	}
	return res.String()
}