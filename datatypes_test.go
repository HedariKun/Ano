package ano

import (
	"testing"
)

func TestNode(t *testing.T) {
	// smaller to bigger
	list1 := []int{1, 5, 4, 3, 9, 12, 62}
	sortList1 := []int{1, 3, 4, 5, 7, 9, 12, 62}
	compare1 := func(A, B int) bool {
		return A < B
	}

	n1 := node[int]{value: 7}
	for _, element := range list1 {
		n1.Add(element, compare1)
	}
	for index, element := range n1.GetList() {
		if sortList1[index] != element {
			t.Fatalf("Node not working: expected %v got %v", sortList1[index], element)
		}
	}

	// nigger to smaller
	list2 := []int{1, 5, 4, 3, 9, 12, 62}
	sortList2 := []int{62, 12, 9, 7, 5, 4, 3, 1}
	compare2 := func(A, B int) bool {
		return A > B
	}

	n2 := node[int]{value: 7}
	for _, element := range list2 {
		n2.Add(element, compare2)
	}
	for index, element := range n2.GetList() {
		if sortList2[index] != element {
			t.Fatalf("Node not working: expected %v got %v", sortList2[index], element)
		}
	}
}
