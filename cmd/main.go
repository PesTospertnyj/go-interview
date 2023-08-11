package main

import (
	"fmt"
)

func main() {
	// var a, b = []string{"Helga", "Jane", "Max", "Jane"}, []string{"Eugene", "John", "Iov", "John"}
	//
	// res := internal.UniqueNames(a, b)
	//
	// fmt.Println(res)
	//
	// newBank := bank.NewBankAccount()
	//
	// err := newBank.Deposit(100)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println(GetFullName("John", "Doe"))
}

func GetName(name string) string {
	return name
}

func GetFullName(firstName, lastName string) func(string) string {
	return func(name string) string {
		return fmt.Sprintf("%s - %s", firstName, lastName)
	}
}
