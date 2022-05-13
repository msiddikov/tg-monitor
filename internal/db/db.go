package db

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

var (
	Monitors []Monitor
	Users    []User
)

func init() {
	_, err := os.Stat(os.Getenv("DB_PATH") + "monitors.json")
	if errors.Is(err, os.ErrNotExist) {
		Monitors = []Monitor{}
		saveMonitors()
	} else {
		loadMonitors()
	}
	_, err = os.Stat(os.Getenv("DB_PATH") + "users.json")
	if errors.Is(err, os.ErrNotExist) {
		Users = []User{}
		saveUsers()
	} else {
		loadUsers()
	}
}

//func loadFile(var interface{},)

func UpsertMonitor(m Monitor) {
	for k, v := range Monitors {
		if v.Id == m.Id {
			Monitors[k] = m
			saveMonitors()
			return
		}
	}
	if len(Monitors) > 0 {
		m.Id = Monitors[len(Monitors)-1].Id + 1
	} else {
		m.Id = 1
	}
	Monitors = append(Monitors, m)
	saveMonitors()
}

func UpsertUser(u User) {
	for k, v := range Users {
		if v.ChatId == u.ChatId {
			Users[k] = u
			saveUsers()
		}
	}
	Users = append(Users, u)
	saveUsers()
}

func saveMonitors() {
	data, _ := json.Marshal(Monitors)
	err := os.WriteFile(os.Getenv("DB_PATH")+"monitors.json", data, 0644)
	if err != nil {
		log.Panic("Unable to create file... Check permissions")
	}
}

func loadMonitors() {
	data, err := os.ReadFile(os.Getenv("DB_PATH") + "monitors.json")
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(data, &Monitors)
	if err != nil {
		log.Panic("broken monitors.json")
	}
}

func saveUsers() {
	data, _ := json.Marshal(Users)
	err := os.WriteFile(os.Getenv("DB_PATH")+"monitors.json", data, 0644)
	if err != nil {
		log.Panic("Unable to create file... Check permissions")
	}
}

func loadUsers() {
	data, err := os.ReadFile(os.Getenv("DB_PATH") + "users.json")
	if err != nil {
		log.Panic(err)
	}
	err = json.Unmarshal(data, &Monitors)
	if err != nil {
		log.Panic("broken users.json")
	}
}
