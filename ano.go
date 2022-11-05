package ano

import "errors"

type Ano[T any] struct {
	list []T
}

type anoHelper struct {
	list []any
}

func Wrap[T any](items []T) Ano[T] {
	ano := Ano[T]{list: items}
	return ano
}

func DefineType[T any](helper anoHelper) (Ano[T], error) {
	ano := Ano[T]{}
	list := []T{}
	for _, element := range helper.list {
		item, ok := element.(T)
		if !ok {
			return ano, errors.New("Type Mismatch")
		}
		list = append(list, item)
	}
	ano.list = list
	return ano, nil
}

func (ano Ano[T]) Get() []T {
	return ano.list
}

func (ano Ano[T]) Map(mapFunc func(element T) T) Ano[T] {
	list := []T{}
	for _, element := range ano.Get() {
		item := mapFunc(element)
		list = append(list, item)
	}
	return Ano[T]{list: list}
}

func (ano Ano[T]) GenericMap(mapFunc func(element T) interface{}) anoHelper {
	list := []any{}
	for _, element := range ano.list {
		item := mapFunc(element)
		list = append(list, item)
	}
	return anoHelper{list: list}
}

func (ano Ano[T]) Filter(filterFunc func(element T) bool) Ano[T] {
	newAno := Ano[T]{}
	list := []T{}
	for _, element := range ano.Get() {
		check := filterFunc(element)
		if check {
			list = append(list, element)
		}
	}
	newAno.list = list
	return newAno
}

func (ano Ano[T]) Sort(sortFunc func(A, B T) bool) Ano[T] {
	newAno := Ano[T]{}
	var mainNode *node[T]
	for _, element := range ano.Get() {
		if mainNode == nil {
			mainNode = &node[T]{value: element}
		} else {
			mainNode.Add(element, sortFunc)
		}
	}

	newAno.list = mainNode.GetList()
	return newAno
}

func (ano Ano[T]) Intersect(otherList []T, intersectFunc func(element T) any) Ano[T] {

	objectsMap := make(map[any]bool)
	intersectedList := []T{}

	for _, element := range otherList {
		key := intersectFunc(element)
		_, ok := objectsMap[key]
		if !ok {
			objectsMap[key] = false
		}
	}

	for _, element := range ano.list {
		key := intersectFunc(element)
		_, ok := objectsMap[key]
		if ok {
			intersectedList = append(intersectedList, element)
		}
	}

	return Wrap(intersectedList)
}

// Todo: add strict intersect for structs to make sure that the data should be the same

func (ano Ano[T]) Union(otherList []T, unionFunc func(element T) any) Ano[T] {
	list := []T{}
	elementMap := make(map[any]bool)

	for _, element := range ano.Get() {
		list = append(list, element)
		key := unionFunc(element)

		_, ok := elementMap[key]
		if !ok {
			elementMap[key] = false
		}
	}

	for _, element := range otherList {
		key := unionFunc(element)

		_, ok := elementMap[key]
		if !ok {
			list = append(list, element)
		}
	}

	return Wrap(list)
}
