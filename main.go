package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func availableCommands() {
	commands := []string{
		"create",
		"list",
		"update",
		"delete",
	}
	for _, c := range commands {
		fmt.Println(c)
	}

}

// first we need to have a defined task. right now it will be just a name
func createTask(task string) (bool, error) {
	//we need to actually create this
	file, err := os.OpenFile("./data/tasks.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//yes... i konw this looks ridculous. bare with me.
	t := fmt.Sprint(
		"name: ", task, ", ",
		" status: ", "todo, ",
		" createdAt: ", time.Now().UTC().Unix(), ", ",
		" updatedAt: ", time.Now().UTC().Unix())
	_, err = file.WriteString(t + "\n")
	if err != nil {
		panic(err)
	}
	return true, nil
}

func listTasks() {
	//now we want to be able to call this list of tasks and loop through reading each one.
	file, err := os.Open("./data/tasks.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func deleteTask(task string) (bool, error) {
	file, err := os.ReadFile("./data/tasks.txt")
	if err != nil {
		return false, err
	}
	lines := strings.Split(string(file), "\n")
	for i, line := range lines {
		t := strings.Split(line, ",")
		if strings.Contains(t[0], task) {
			lines[i] = ""
		}
		output := strings.Join(lines, "\n")
		err = os.WriteFile("./data/tasks.txt", []byte(output), 644)
		if err != nil {
			return false, nil
		}
	}
	return true, nil
}
func updateTask(task, status string) (bool, error) {
	file, err := os.ReadFile("./data/tasks.txt")
	if err != nil {
		return false, nil

	}
	lines := strings.Split(string(file), "\n")
	for i, line := range lines {
		t := strings.Split(line, ",")
		if strings.Contains(t[0], task) {
			nc := fmt.Sprint(
				"name: ", task, ", ",
				" status: ", status, ", ",
				" createdAt: ", t[3], ", ",
				" updatedAt: ", time.Now().UTC().Unix())
			lines[i] = nc
		}
		output := strings.Join(lines, "\n")
		err = os.WriteFile("./data/tasks.txt", []byte(output), 644)
		if err != nil {
			return false, nil
		}
	}
	return true, nil
}
func main() {
	/*Lets list out what we would want here in a small CLI
		- list tasks x
		- create task x
		- list commands x
		- update task x
		- delete task

	To do this with a cli we should create a list of these commands.
	*/
	if len(os.Args) < 2 {
		availableCommands()
		os.Exit(1)
	}
	flag := os.Args[1]
	switch flag {
	case "list":
		listTasks()
	case "create":
		task := os.Args[2]
		createTask(task)
	case "update":
		task := os.Args[2]
		status := os.Args[3]
		updateTask(task, status)
	case "delete":
		task := os.Args[2]
		deleteTask(task)
	default:
		availableCommands()
	}
}
