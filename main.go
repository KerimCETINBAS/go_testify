package main

import "fmt"

type UserController struct {
}

func (m *UserController) DoSomething(number int) (bool, error) {

	if number != 1 {
		fmt.Println("number 1 degil")
	}
	return true, nil
}

func main() {
}
