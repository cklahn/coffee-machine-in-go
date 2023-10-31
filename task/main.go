package main

import "fmt"

func remaining(water, milk, beans, cocoa, cups, money int) {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%v ml of water\n", water)
	fmt.Printf("%v ml of milk\n", milk)
	fmt.Printf("%v g of coffee beans\n", beans)
	fmt.Printf("%v g of cacoa\n", cocoa)
	fmt.Printf("%v disposable cups\n", cups)
	fmt.Printf("$%v of money\n", money)
	fmt.Println()
}

func outOfRessources(missing string) {
	fmt.Printf("Sorry, not enough %s!\n", missing)
	fmt.Println()
}

func updateState(water, milk, beans, cacoa, cups, money *int, step string, selection int) {
	if step == "buy" {
		switch selection {
		case 1:
			*water -= 250
			*beans -= 16
			*cups -= 1
			*money += 4
		case 2:
			*water -= 350
			*milk -= 75
			*beans -= 20
			*cups -= 1
			*money += 7
		case 3:
			*water -= 200
			*milk -= 100
			*beans -= 12
			*cups -= 1
			*money += 6
		case 4:
			*milk -= 200
			*cacoa -= 25
			*cups -= 1
			*money += 8
		}

	}
}

func buy(water, milk, beans, cacoa, cups, money *int) {
	var coffee int
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, 4 - cacoa:")
	fmt.Scan(&coffee)
	fmt.Println()

	switch coffee {
	case 1:
		if *water >= 250 && *beans >= 16 && *cups >= 1 {
			updateState(water, milk, beans, cacoa, cups, money, "buy", coffee)
			fmt.Println("I have enough resources, making you a coffee!")
			fmt.Println()
		} else if *water < 250 {
			outOfRessources("water")
		} else if *beans < 16 {
			outOfRessources("beans")
		} else if *cups < 1 {
			outOfRessources("cups")
		}
	case 2:
		if *water >= 350 && *milk >= 75 && *beans >= 20 && *cups >= 1 {
			updateState(water, milk, beans, cacoa, cups, money, "buy", coffee)
			fmt.Println("I have enough resources, making you a coffee!")
			fmt.Println()
		} else if *water < 350 {
			outOfRessources("water")
		} else if *milk < 75 {
			outOfRessources("milk")
		} else if *beans < 20 {
			outOfRessources("beans")
		} else if *cups < 1 {
			outOfRessources("cups")
		}
	case 3:
		if *water >= 200 && *milk >= 100 && *beans >= 12 && *cups >= 1 {
			updateState(water, milk, beans, cacoa, cups, money, "buy", coffee)
			fmt.Println("I have enough resources, making you a coffee!")
			fmt.Println()
		} else if *water < 200 {
			outOfRessources("water")
		} else if *milk < 100 {
			outOfRessources("milk")
		} else if *beans < 12 {
			outOfRessources("beans")
		} else if *cups < 1 {
			outOfRessources("cups")
		}
	case 4:
		if *milk >= 200 && *cacoa >= 25 && *cups >= 1 {
			updateState(water, milk, beans, cacoa, cups, money, "buy", coffee)
			fmt.Println("I have enough resources, making you a cacoa!")
			fmt.Println()
		} else if *milk < 100 {
			outOfRessources("milk")
		} else if *cacoa < 12 {
			outOfRessources("cacoa")
		} else if *cups < 1 {
			outOfRessources("cups")
		}
	default:
		fmt.Println("not possible")
	}
}

func fill(exWater, exMilk, exBeans, exCacoa, exCups *int) {
	var water, milk, beans, cacoa, cups int
	fmt.Println("Write how many ml of water you want to add:")
	fmt.Scan(&water)
	fmt.Println("Write how many ml of milk you want to add:")
	fmt.Scan(&milk)
	fmt.Println("Write how many grams of coffee beans you want to add:")
	fmt.Scan(&beans)
	fmt.Println("Write how many grams of cacoa you want to add:")
	fmt.Scan(&cacoa)
	fmt.Println("Write how many disposable cups you want to add:")
	fmt.Scan(&cups)
	fmt.Println()

	*exWater += water
	*exMilk += milk
	*exBeans += beans
	*exCacoa += cacoa
	*exCups += cups
}

func take(money *int) {
	fmt.Printf("I gave you $%v\n", *money)
	fmt.Println()
	*money = 0
}

func main() {
	// write your code here
	var action string
	on := true
	water := 400
	milk := 540
	beans := 120
	cacoa := 100
	cups := 9
	money := 550

	for on {
		fmt.Println("Write action (buy, fill, take, remaining, exit): ")
		fmt.Scan(&action)
		fmt.Println()

		switch action {
		case "remaining":
			remaining(water, milk, beans, cacoa, cups, money)
		case "fill":
			fill(&water, &milk, &beans, &cacoa, &cups)
		case "take":
			take(&money)
		case "buy":
			buy(&water, &milk, &beans, &cacoa, &cups, &money)
		case "exit":
			on = false
		default:
			fmt.Println("Please use buy, fill, take, remaining or exit")
		}

	}
}
