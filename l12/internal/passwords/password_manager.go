package passwords

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	passFileName = "/assets/passwords.txt"
)

type PasswordManager struct {
	Passwords map[string]string `json:"Passwords"`
}

func NewManager() *PasswordManager {
	manager := &PasswordManager{}
	manager.Passwords = make(map[string]string)

	wd, _ := os.Getwd()
	file, err := os.ReadFile(wd + passFileName)
	if err != nil {
		fmt.Println(err.Error())
		return manager
	}
	json.Unmarshal(file, manager)
	return manager
}

func (m *PasswordManager) Save() {
	bytes, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	wd, _ := os.Getwd()
	err = os.WriteFile(wd+passFileName, bytes, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func (m *PasswordManager) Add(name, pass string) {
	_, ok := m.Passwords[name]
	if ok {
		fmt.Println("Password already exist. Replacing.")
	}
	m.Passwords[name] = pass
}

func (m *PasswordManager) Get(name string) (string, bool) {
	pass, ok := m.Passwords[name]
	return pass, ok
}
