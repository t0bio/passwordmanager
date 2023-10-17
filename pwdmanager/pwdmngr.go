package main

import (
	"fmt"
	"os"
)

const pwd = "passwords.db"

func main() {
	// var pwds []string
	fmt.Println("Password Manager")
	fmt.Println("1. Add password")
	fmt.Println("2. Show saved passwords")
	fmt.Println("3. Exit")
	fmt.Print("Enter choice: ")
	var choice int
	fmt.Scanln(&choice)
	if choice == 3 {
		os.Exit(0)
	}
	for choice != 3 {
		switch choice {
		case 1:
			addPassword()
		case 2:
			showPassword()
		default:
			fmt.Println("Invalid choice")
		}
		fmt.Println("1. Add password")
		fmt.Println("2. Show saved passwords")
		fmt.Println("3. Exit")
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)
	}

}

func addPassword() {
	fmt.Println("Enter password: ")
	var password string
	var website string
	var username string
	fmt.Scanln(&password)
	fmt.Println("Enter website: ")
	fmt.Scanln(&website)
	fmt.Println("Enter username: ")
	fmt.Scanln(&username)
	f, err := os.OpenFile(pwd, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	if _, err = f.WriteString(website + " " + username + " " + password + "\n"); err != nil {
		fmt.Println(err)
	}
}

func showPassword() {
	fmt.Println("Enter website: ")
	var website string
	fmt.Scanln(&website)
	f, err := os.Open(pwd)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	var username string
	var password string
	for {
		_, err := fmt.Fscanln(f, &username, &password)
		if err != nil {
			break
		}
		if username == website {
			fmt.Println(password)
		}
	}
}
