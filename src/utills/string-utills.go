package utills

func Check_prefix(str, prefix string) bool {
	strB := []byte(str)
	prefixB := []byte(prefix)

	if len(prefixB) > len(strB) {
		return false
	}

	for i := range prefixB {
		if prefixB[i] != strB[i] {
			return false
		}
	}

	return true
}

func Rem_prefix(str, prefix string) string {
	strB := []byte(str)
	prefixB := []byte(prefix)

	if len(prefixB) > len(strB) {
		return ""
	}

	for i := range prefixB {
		if prefixB[i] != strB[i] {
			return string(strB[:i])
		}
	}
	return string(strB[len(prefixB):])
}
