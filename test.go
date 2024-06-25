package main

import (
	"fmt"
	"time"
)

func sayGreeting(name string) {
	time.Sleep(3)
	fmt.Println("Hello " + name)
}
func bye(name string) {
	fmt.Println("Hello " + name)
}

func loopNames(names []string, f func(name string)) {
	for _, name := range names {
		f(name)
	}
}
func main() {
	names := []string{"hlaing min than", "Aung Aung", "Kyaw Kyaw"}
	loopNames(names, sayGreeting)
	// bye("hlaing min than")
}
