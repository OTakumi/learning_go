package main

import (
	"fmt"

	"github.com/OTakumi/learning_go/internal/model"
)

func main() {
	person := model.MakePerson("aaa", "bbb", 10)

	personPointer := model.MakePersonPointer("aaa", "bbb", 10)

	if person.FirstName != "" || personPointer != nil {
		fmt.Println("ok")
	}
}
