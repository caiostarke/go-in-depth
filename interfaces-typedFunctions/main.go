package main

import (
	"fmt"
	"strings"
)

// ****************************************************************
// func transformString(s string) string {
// 	return strings.ToUpper(s)
// }
// It's bad, cuz we have a dependency on strings.ToUpper, instead we do

type stringTransformer func(s string) string

func transformString(s string, fn stringTransformer) string {
	return fn(s)
}

func Upper(s string) string {
	return strings.ToUpper(s)
}

func Lower(s string) string {
	return strings.ToLower(s)
}

func Prefixer(prefix string) stringTransformer {
	return func(s string) string {
		return prefix + s
	}
	// bad too, dependending on 'PREFIX_ hardcoded value'
	// return "PREFIX_" + s
}

func main() {
	fmt.Println(transformString("Hello my friend!", Upper))
	fmt.Println(transformString("Hello my friend!", Lower))
	fmt.Println(transformString("Hello my friend!", Prefixer("PREFIX_")))
	fmt.Println(transformString("Hello my friend!", Prefixer("2023_")))
}
