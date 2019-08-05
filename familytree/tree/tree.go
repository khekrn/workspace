package tree

import (
	"errors"
	"familytree/person"
)

// Error codes
var (
	ErrUniqueName = errors.New("person already exist, please add new person")

	ErrPersonNotFound = errors.New("person not found in family tree")

	ErrAddingChildFailed = errors.New("expected gender to be female but found male")

	ErrChildExist = errors.New("child already exist in the family tree")
)

// Tree interface defines contract for family tree implementation
type Tree interface {
	Add(psn person.Person) error
	GetPerson(name string) (person.Person, bool)
	ListPeople() map[string]person.Person
	AddChild(motherName, childName, gender string) error
	Exist(name string) bool
}

// FamilyTree holds family tree structure
type FamilyTree struct {
	people map[string]person.Person
}

// NewFamilyTree creates the new instance of FamilyTree
func NewFamilyTree() Tree {
	return &FamilyTree{people: map[string]person.Person{}}
}

// Add adding new people to the family tree
func (ft *FamilyTree) Add(psn person.Person) error {
	if _, found := ft.people[psn.Name()]; found {
		return ErrUniqueName
	}
	ft.people[psn.Name()] = psn
	return nil
}

// GetPerson gives you the Person struct if the given name exist
func (ft *FamilyTree) GetPerson(name string) (person.Person, bool) {
	if psn, found := ft.people[name]; found {
		return psn, true
	}
	return nil, false
}

// ListPeople returns map of pepole in the tree
func (ft *FamilyTree) ListPeople() map[string]person.Person {
	return ft.people
}

// AddChild adding given child to the given parent
func (ft *FamilyTree) AddChild(motherName, childName, gender string) error {
	mother, found := ft.GetPerson(motherName)
	if !found {
		return ErrPersonNotFound
	}

	if mother.Gender() != "female" {
		return ErrAddingChildFailed
	}

	if ft.Exist(mother.Name()) {
		return ErrChildExist
	}

	newChild := person.NewPerson(childName, gender)
	newChild.SetMother(mother)

	if mother.HasSpouse() {
		newChild.SetFather(mother.Spouse())
	}

	ft.Add(newChild)

	return nil
}

// Exist is user exist
func (ft *FamilyTree) Exist(name string) bool {
	if _, found := ft.people[name]; found {
		return found
	}
	return false
}
