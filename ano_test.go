package ano

import (
	"fmt"
	"testing"
)

type user struct {
	Name     string
	canDrive bool
}

func TestM(t *testing.T) {
	u1 := user{Name: "mike", canDrive: true}
	u2 := user{Name: "jessica", canDrive: false}
	u3 := user{Name: "atrox", canDrive: true}
	u4 := user{Name: "rawa", canDrive: false}

	list := []user{u1, u2, u3, u4}
	ano := Wrap(list)
	ano = ano.Filter(func(e user) bool { return !e.canDrive }).Sort(func(a, b user) bool { return a.Name < b.Name })
	fmt.Printf("%+v", ano.Get())
}

func TestMap(t *testing.T) {
	list := []int{1, 2, 3, 4}
	checkList := []int{2, 4, 6, 8}
	ano := Wrap(list).Map(func(element int) int { return element * 2 })
	for index, element := range ano.Get() {
		if element != checkList[index] {
			t.Fatalf("Map not working: Expected %d but got %d", checkList[index], element)
		}
	}
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

func TestSort(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fatalf("Sort not working: %v", err)
		}
	}()

	// sort runes alphabatically
	list1 := []int{5, 7, 12, 4, 3, 6, 8, 9, 1, 0, 10, 11, 2}
	sortedList1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	ano1 := Wrap(list1)
	ano1 = ano1.Sort(func(A, B int) bool {
		return A < B
	})
	anoList1 := ano1.Get()
	for index, element := range sortedList1 {
		if element != anoList1[index] {
			t.Fatalf("Sort not working: expected %d got %d", element, anoList1[index])
		}
	}

	list2 := []rune{'a', 'c', 'd', 'p', 'b', 'e', 'i', 'g', 'h', 'f', 'j', 'x', 'y', 'z', 'l', 'k', 'o', 'q', 'm', 'n', 'r', 'w', 'u', 's', 't', 'v'}
	sortedList2 := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	ano2 := Wrap(list2)
	ano2 = ano2.Sort(func(A, B rune) bool {
		return A < B
	})
	anoList2 := ano2.Get()
	for index, element := range sortedList2 {
		if element != anoList2[index] {
			t.Fatalf("Sort not working: expected %v got %v", string(element), string(anoList2[index]))
		}
	}
}

func TestIntersect(t *testing.T) {
	listA := []int{1, 2, 3, -123, 4, 5, 54234}
	listB := []int{7, 5, 54234, -123, 8, 4, 9}
	checkList := []int{-123, 4, 5, 54234}
	anoList := Wrap(listA).Intersect(listB, func(element int) any { return element }).Sort(func(a, b int) bool { return a < b }).Get()
	if len(checkList) != len(anoList) {
		t.Fatalf("Intersect not working: expected size of %v but got %v", len(checkList), len(anoList))
	}
	for index, element := range checkList {
		if anoList[index] != element {
			t.Fatalf("Intersect not working: expected %v but got %v", element, anoList[index])
		}
	}
}

func TestIntersectWithStruct(t *testing.T) {
	type User struct {
		id   int
		name string
	}
	user1 := User{id: 1, name: "kevin"}
	user2 := User{id: 2, name: "sarah"}
	user3 := User{id: 3, name: "paul"}
	user4 := User{id: 4, name: "mike"}
	user5 := User{id: 5, name: "rawa"}
	user6 := User{id: 6, name: "ali"}
	user7 := User{id: 7, name: "user"}
	user8 := User{id: 8, name: "mike"}

	teamA := []User{user1, user4, user7, user3, user6}
	teamB := []User{user7, user8, user1, user2, user5, user3}

	checkTeam := []User{user1, user3, user7}

	intersectUsers := Wrap(teamA).Intersect(teamB, func(element User) any { return element.id }).Sort(func(A, B User) bool { return A.id < B.id }).Get()
	if len(checkTeam) != len(intersectUsers) {
		t.Fatalf("Intersection not working: expected length %v got %v", len(checkTeam), len(intersectUsers))
	}
	for index, element := range checkTeam {
		if intersectUsers[index].id != element.id {
			t.Fatalf("Intersection not working: expected element %v got %v", element, intersectUsers[index])
		}
	}

}

func TestUnion(t *testing.T) {
	listA := []int{1, 2, 3, 4, 5, 6}
	listB := []int{4, 5, 6, 7, 8, 9}
	checkList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	unionList := Wrap(listA).Union(listB, func(element int) any { return element }).Get()

	if len(unionList) != len(checkList) {
		t.Fatalf("Union not working: expected length of %v but got %v", len(checkList), len(unionList))
	}
	for index, element := range checkList {
		if element != unionList[index] {
			t.Fatalf("Union not working: expected %v but got %v", element, unionList[index])
		}
	}
}
