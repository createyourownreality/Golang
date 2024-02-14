package main

import "fmt"

type str string // str is a kind of string

func (txt str) log() { // here txt is the name of the str

	fmt.Println(txt)
}

func main() {

	var name str = "Max"

	name.log()

}
