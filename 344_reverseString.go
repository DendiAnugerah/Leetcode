package leetcode

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1
	for i < j {
		buf := s[j]
		s[j] = s[i]
		s[i] = buf
		i++
		j--
	}
}
