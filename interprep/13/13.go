package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["key"] = 42

	delete(m, "key")
	fmt.Println(m["key"])
}
