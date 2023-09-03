package utils

func InListStr(element string, list []string) bool {
	for _, item := range list {
		if element == item {
			return true

		}
	}
	return false
}
