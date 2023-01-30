package storage

import "strings"

const alphabet = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ12345678"

var alphabetLen = uint32(len(alphabet))

func ShortenURL(id uint32) string {
	var nums []uint32
	for num := id; num > 0; {
		nums = append(nums, num%alphabetLen)
		num /= alphabetLen
	}
	reverse(nums)

	var builder strings.Builder
	for _, num := range nums {
		builder.WriteString(string(alphabet[num]))
	}
	return builder.String()
}

func reverse(s []uint32) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
