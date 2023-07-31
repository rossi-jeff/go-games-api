package utilities

func DeleteStringSliceIndex(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func StringSliceIndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func IntSliceIndexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
