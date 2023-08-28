package service

import "gymshark-test/internal/models"

func calculatePacks(packages []*models.Package, itemsOrdered uint) map[*models.Package]uint {
	result := map[*models.Package]uint{}
	j := len(packages) - 1

	for j >= 0 {
		if itemsOrdered >= packages[j].Size {
			itemsOrdered -= packages[j].Size
			result[packages[j]]++
		} else if j == 1 && itemsOrdered > packages[0].Size {
			result[packages[j]]++
			break
		} else if j == 0 && itemsOrdered > 0 {
			result[packages[j]]++
			j--
		} else {
			j--
		}
	}

	return result
}
