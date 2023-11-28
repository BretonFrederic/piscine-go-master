package main

import "github.com/01-edu/z01"

const (
	CLOSE = false
	OPEN  = true
)

type Door struct {
	state bool
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...\n")
	return ptrDoor.state
}

func OpenDoor(ptrDoor *Door) bool {
	PrintStr("Door Opening...\n")
	return ptrDoor.state
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?\n")
	if ptrDoor.state {
		ptrDoor.state = CLOSE
	}
	return ptrDoor.state
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?\n")
	if ptrDoor.state {
		ptrDoor.state = OPEN
	}
	return ptrDoor.state
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	if IsDoorOpen(door) {
		CloseDoor(door)
	}
	if door.state {
		CloseDoor(door)
	}
	if door.state {
		OpenDoor(door)
	}
	CloseDoor(door)
}
