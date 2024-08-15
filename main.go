package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Name string `json:"name"`
	//TODO: make this a more defined type
	Status    string `json:"status"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

func getUnixTime() int64 {
	return time.Now().UTC().Unix()
}

// Create the task and store it in a file
func createTask(task string) (Task, error) {
	//forgot to turn this into a file.
	fileName := fmt.Sprint(task, "_", getUnixTime(), ".json")
	file, err := os.Create("./data/" + fileName)
	if err != nil {
		return Task{}, err
	}
	defer file.Close()
	content := Task{
		Name:      task,
		Status:    "pending",
		CreatedAt: getUnixTime(),
		UpdatedAt: getUnixTime(),
	}
	jsonData, err := json.Marshal(content)
	if err != nil {
		return Task{}, err
	}
	str := string(jsonData)
	_, err = file.WriteString(str)
	if err != nil {
		return Task{}, err
	}
	return Task{}, nil
}

func listCommands() []string {
	commands := []string{
		"create",
	}
	return commands
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(listCommands())
		os.Exit(1)
	}
	command := os.Args[1]
	switch command {
	case "create":
		task := os.Args[2]
		createTask(task)
	}
}
