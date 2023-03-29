package main

import (
	"errors"
	"fmt"
)

var err1 = errors.New("Error: el salario es menor a 10.000")

func main() {
	salary := 50
	//e := Error()
	err := esMenor(salary)
	if err != nil {
		if errors.Is(err, err1) {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(salary)
}

func esMenor(salary int) error {
	if salary <= 10000 {
		return err1
	}
	return nil
}
