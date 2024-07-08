package main

import "fmt"

var ESPRESSO = [4]int{250, 0, 16, 4}
var LATTE = [4]int{350, 75, 20, 7}
var CAPPUCCINO = [4]int{200, 100, 12, 6}

func main() {
	water := new(int)
	*water = 400
	milk := new(int)
	*milk = 540
	beans := new(int)
	*beans = 120
	cups := new(int)
	*cups = 9
	money := new(int)
	*money = 550

	var action string
	for action != "exit" {
		action = askUser()
		switch action {
		case "buy":
			buy(water, milk, beans, cups, money)
		case "fill":
			fill(water, milk, beans, cups)
		case "take":
			take(money)
		case "remaining":
			printCurrentState(water, milk, beans, cups, money)
		case "exit":
			continue
		}
		fmt.Println("")
	}
}

func askUser() string {
	var action string
	fmt.Println("Write action (buy, fill, take, remaining, exit):")
	fmt.Scanln(&action)
	return action
}

func take(money *int) {
	fmt.Println("I gave you $", *money)
	*money = 0
}

func fill(water *int, milk *int, beans *int, cups *int) {
	var addWater int
	var addMilk int
	var addBeans int
	var addCups int
	fmt.Println("Write how many ml of water do you want to add:")
	fmt.Scanln(&addWater)
	fmt.Println("Write how many ml of milk do you want to add:")
	fmt.Scanln(&addMilk)
	fmt.Println("Write how many grams of coffee beans do you want to add:")
	fmt.Scanln(&addBeans)
	fmt.Println("Write how many disposable cups of coffee do you want to add:")
	fmt.Scanln(&addCups)

	*water += addWater
	*milk += addMilk
	*beans += addBeans
	*cups += addCups
}

func printCurrentState(water *int, milk *int, beans *int, cups *int, money *int) {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d of water\n", *water)
	fmt.Printf("%d of milk\n", *milk)
	fmt.Printf("%d of coffee beans\n", *beans)
	fmt.Printf("%d of disposable cups\n", *cups)
	fmt.Printf("%d of money\n", *money)
}

func buy(water *int, milk *int, beans *int, cups *int, money *int) {
	var coffeeType string
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	fmt.Scanln(&coffeeType)

	var coffee [4]int
	switch coffeeType {
	case "1":
		coffee = ESPRESSO
	case "2":
		coffee = LATTE
	case "3":
		coffee = CAPPUCCINO
	case "back":
		return
	default:
		coffee = [4]int{0, 0, 0, 0}
	}

	if *water < coffee[0] {
		fmt.Println("Sorry, not enough water!")
	} else if *milk < coffee[1] {
		fmt.Println()
	} else if *beans < coffee[2] {
		fmt.Println("Sorry, not enough coffee beans!")
	} else if *cups < 1 {
		fmt.Println("Sorry, not enough disposable cups!")
	} else {
		fmt.Println("I have enough resources, making you a coffee!")
		*water -= coffee[0]
		*milk -= coffee[1]
		*beans -= coffee[2]
		*cups--
		*money += coffee[3]
	}

}
