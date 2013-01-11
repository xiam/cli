package cli

import (
	"flag"
	"fmt"
	"testing"
)

/* Flags */
var age = flag.Uint("age", 0, "Your age.")
var city = flag.String("city", "", "Your city of residence.")

/* Command definition */
type command1 struct {
}

func (self *command1) Help() error {
	fmt.Println("This command shows help topics.")
	return nil
}

func (self *command1) Usage() error {
	fmt.Println("This command requires no arguments.")
	return nil
}

func (self *command1) Execute() error {

	if *age == 0 {
		fmt.Printf("You didn't tell me your age.\n")
	} else {
		fmt.Printf("I see, you're %d years old.\n", *age)
	}

	if *city == "" {
		fmt.Printf("You didn't tell me the name of your city.\n")
	} else {
		fmt.Printf("So, you live in %s\n", *city)
	}

	return nil
}

func TestRegister(t *testing.T) {
	err := Register("command1", &command1{})
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
}

func TestUsage(t *testing.T) {
	err := Usage("command1")
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
}

func TestExecute(t *testing.T) {
	err := Execute("command1")
	if err != nil {
		t.Errorf("Error: %v\n", err)
	}
}
