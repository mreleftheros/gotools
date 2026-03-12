package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/mreleftheros/gotools/srv/pwhash"
)

type User struct {
	Username string
	Hash     string
}

func main() {
	users := make(map[string]User)
	var cmd string
	var username string
	var password string

	for {
		fmt.Print("Enter: ")
		_, err := fmt.Scanf("%s %s %s\n", &cmd, &username, &password)
		if err != nil {
			log.Fatal(err)
		}

		switch cmd {
		case "signup":
			fmt.Println("Signup...")
			_, ok := users[username]
			if ok {
				fmt.Println("user already exists")
				continue
			}

			h, err := pwhash.Hash(password)
			if err != nil {
				fmt.Println(err)
				continue
			}
			users[username] = User{username, h}
			fmt.Println(h)
			fmt.Println("Signup successful!")
		case "login":
			fmt.Println("Login...")
			v, ok := users[username]
			if !ok {
				fmt.Println("user does not exist")
				continue
			}

			matches, err := pwhash.Verify(v.Hash, password)
			if err != nil {
				fmt.Println("err", err)
				continue
			}
			if matches {
				fmt.Println("Login successful!")
			} else {
				fmt.Println("Login failed")
			}
		default:
			log.Fatal(errors.New("no such command"))
		}
	}
}
