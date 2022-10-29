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
