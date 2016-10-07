package main

import "fmt"

func main() {
	events := getEvents()
	fmt.Printf("There are %v events!", len(events))
}

func getEvents() []Event {
	return []Event{
		{Name: "ToulouseJS", Location: "Etincelle Coworking"},
		{Name: "AngularToulouse", Location: "At Home"},
	}
}

type Event struct {
	Name     string
	Location string
}
