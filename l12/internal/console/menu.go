package console

import (
	"encoding/json"
	"fmt"
	"main/internal/passwords"
	"os"
)

const (
	cmdExit = 0
	cmdList = 1
	cmdSave = 2
	cmdGet  = 3
)

func printMenu() {
	fmt.Println("1. List saved passwords")
	fmt.Println("2. Save password. Args: name password")
	fmt.Println("3. Get password")
	fmt.Println("0. Exit")
	fmt.Println("Please, enter an option: ")
}

func App() {
	manager, err := passwords.NewManager()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for {
		printMenu()
		var option int
		_, err := fmt.Scan(&option)
		if err != nil {
			continue
		}
		if option == cmdExit {
			fmt.Println("Bye")
			break
		}
		switch option {
		case cmdList:
			bytes, err := json.Marshal(manager)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(string(bytes))
		case cmdSave:
			var name, pass string
			args, err := fmt.Scan(&name, &pass)
			if args < 2 {
				fmt.Println(err)
				break
			}
			if err != nil {
				fmt.Println(err)
				break
			}
			manager.Add(name, pass)
			fmt.Println("Ok")
		case cmdGet:
			var name string
			_, err := fmt.Scan(&name)
			if err != nil {
				fmt.Println(err)
				break
			}
			pass, ok := manager.Get(name)
			if !ok {
				fmt.Println("Password not found")
				break
			}
			fmt.Println(pass)
		default:
			fmt.Println("Not a valid option")
		}
		fmt.Print("==========\n\n")
	}
	err = manager.Save()
	if err != nil {
		fmt.Println(err)
	}
}
