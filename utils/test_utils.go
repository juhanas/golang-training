package utils

// CompareLists returns if the two lists contain the same data
func CompareLists(list1, list2 []string) bool {
	if list1 == nil && list2 != nil || list1 != nil && list2 == nil {
		return false
	}
	if len(list1) != len(list2) {
		return false
	}
	for i := 0; i < len(list1); i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// CopyMap deep-copies the given map
func CopyMap(orig map[string]string) map[string]string {
	new := map[string]string{}
	for k, v := range orig {
		new[k] = v
	}
	return new
}
