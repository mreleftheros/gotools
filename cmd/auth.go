package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/mreleftheros/gotools/srv/password"
)

const salt = "ff1nQkQQBx+At0yPAjDeMQ"

type User struct {
	Name string
	Hash []byte
}

func main() {
	users := make(map[string]User)
	var cmd string
	var name string
	var pw string

	for {
		fmt.Print("Enter: ")
		_, err := fmt.Scanf("%s %s %s\n", &cmd, &name, &pw)
		if err != nil {
			log.Fatal(err)
		}

		switch cmd {
		case "signup":
			fmt.Println("Signup")
			_, ok := users[name]
			if ok {
				fmt.Println("user already exists")
				continue
			}

			h := password.Hash(pw, salt)
			users[name] = User{name, h}
			fmt.Println("Signup successful")
		case "login":
			fmt.Println("Login")
			v, ok := users[name]
			if !ok {
				fmt.Println("user does not exist")
				continue
			}

			ok = password.Compare(pw, salt, v.Hash)
			if ok {
				fmt.Println("Login successful")
			} else {
				fmt.Println("Login failed")
			}
		default:
			log.Fatal(errors.New("no such command"))
		}
	}
}
