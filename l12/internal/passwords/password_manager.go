package passwords

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const (
	passFileName = "/assets/passwords.txt"
)

type PasswordManager struct {
	Passwords map[string]string `json:"Passwords"`
}

func NewManager() (*PasswordManager, error) {
	manager := &PasswordManager{}
	manager.Passwords = make(map[string]string)

	wd, _ := os.Getwd()
	file, err := os.ReadFile(wd + passFileName)
	if err != nil {
		return nil, errors.New("cannot parse password manager file")
	}
	json.Unmarshal(file, manager)
	return manager, nil
}

func (m *PasswordManager) Save() error {
	bytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	wd, _ := os.Getwd()
	err = os.WriteFile(wd+passFileName, bytes, 0644)
	if err != nil {
		return err
	}
	return nil
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
