package main

import "fmt"

func main() {
	fmt.Println("Hello qr codes!")

	// This microservice doesn't need to persist anything
	// and should only create and return QR codes
	// that should be stored in the invites microservice
}
