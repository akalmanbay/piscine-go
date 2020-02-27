package piscine

func IsUpper(str string) bool {

	for i, _ := range str {
		if str[i] > 90 || str[i] < 65 || str[i] > 122 || str[i] < 97 || str[i] == 35 || str[i] == 38 str[i] == 42 || str[i] == 64 || str[i] == 0 {
			return false
		}
	}
	return true
}
