package main

import "github.com/linhptr/rivercrossing/state"

func main() {
    state.PutInBoat("kylling")
    state.CrossRiver("kylling")
    state.PutInBoat("rev")
    state.CrossRiver("rev")
    state.PutInBoat("korn")
    state.CrossRiver("korn")

}