package models

type Person struct {
	Name  string
	Age   int
	Email string
}

//Display to show details
func Display() (Person, error) {
	var person = Person{
		Name:  "chinedum",
		Age:   20,
		Email: "email@email.com",
	}
	return person, nil
}
