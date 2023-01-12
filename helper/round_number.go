package helper

import (
	"fmt"
	"strconv"
)

func ToFixed(num float64) float64 {

	y := fmt.Sprintf("%.2f", num)
	num, _ = strconv.ParseFloat(y, 64)

	return num
}
