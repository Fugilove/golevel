package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: golevel new <project_name>")
		return
	}

	command := os.Args[1]
	projectName := os.Args[2]

	if command == "new" && projectName != "" {
		createProject(projectName)
	} else {
		fmt.Println("Invalid command or project name")
	}
}

func createProject(projectName string) {
	fmt.Println("Creating project:", projectName)

	// Clone the GoLevel repository
	cmd := exec.Command("git", "clone", "https://github.com/Fugilove/golevel", projectName)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error cloning repository:", err)
		return
	}

	// Change directory to the project
	err = os.Chdir(projectName)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	// Initialize Go module
	cmd = exec.Command("go", "mod", "init", projectName)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error initializing Go module:", err)
		return
	}

	// Install dependencies
	cmd = exec.Command("go", "mod", "tidy")
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error installing dependencies:", err)
		return
	}

	fmt.Println("Project", projectName, "created successfully!")
}
