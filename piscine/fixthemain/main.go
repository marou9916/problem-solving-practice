package main

import "github.com/01-edu/z01"

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state == "OPEN" {
		CloseDoor(door)
	}
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

type Door struct {
	state string
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = "CLOSE"
	return true
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...")
	ptrDoor.state = "OPEN"
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	if ptrDoor.state == "OPEN" {
		return true
	} else {
		return false
	}
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	if ptrDoor.state == "CLOSE" {
		return true
	} else {
		return false
	}
}
