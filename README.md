# Ano
A way to wrap arrays in golang to be able to manipulate them without writing any additional functions or methods

# Usage
you can start using it by adding the module using `go get github.com/hedarikun/ano` then you can import it in your project.

## Example usage of Ano
```go 
package main

import (
	"github.com/hedarikun/ano"
)

func main() {
	list := []int{1, 2, 3, 4, 5}
	list = ano.Wrap(list).Map(func(element int) int { return element * 2 }).Get()
	for _, item := range list {
		println(item)
	}
}
```
# What functions does it have
- `Map` to Map the values of a list to different Values with the same type
- `GenericMap` to Map the values of a list to different Values With different Type - it will return anoHelper object that you can get a new Ano object with the desired Type by using `ano.DefineType` Function
- `Filter` to Filter the values of a list based on a rule
