package main

import (
	"fmt"
	"log"

	"go-sample/person"
)

func main() {
	var favLang string

	john, err := person.New("John Doe", "1994-01-01")

	fmt.Println(john.Greeter(), person.HandleChildren(john))
	// err != nil is a common check, and its okay for it to be  repeated multiple times.
	// dont try to over engineer this. go is all about simplicity.
	if err != nil {
		log.Fatalln("this date should not error!", err)
	}

	persident, err := person.PresidentGenerator("Not a bad person")

	if err != nil {
		log.Fatalln("this date should not error!", err)
	}

	fmt.Println(persident.MakePolicyDecision("go"))

	fmt.Printf("%s's date of birth is %s and they are an %s\n", john.Name, john.DayAndTimeOfBirth.String(), john.AgeClassification())

	biden, err := person.NewPresident("GOVERNMENT", "Joe Biden", "1942-11-20")

	if err != nil {
		log.Fatalln("this date should not error!", err)
	}

	fmt.Printf("%s's date of birth is %s and they are an %s\n", biden.Name, biden.DayAndTimeOfBirth.String(), biden.AgeClassification())

	fmt.Println("Enter your favorite programming language")
	_, err = fmt.Scanln(&favLang)

	if err != nil {
		log.Fatalln("Sorry could not get your favorite language.")
	}

	fmt.Printf("%s's thoughts are\n", biden.Name)
	thoughts, err := biden.MakePolicyDecision(favLang)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(thoughts)

}
