package storage

import "strings"

const alphabet = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ12345678"

var alphabetLen = len(alphabet)

func ShortenURL(id uint) string {
	var nums []uint
	for num := id; num > 0; {
		nums = append(nums, num%uint(alphabetLen))
		num /= uint(alphabetLen)
	}
	reverse(nums)

	var builder strings.Builder
	for _, num := range nums {
		builder.WriteString(string(alphabet[num]))
	}
	return builder.String()
}

func reverse(s []uint) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
