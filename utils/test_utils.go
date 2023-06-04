package utils

import (
	"github.com/juhanas/golang-training/models"
)

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

// CompareDogs returns if the two lists of dogs contain the same data
func CompareDogs(list1, list2 []models.Dog) bool {
	if list1 == nil && list2 != nil || list1 != nil && list2 == nil {
		return false
	}
	if len(list1) != len(list2) {
		return false
	}
	for i := 0; i < len(list1); i++ {
		if list1[i].Name != list2[i].Name {
			return false
		}
		if list1[i].Color != list2[i].Color {
			return false
		}
		if list1[i].Pack != list2[i].Pack {
			return false
		}
	}
	return true
}

// CreateDog creates a Dog struct with the given name and other placeholder values
func CreateDog(name string) models.Dog {
	return models.Dog{
		Animal: models.Animal{
			Name:  name,
			Color: "black",
		},
		Pack: "Who let the dogs out?",
	}
}
