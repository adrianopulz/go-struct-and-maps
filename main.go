package main

import "fmt"

func main() {
	// New person
	adriano := person{
		firstName: "Adriano",
		lastName:  "Pulz",
	}

	// New contact
	contact := contactInfo{
		email:   "adrianopulz@gmail.com",
		zipCode: 88030000,
	}

	// Adding Phones
	contact.addPhone("Home", 5134456699)
	contact.addPhone("Mobile", 51988998800)

	// Add the contact to the person
	adriano.contact = contact

	adriano.print()
}

// ContacInfo struct custom type
type contactInfo struct {
	email   string
	zipCode int
	phones  map[string]int
}

// person struct custom type
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

// Add phone number on ContactInfo Map property
func (c *contactInfo) addPhone(phoneType string, phoneNumber int) {
	if c.phones == nil {
		c.phones = make(map[string]int)
	}

	c.phones[phoneType] = phoneNumber
}

// Print person data on the screen
func (p person) print() {
	fmt.Println("My name is:", p.firstName, p.lastName)
	fmt.Println("My contact ifos are:")
	fmt.Println("  Email:", p.contact.email)
	fmt.Println("  Zip Code:", p.contact.zipCode)

	contacts := p.contact.phones
	for phoneType, phonrNumber := range contacts {
		fmt.Println(" ", phoneType+" phone:", phonrNumber)
	}
}
