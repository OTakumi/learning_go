package model

// Person model

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	p := Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

	return p
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	p := Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

	return &p
}
