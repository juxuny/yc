package validator

// Luhn算法验证银行卡号合法性
func luhn(s string) bool {
	var t = []int{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
	odd := len(s) & 1
	var sum int
	for i, c := range s {
		if c < '0' || c > '9' {
			return false
		}
		if i&1 == odd {
			sum += t[c-'0']
		} else {
			sum += int(c - '0')
		}
	}
	return sum%10 == 0
}

func IsBankCardAccount(in string) bool {
	return luhn(in)
}
