package shorten

import "strings"

type Shorten struct {
}

func New() *Shorten {
	return &Shorten{}
}

const alphabet = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ12345678"

var alphabetLen = uint32(len(alphabet))

func (s *Shorten) ShortenURL(id uint32) string {
	var (
		nums    []uint32
		num     = id
		builder strings.Builder
	)
	for num > 0 {
		nums = append(nums, num%alphabetLen)
		num /= alphabetLen
	}
	Reverse(nums)

	for _, num := range nums {
		builder.WriteString(string(alphabet[num]))
	}
	return builder.String()
}

func Reverse(s []uint32) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
