package service

import (
	"fmt"
	"reflect"
	"testing"

	"gymshark-test/internal/models"
)

func Test(t *testing.T) {
	type test struct {
		in  uint
		out map[*models.Package]uint
	}

	packSizes := []*models.Package{
		{
			ID:   0,
			Size: 250,
		},
		{
			ID:   1,
			Size: 500,
		},
		{
			ID:   2,
			Size: 1000,
		},
		{
			ID:   3,
			Size: 2000,
		},
		{
			ID:   4,
			Size: 5000,
		},
	}

	tests := []test{
		{
			in: 12501,
			out: map[*models.Package]uint{
				packSizes[4]: 2,
				packSizes[3]: 1,
				packSizes[1]: 1,
				packSizes[0]: 1,
			},
			//[]int{5000, 5000, 2000, 500, 250},
		},
		{
			in: 501,
			out: map[*models.Package]uint{
				packSizes[1]: 1,
				packSizes[0]: 1,
			},
			//[]int{500, 250},
		},
		{
			in: 1,
			out: map[*models.Package]uint{
				packSizes[0]: 1,
			},
			//[]int{250},
		},
		{
			in: 250,
			out: map[*models.Package]uint{
				packSizes[0]: 1,
			},
			//[]int{250},
		},
		{
			in: 251,
			out: map[*models.Package]uint{
				packSizes[1]: 1,
			},
			//[]int{500},
		},
		{
			in: 499,
			out: map[*models.Package]uint{
				packSizes[1]: 1,
			},
			//[]int{500},
		},
		{
			in: 1001,
			out: map[*models.Package]uint{
				packSizes[2]: 1,
				packSizes[0]: 1,
			},
			//[]int{1000, 250},
		},
	}

	for _, unit := range tests {
		result := calculatePacks(packSizes, unit.in)
		if !reflect.DeepEqual(result, unit.out) {
			fmt.Println(unit, result)
			panic("error")
		}
	}
}
