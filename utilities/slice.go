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

func StringSliceUnique(data []string) []string {
	var unique []string
	for i := 0; i < len(data); i++ {
		letter := data[i]
		if letter != "" && StringSliceIndexOf(letter, unique) == -1 {
			unique = append(unique, data[i])
		}
	}
	return unique
}

func IntSliceIndexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
