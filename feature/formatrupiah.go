package feature

import (
	"fmt"
	"strings"
)

func FormatRupiah(angka int) string {
	str := fmt.Sprintf("%d", angka)
	var result []string

	// Memecah string dari belakang untuk menambahkan titik setiap 3 digit
	for i := len(str); i > 0; i -= 3 {
		start := i - 3
		if start < 0 {
			start = 0
		}
		result = append([]string{str[start:i]}, result...)
	}

	return "Rp " + strings.Join(result, ".")
}