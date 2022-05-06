package homework

import "sort"

/*
Task 3: Maps
function that returns map values sorted in order of increasing keys.
Input -> {2: "a", 0: "b", 1: "c"}
Output -> ["b", "c", "a"]
Input -> {10: "aa", 0: "bb", 500: "cc"}
Output -> ["bb", "aa", "cc"]
*/
func sortMapValues(input map[int]string) (result []string) {
	var keys []int
	for k := range input {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for i := 0; i < len(keys); i++ {
		result = append(result, input[keys[i]])
	}

	return
}
