package utils

import "strconv"

func StrToUInt(s string) (uint, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}
