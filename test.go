package main

import (
	"fmt"
	"strings"
)

func initials(name string) (string, string) {
	initials := []string{}
	splittedNames := strings.Split(name, " ")
	for _, name := range splittedNames {
		upperCase := strings.ToUpper(name)
		initials = append(initials, upperCase[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

func main() {
	name1 := "Hlaing Min"
	name2 := "Zaw"

	s1, s2 := initials(name1)
	fmt.Println(s1, s2)

	s3, s4 := initials(name2)
	fmt.Println(s3, s4)
}
