package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	// secretAgent "is-a" person, so include person as a field type, but don't include a field name
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) {} { code...}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

// func (r receiver) identifier(parameters) (return(s)) {} { code...}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

// a type implements an interface by implementing its methods
// any type that has the method speak, is ALSO of type human, and therefore implements
// the human interface
type human interface {
	speak()
}

// now we can create a generalized/abstract func that USES the INTERFACE as a parameter, which
// allows us to pass the [super]type HUMAN into bar and then just call the bar func (which uses an interface for a type)
// allowing you to pass in different types (see bar calls at bottom)
func bar(h human) {
	fmt.Println("I was passed into bar", h)
}

func main() {
	sa1 := secretAgent{ //the below body is the VALUE of TYPE secretAgent
		//and is assigned to the identifier sa1, AND is also
		//of the type interface (i.e. VALUES can be MORE THAN ONE TYPE!!)
		//This means that this body/value, is of type secretAgent AND
		//type human, because the method speak() is attached to it.
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	// can call bar because values of type secretAgent and person are also of type human, and
	// we've asked type bar to take in humaN (all of this is polymorphism)
	bar(sa1)
	bar(sa2)
	bar(p1)

}
