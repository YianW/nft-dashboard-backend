package main

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/term"
	"os"
	"strings"
	"syscall"
	"tulip/backend/models"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter username (default 'tulipadmin'): ")
	username, err := reader.ReadString('\n')
	username = strings.Replace(username, "\n", "", -1)
	if username == "" {
		username = "tulipadmin"
	}
	fmt.Printf("::Using %q as username...\n", username)
	fmt.Println("Enter a password:")
	bytepw, err := term.ReadPassword(syscall.Stdin)
	if len(bytepw) < 4 {
		fmt.Println("Password cannot be empty")
		return
	}
	if err != nil {
		fmt.Printf("Error during password enter: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("Enter your password again:")
	bytepw2, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		fmt.Printf("Error during password enter: %s\n", err.Error())
		os.Exit(1)
	}
	if bytes.Equal(bytepw, bytepw2) {
		fmt.Println("Passwords do not match.")
		return
	}

	// Create user here
	models.ConnectDB()
	var u models.User
	u.Username = username
	u.Password = string(bytepw)
	u.Role = "superuser"
	u.Active = true
	_, err = u.SaveUser()
	if err != nil {
		fmt.Println("Error during save: ", err.Error())
	}
	fmt.Println("Successfully created user")
}
