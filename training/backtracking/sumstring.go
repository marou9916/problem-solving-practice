package backtracking

func CalculateSum(s1 string, s2 string) string {
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}

	m := len(s1)
	n := len(s2)
	carry := 0
	sum := ""

	for i := 0; i < m; i++ {
		ds := (int(s1[m-1-i]-'0') + int(s2[n-1-i]-'0') + carry) % 10

		carry = (int(s1[m-1-i]-'0') + int(s2[n-1-i]-'0') + carry) / 10

		sum = string(ds + '0') + sum
	}

	for i := m; i < n; i++ {
		ds := (int(s1[m-1-i]-'0') + carry) % 10

		carry = (int(s1[m-1-i]-'0') + carry) / 10

		sum = string(ds + '0') + sum
	}

	if carry > 0 {
		sum = string(carry + '0') + sum
	}

	return sum

}

func CheckStringSum(s string, beg int, len1 int, len2 int) bool {
	s1 := s[beg : beg+len1]
	s2 := s[beg+len1 : beg+len1+len2]
	s3 := CalculateSum(s1, s2)

	if len(s3) > len(s)-(beg+len1+len2) {
		return false
	}

	if s3 == s[beg+len(s1)+len(s2):beg+len(s1)+len(s2)+len(s3)] {
		if beg+len(s1)+len(s2)+len(s3) == len(s) {
			return true
		}

		return CheckStringSum(s, beg, len2, len(s3))
	}

	return false
}

func StringSum(s string) bool {
	for i := 1; i < len(s); i++ {
		for j := 1; i+j < len(s); j++ {
			if CheckStringSum(s, 0, i, j) {
				return true
			}
		}
	}
	return false
}
