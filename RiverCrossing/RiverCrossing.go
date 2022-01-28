package RiverCrossing

import "fmt"

var state = "Nothing is in the boat"

func ViewState() string {
	return state
}

func PutInFox() {
	state = "[kylling korn hs ---\\rev/ _/---]"
}

func PutInWhat(what string) {
	fmt.Printf("\\___%s___/", what)
}

func PutInHS() {
	state = "[kylling korn rev ---\\hs/ _/---]"
}

func PutInCorn() {
	state = "[kylling rev hs ---\\korn/ _/---]"
}

func PutInChicken() {
	state = "[rev korn hs ---\\kylling/ _/---]"
}
