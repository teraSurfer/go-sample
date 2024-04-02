package person

import (
	"errors"
	"time"
)

// Person with capital P means this struct can be accessed outside the package.
type Person struct {
	Name              string
	DayAndTimeOfBirth time.Time
}

// functional params, anyone can call this.
func New(name string, dateOfBirth string) (Person, error) {
	dob, err := time.Parse("2006-01-02", dateOfBirth)

	if err != nil {
		return Person{}, err
	}

	return Person{
		Name:              name,
		DayAndTimeOfBirth: dob,
	}, nil
}

// receiver, accessible only on person object.
func (p Person) AgeClassification() string {

	eighteenYearsAgo := time.Now().AddDate(-18, 0, 0)

	if p.DayAndTimeOfBirth.Before(eighteenYearsAgo) || p.DayAndTimeOfBirth.Equal(eighteenYearsAgo) {
		return "Adult"
	} else {
		return "Child"
	}
}

const electoralBody = "GOVERNMENT"

// persident with small letter means this struct is not accessible outside the package.
type president struct {
	// anonmymous field, this is composition where all properties of person are also available in president.
	// capital p means this can be accessed outside
	Person
	// this cannot be accessed from outside the package
	secrets []string
}

// this function is accessible from outside. returning errors is a good practice.
func NewPresident(elector string, name string, dateAndTimeOfBirth string) (*president, error) {

	p, err := New(name, dateAndTimeOfBirth)

	if err != nil {
		return nil, errors.New("president needs a valid date of birth.")
	}

	if electoralBody == elector {

		// return pointer not value.
		return &president{
			// Person: New("Biden", time.Date(1942, time.November, 20, 0, 0, 0, 0, time.UTC)),
			Person:  p,
			secrets: []string{"Nuclear codes", "Moon landing documents"},
		}, nil
	}

	return nil, errors.New("only valid electors are able to create new president")
}

// only president can access this method
func (p *president) dislike() string {
	return `
                      __________________
                      \                 \
                        \                 \
                          \                 \
                            \                 \
           /-------------------------------------
         //---------------//                  / |
       //               //                  / __|
     //               //                  / /  ||
   //               //                  / /    ||
 //_______________//   o o            / /      ||      ___/-\___
------------------------------------/   ------- |     |---------|
| DO NOT PLAY |         | HOUSEHOLD |           |      | | | | |
| ON OR AROUND|         |WASTE ONLY |           |      | | | | |
|--------------         ------------|           |      | | | | |
|                                   |           |      | | | | |
-------------------------------------------------      |_______|
`
}

func (p *president) like() string {
	return `
                                ████          
                                ██  ██        
                              ▓▓    ██        
                              ▓▓    ██        
                            ██    ██          
                            ██    ██          
                          ██      ██          
                          ██    ██            
  ██████████            ██      ██            
██░░▒▒▒▒░░▒▒██    ██████        ██████████████
██░░░░░░░░▒▒██  ██                          ██
██▒▒▒▒▒▒░░▒▒██  ██                          ██
██▓▓▓▓▓▓▓▓▓▓██  ██                ██████████  
██▓▓▓▓▒▒▓▓▓▓██  ██                        ██  
██▓▓▓▓▒▒▒▒▒▒██  ▓▓                ░░░░    ██  
██▓▓▓▓▓▓▓▓▓▓██  ██                ████████    
██▓▓▓▓▓▓▓▓▓▓██  ██                      ██    
██▓▓▓▓▓▓▓▓▓▓██  ██                      ██    
██▓▓▓▓▓▓▓▓▓▓██  ██                ██████      
██▓▓▓▓▓▓  ▓▓██  ██                    ██      
██▓▓▓▓▓▓▓▓▓▓██    ████████████████████        
  ██████████                             
	`
}

func (p *president) MakePolicyDecision(topic string) (string, error) {
	switch topic {
	case "Golang", "Go", "go", "golang":
		return p.like(), nil
	case "Java", "java":
		return p.dislike(), nil
	default:
		return "", errors.New("Not an acceptable topic for discussion.")
	}
}
