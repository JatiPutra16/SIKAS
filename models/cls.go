package models

import (
	"fmt"
)

func Clearscreen() {
	fmt.Print("\033[H\033[2J")
}
