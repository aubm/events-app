package main

import "testing"

func TestGetEvents(t *testing.T) {
	events := getEvents()
	nb := len(events)

	if nb != 2 {
		t.Errorf("expected nb of events to be 2, got: %v", nb)
	}
}
