package main

import "testing"

func TestMain(t *testing.T) {
	var car carette
	path := readArrayOfIntFromFile("input.txt")
	if temp := car.countSteps(path); temp != 22570529 {
		t.Errorf("Expected %d", temp)
	}
}
