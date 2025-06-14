func intToRoman(num int) string {
	var ans strings.Builder
	m := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}

	// Extract digits
	th := num / 1000
	h := (num / 100) % 10
	t := (num / 10) % 10
	u := num % 10

	// Convert each digit
	if th > 0 {
		ans.WriteString(strings.Repeat(m[1000], th))
	}
	if h > 0 {
		ans.WriteString(convertDigit(h, 100, m))
	}
	if t > 0 {
		ans.WriteString(convertDigit(t, 10, m))
	}
	if u > 0 {
		ans.WriteString(convertDigit(u, 1, m))
	}

	return ans.String()
}

func convertDigit(digit, place int, m map[int]string) string {
	var res strings.Builder
	base := m[place]
	fiveBase := m[place*5]
	tenBase := m[place*10]

	switch {
	case digit == 9:
		res.WriteString(base + tenBase)
	case digit >= 5:
		res.WriteString(fiveBase)
		res.WriteString(strings.Repeat(base, digit-5))
	case digit == 4:
		res.WriteString(base + fiveBase)
	default:
		res.WriteString(strings.Repeat(base, digit))
	}
	return res.String()
}