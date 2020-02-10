package main

func contains(arr []string, target string) bool {
	for _, element := range arr {
		if target == element {
			return true
		}
	}
	return false
}
