package console

import (
	"encoding/json"
	"fmt"
	"main/internal/passwords"
)

func printMenu() {
	fmt.Println("1. List saved passwords")
	fmt.Println("2. Save password. Args: name password")
	fmt.Println("3. Get password")
	fmt.Println("0. Exit")
	fmt.Println("Please, enter an option: ")
}

func App() {
	manager := passwords.NewManager()
	for {
		printMenu()
		var option int
		_, err := fmt.Scan(&option)
		if err != nil {
			continue
		}
		if option == 0 {
			fmt.Println("Bye")
			break
		}
		switch option {
		// list passwords
		case 1:
			bytes, err := json.Marshal(manager)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println(string(bytes))
		// save password
		case 2:
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
		// get password
		case 3:
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
	manager.Save()
}
