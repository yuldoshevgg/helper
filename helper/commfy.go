package helper

import "strings"

func Commfy(s string) string {
	var arr []string

	var desimal string

	if strings.Contains(s, ".") {
		desimal = strings.Split(s, ".")[1]
		s = strings.Split(s, ".")[0]
	}

	res1 := func(s string) string {
		for len(s) > 3 {
			s_byte := []byte(s)
			s_len := len(s)
			_nums := string(s_byte[s_len-3:])
			arr = append(arr, _nums)
			s = string(s_byte[:s_len-3])
		}

		arr = append(arr, s)
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}
		res := strings.Join(arr, " ")
		return res
	}(s)

	return res1 + "." + desimal
}
