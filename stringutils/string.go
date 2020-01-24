package stringutils

func Contains(x string, list []string) bool {
	for _, y := range list {
		if x == y {
			return true
		}
	}
	return false
}
