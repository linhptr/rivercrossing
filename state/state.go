package state

import "fmt"

var States = [7]string{
    "[V---\\___kylling rev korn HS___\\ \\______/______/--Ø]",
    "[V---\\___rev korn___\\ \\___kylling+HS___/______/--Ø]",
    "[V---\\___rev korn___\\ \\___HS___/____kylling__/--Ø]",
    "[V---\\___korn___\\ \\___rev+HS___/____kylling__/--Ø]",
    "[V---\\___korn___\\ \\___HS___/____rev kylling__/--Ø]",
    "[V---\\______\\ \\___korn+HS___/____rev kylling__/--Ø]",
    "[V---\\______\\ \\______/____HS korn rev kylling__/--Ø]"}

var Status = States[0]

func PutInBoat(item string) string {
    if item == "kylling" {
        Status = States[1]
        fmt.Println(States[1])
        return States[1]
    }

    if item == "rev" {
        Status = States[3]
        fmt.Println(States[3])
        return States[3]
    }

    if item == "korn" {
        Status = States[5]
        fmt.Println(States[5])
    }
        return States[5]
}

func CrossRiver(item string) string{
    if item == "kylling" {
        Status = States[2]
        fmt.Println(States[2])
        return States[2]
    }
    if item == "rev" {
        Status = States[4]
        fmt.Println(States[4])
        return States[4]
    }
    if item == "korn" {
            Status = States[6]
            fmt.Println(States[6])
        }
        return States[6]
}