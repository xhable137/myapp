package models

type Person struct {
	Name string
	Age  int
}

func (p *Person) Greet() string {
	return "Hello, my name is " + p.Name
}

func (p *Person) HaveBirthday() int {
	p.Age++
	return p.Age
}
