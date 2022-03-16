package testdata

type Person struct {
	Name string
}

type PeopleRepo struct{}

func (r *PeopleRepo) AddPerson(person *Person) {}
