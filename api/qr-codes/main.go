package main

import "fmt"

// This microservice doesn't need to persist anything
// and should only create and return QR codes
// that should be stored in the invites microservice
func main() {
	fmt.Println("Hello qr codes!")

}
