package base

import "fmt"

type alertCounter int
type AlertCounter int
type Counter struct {
	Name  string
	Value int
	info
}

type info struct {
	Date *Time
}

func New(value int) alertCounter {
	return alertCounter(value)
}
func init() {
	fmt.Printf("\033[1;33;40m%s\033[0m\n", "counters")
}
