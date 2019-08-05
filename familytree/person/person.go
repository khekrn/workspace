package person

import "strings"

// Person interface defines person each individual data
type Person interface {
	Name() string
	Gender() string
	Mother() Person
	Father() Person
	Spouse() Person
	HasSpouse() bool
	Children() []Person
	Sibling() []Person

	AddChild(child Person)
	SetMother(mother Person)
	SetFather(father Person)
	SetSpouse(spuse Person)
}

// person data
type person struct {
	name     string
	gender   string
	mother   Person
	father   Person
	spouse   Person
	children []Person
}

// NewPerson creates new Person with given name and gender
func NewPerson(name, gender string) Person {
	return &person{name: name, gender: gender}
}

// AddChild new child will be added to the person
func (p *person) AddChild(child Person) {
	for _, ch := range p.children {
		if ch == child {
			return
		}
	}
	p.children = append(p.children, child)
}

// SetMother updates the current person mother
func (p *person) SetMother(mother Person) {
	p.mother = mother
	p.mother.AddChild(p)
}

// SetFather updates the current person father
func (p *person) SetFather(father Person) {
	p.father = father
	p.father.AddChild(p)
}

// SetSpouse updates the current person spouse
func (p *person) SetSpouse(spouse Person) {
	p.spouse = spouse
}

// Name returns person name
func (p *person) Name() string {
	return p.name
}

// Gender returns person gender in lowercase
func (p *person) Gender() string {
	return strings.ToLower(p.gender)
}

// Mother person mother
func (p *person) Mother() Person {
	return p.mother
}

// Father person father
func (p *person) Father() Person {
	return p.father
}

// Spouse person spouse
func (p *person) Spouse() Person {
	return p.spouse
}

// HasSpouse returns whether the person has spouse or not
func (p *person) HasSpouse() bool {
	return p.spouse != nil
}

// Children person childrens
func (p *person) Children() []Person {
	return p.children
}

// Sinling person siblings
func (p *person) Sibling() []Person {
	siblings := make([]Person, 0)
	if p.mother != nil {
		for _, child := range p.Mother().Children() {
			if child.Name() == p.name {
				continue
			}
			siblings = append(siblings, child)
		}
	}
	return siblings
}
