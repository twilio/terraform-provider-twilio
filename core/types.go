package core

import "strconv"

func IntToString(input int) string {
	return strconv.Itoa(input)
}

func StringToInt(input string) (int, error) {
	return strconv.Atoi(input)
}
