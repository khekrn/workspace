package person

import (
	"fmt"
	"testing"
)

type personStub struct {
	name   string
	gender string
}

var testPersons = []struct {
	description string
	stub        personStub
}{
	{
		"random male",
		personStub{"John Doe", "male"},
	},
	{
		"random female",
		personStub{"Lilliana Medina", "female"},
	},
}

func TestNewPerson(t *testing.T) {
	for i, test := range testPersons {
		mother := NewPerson(fmt.Sprintf("mother%d", i), "female")
		father := NewPerson(fmt.Sprintf("father%d", i), "male")
		spouse := NewPerson(fmt.Sprintf("spouse%d", i), "female")
		person := NewPerson(test.stub.name, test.stub.gender)
		if person.Name() != test.stub.name {
			t.Fatalf("failed to assert that %s is %s", person.Name(), test.stub.name)
		}
		if person.Gender() != test.stub.gender {
			t.Fatalf("failed to assert that %s is %s", person.Gender(), test.stub.gender)
		}
		person.SetMother(mother)
		if person.Mother() != mother {
			t.Fatalf("failed to assert that %s is %s", person.Gender(), test.stub.gender)
		}
		person.SetFather(father)
		if person.Father() != father {
			t.Fatalf("failed to assert that %s is %s", person.Gender(), test.stub.gender)
		}
		person.SetSpouse(spouse)
		if person.Spouse() != spouse {
			t.Fatalf("failed to assert that %s is %s", person.Spouse().Name(), spouse.Name())
		}
		if person.HasSpouse() == false {
			t.Fatalf("failed to assert that a spouse exists")
		}
		if len(person.Children()) != 0 {
			t.Fatalf("failed to assert that %s is %s", person.Gender(), test.stub.gender)
		}
		if len(person.Sibling()) != 0 {
			t.Fatal("failed to assert that the person has no siblings.")
		}
	}
}

func TestAddChild(t *testing.T) {
	person := NewPerson("parent", "female")
	child := NewPerson("child", "female")
	if len(person.Children()) != 0 {
		t.Fatal("failed to assert that the person has zero children")
	}
	person.AddChild(child)
	if len(person.Children()) != 1 {
		t.Fatal("failed to assert that the person has 1 child")
	}
	if person.Children()[0].Name() != "child" {
		t.Fatal("failed to assert the child's name is child")
	}
	// attempting to add the same child twice does nothing.
	person.AddChild(child)
	if len(person.Children()) != 1 {
		t.Fatal("failed to assert that the person has 1 child")
	}
}

func TestSibling(t *testing.T) {
	mother := NewPerson("mother", "female")
	father := NewPerson("father", "male")
	child1 := NewPerson("child1", "male")
	child1.SetFather(father)
	child1.SetMother(mother)
	child2 := NewPerson("child2", "female")
	child2.SetFather(father)
	child2.SetMother(mother)
	if len(child1.Sibling()) != 1 {
		t.Fatal("failed to assert the child has one sibling")
	}
	if child1.Sibling()[0].Name() != child2.Name() {
		t.Fatal("failed asserting that the child1's sibling is child2")
	}
}
