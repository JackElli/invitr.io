package main

import "fmt"

// Need to make sure that database is set up correctly

// This microservice holds all of the info about invites
// it should be able to create, delete, edit invites
// add QR codes and retrieve info from the users microservice.

// The idea is that a user creates an invite to an event,
// which either:
// 1. they send to the recepient via email (either with QR code or direct to
//    calender)
// 2. they print out invites on a special card (for weddings, etc...)

// then the recipient can say whether they are going, add to calender,
// add a message etc..

// The event should hold info like date, place, who is organising etc..
// We could also add links to buy things like flowers, drinks, food
func main() {
	fmt.Println("Hello invites!")
}
