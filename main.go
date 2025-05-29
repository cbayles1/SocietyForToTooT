package main

import (
	"fmt"
)

type Thing struct {
	name      string
	isOnTopOf []Thing
}

var allThings = map[string]Thing{}

//var allThings []Thing

// func putAOnTopOfB(a *Thing, b Thing) {
// 	a.isOnTopOf = append(a.isOnTopOf, b)
// 	b.isBelow = append(b.isBelow, a)
// }

// func putAOnTopOfMultiple(a Thing, y []Thing) {
// 	for i := range y {
// 		a.isOnTopOf = append(a.isOnTopOf, y[i])
// 		y[i].isBelow = append(y[i].isBelow, a)
// 	}
// }

// func putMultipleOnTopOfB(x []Thing, b Thing) {
// 	for i := range x {
// 		x[i].isOnTopOf = append(x[i].isOnTopOf, b)
// 		b.isBelow = append(b.isBelow, x[i])
// 	}
// }

// func putMultipleOnTopOfMultiple(x []Thing, y []Thing) {
// 	for i := range x {
// 		for j := range y {
// 			x[i].isOnTopOf = append(x[i].isOnTopOf, y[j])
// 			y[j].isBelow = append(y[j].isBelow, x[i])
// 		}
// 	}
// }

func main() {
	fmt.Println("\"If there was not one thing that was not on top of another thing our society would be nothing more than a meaningless body of men that had gathered together for no good purpose.\"")
	fmt.Println(" - Sir William, President of the Royal Society for Putting Things on Top of Other Things")
	fmt.Println()

	var something Thing
	something.name = "something"
	allThings["something"] = something
	//allThings = append(allThings, something)

	for {
		thing := MakeNewThing()
		StackThing(thing)
		fmt.Println(len(allThings) - 1)
	}
}

func MakeNewThing() Thing {
	var newThing Thing
	fmt.Print("Enter a new thing: ")
	fmt.Scanln(&newThing.name)
	return newThing
}

func StackThing(thing Thing) {
	var onTopOf Thing

	if len(allThings) > 1 {
		fmt.Println("These are all the things which are currently on top of other things:")
		printAllThings()
		fmt.Printf("\nWhich shall %s recline on top of? ", thing.name)
		fmt.Scanln(&onTopOf.name)
	} else {
		onTopOf.name = "something"
	}

	found := false
	for name, lowerThing := range allThings { // linear search
		if name == onTopOf.name {
			thing.isOnTopOf = append(thing.isOnTopOf, lowerThing)
			found = true
			break
		}
	}

	if !found {
		fmt.Println("TRY AGAIN.")
		// TODO: Have user try again
	}

	allThings[thing.name] = thing
}

func printAllThings() {
	for name, thing := range allThings {

		if name == "something" { // don't print "something", it's the secret thing not on top of any other thing
			continue
		}

		fmt.Printf("%s is on top of ", name)

		if len(thing.isOnTopOf) <= 0 {
			fmt.Print("nothing. Shame! Shame!")
			//StackThing(thing)
			continue
		} else if len(thing.isOnTopOf) == 1 {
			fmt.Print(thing.isOnTopOf[0].name)
		} else {
			for i := range thing.isOnTopOf {
				fmt.Print(thing.isOnTopOf[i].name)

				if i < len(thing.isOnTopOf)-2 {
					fmt.Print(", ")
				} else if i < len(thing.isOnTopOf)-1 {
					fmt.Print(", and ")
				}
			}
		}
		fmt.Print(".")
	}
}
