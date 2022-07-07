package main

type contactInfo struct {
	Address string
	PhoneNo string
}

type Person struct {
	FirstName string
	LastName  string
	Contact   contactInfo
}

type Person1 struct {
	FirstName string
	LastName  string
	contactInfo
}

func main() {

}