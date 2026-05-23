package Utils

import (
	"fmt"
	"strings"
)

func PtrTernary[T any](p *T, fallback T) T{
	//
	if p == nil {
		return fallback
	}

	return *p
}

func ArrayToString[T int | string](arr []T) string{
	//
	var str strings.Builder

	for i, item := range arr{
		val := fmt.Sprint(item)
		if i < len(arr) - 1{
			val = val + ", "
		}

		fmt.Fprint(&str, val)
	}

	return str.String()
}
