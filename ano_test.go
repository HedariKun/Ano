package ano

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {

}

func TestGenericMap(t *testing.T) {
	// Test one - test if numbers will be multiplied by 4
	list1 := []int{1, 2, 3, 4}
	checkList1 := []int{4, 8, 12, 16}
	ano1 := Wrap(list1)
	helper1 := ano1.GenericMap(func(element int) interface{} {
		return element * 4
	})
	ano1, _ = DefineType[int](helper1)
	for index, item := range ano1.Get() {
		if checkList1[index] != item {
			t.Fatalf("Map not working: Expected %d but got %d", checkList1[index], item)
		}
	}

	// Test two - test if type will change from number to string
	list2 := []int{1, 2, 3, 4}
	checkList2 := []string{"1.", "2.", "3.", "4."}
	ano2 := Wrap(list2)
	helper2 := ano2.GenericMap(func(element int) interface{} {
		return fmt.Sprintf("%d.", element)
	})
	ano3, _ := DefineType[string](helper2)
	for index, item := range ano3.Get() {
		if checkList2[index] != item {
			t.Fatalf("Map not working: Expected %s but got %s", checkList2[index], item)
		}
	}

	// Test three - test if type is missmatched

	list3 := []int{1, 2, 3, 4}
	ano4 := Wrap(list3)
	helper3 := ano4.GenericMap(func(element int) interface{} {
		return element * 2
	})
	_, err := DefineType[string](helper3)
	if err == nil {
		t.Fatalf("Map not working: didn't detect type missmatch")
	}
}

func TestFilter(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	checkList := []int{1, 2, 3, 4}
	ano := Wrap(list)
	ano = ano.Filter(func(element int) bool { return element < 5 })
	for index, element := range checkList {
		if ano.Get()[index] != element {
			t.Fatalf("Filter not working: expected %d but got %d", element, ano.Get()[index])
		}
	}
	if len(ano.Get()) != len(checkList) {
		t.Fatal("Filter not working: the filtered list doesn't match the expected list")
	}
}
