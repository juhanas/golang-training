package utils

import (
	"fmt"
	"sort"
)

// SortKeys returns a list of the keys in the given map, sorted alphabetically
func SortKeys(data map[string]string) []string {
	keys := []string{}
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// FindNameKey returns the key for the given name
// Returns error if name is reserved or key could not be found
func FindNameKey(name string, dict map[string]string, i int) (string, error) {
	if i > len(name) {
		return "", fmt.Errorf("index out of bounds")
	}
	if i == 0 {
		return "", fmt.Errorf("index must be greater than 0")
	}

	key := name[:i]
	_, ok := dict[key]
	if !ok {
		return key, nil
	}
	return FindNameKey(name, dict, i+1)
}

// ParseCSVDataToArray parses the given data. Data should be separated by comma: ','.
func ParseCSVDataToArray(data []byte) []string {
	final := []string{}
	str := ""
	for _, b := range data {
		if b != ',' {
			str += string(b)
			continue
		}
		final = append(final, str)
		str = ""
	}
	final = append(final, str)
	return final
}
