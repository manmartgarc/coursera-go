package main

import (
	"bufio"
	"fmt"
	"os"
)

type Animal struct {
	food, locomotion, noise string
}

func (a *Animal) Eat() string {
	return a.food
}

func (a *Animal) Move() string {
	return a.locomotion
}

func (a *Animal) Speak() string {
	return a.noise
}

func NewAnimal(food, locomotion, noise string) *Animal {
	return &Animal{food: food, locomotion: locomotion, noise: noise}
}

func main() {
	animals := map[string]*Animal{
		"cow":   NewAnimal("grass", "walk", "moo"),
		"bird":  NewAnimal("worms", "fly", "peep"),
		"snake": NewAnimal("mice", "slither", "hsss"),
	}

	fmt.Println("Available animals: cow, bird, snake")
	fmt.Println("Available actions: eat, move, speak")

	for {
		var animal, action string
		fmt.Println("\nEnter animal and action (or type 'exit' to quit): ")
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}
		fmt.Sscanf(input, "%s %s", &animal, &action)

		if _, ok := animals[animal]; !ok {
			println("Invalid animal")
			continue
		}

		switch action {
		case "eat":
			println(animals[animal].Eat())
		case "move":
			println(animals[animal].Move())
		case "speak":
			println(animals[animal].Speak())
		default:
			println("Invalid action")
		}
	}
}
