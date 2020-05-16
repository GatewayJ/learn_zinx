package main

import "fmt"

func main() {
	var t *test = nil
	t.Name = "ACE"
	t.test()
}

type test struct {
	Name string
}

func (this_ test) test() {
	fmt.Println(this_.Name)
}
