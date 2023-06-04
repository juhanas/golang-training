package models

import "fmt"

// Animal shares the common characteristics among all pets
type Animal struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

type Cat struct {
	Animal
}

type Dog struct {
	Animal
	Pack string `json:"pack"`
}

// GetArrIdx returns the idx of the object in the given slice (array)
//Returns an error if item is not found
func (item Dog) GetArrIdx(slice []Dog) (int, error) {
	for idx, v := range slice {
		if v.Name == item.Name {
			return idx, nil
		}
	}
	return 0, fmt.Errorf("no index found for item %v in slice %v", item, slice)
}
