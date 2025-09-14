package utils

import (
	"fmt"
	"strings"
)

func FormatMoney(num float64) string {
	intPart := int64(num)
	decPart := strings.Split(fmt.Sprintf("%.2f", num), ".")[1]
	intStr := fmt.Sprintf("%d", intPart)

	// Add commas
	var result []string
	for i, v := range intStr {
		if i > 0 && (len(intStr)-i)%3 == 0 {
			result = append(result, ",")
		}
		result = append(result, string(v))
	}

	return strings.Join(result, "") + "." + decPart
}
