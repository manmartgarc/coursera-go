package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	fmt.Println("Enter a request or type 'exit' to quit:")
	fmt.Println("Available commands: newanimal, query, help")
	scanner := bufio.NewScanner(os.Stdin)
	animals := make(map[string]Animal)
	for fmt.Print("> "); scanner.Scan(); fmt.Print("> ") {
		input := scanner.Text()

		switch input {
		case "exit":
			return
		case "help":
			fmt.Printf("%s\n", strings.Repeat("-", 50))
			fmt.Println("Available commands are: newanimal, query, help")
			fmt.Println("\tnewanimal <name> <type>: create a new animal")
			fmt.Println("\t\ttypes: {cow, bird, snake}")
			fmt.Println("\tquery <name> <action>: query an animal")
			fmt.Println("\t\t actions: {eat, move, speak}")
			fmt.Printf("%s\n", strings.Repeat("-", 50))
		default:
			var command, name, actiontype string
			_, err := fmt.Sscanf(input, "%s %s %s", &command, &name, &actiontype)
			if err != nil {
				fmt.Println("Invalid input")
				continue
			}
			nAnimals := len(animals)

			switch command {
			case "newanimal":
				switch actiontype {
				case "cow":
					animals[name] = Cow{}
				case "bird":
					animals[name] = Bird{}
				case "snake":
					animals[name] = Snake{}
				default:
					fmt.Println("Invalid animal type")
				}
			case "query":
				if animal, ok := animals[name]; ok {
					value := reflect.ValueOf(animal)
					method := value.MethodByName(cases.Title(language.English, cases.Compact).String(actiontype))
					if !method.IsValid() {
						fmt.Println("Invalid action type")
						continue
					}
					method.Call(nil)
				} else {
					fmt.Println("Animal not found")
				}
			default:
				fmt.Println("Invalid command")
			}
			if len(animals) > nAnimals {
				fmt.Println("Created it!")
			}
		}
	}
}
