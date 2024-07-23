package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: golevel new <project_name>")
		return
	}

	command := os.Args[1]
	projectName := os.Args[2]

	if command == "new" && projectName != "" {
		err := createProject(projectName)
		if err != nil {
			fmt.Println("Error creating project:", err)
		}
	} else if command == "install" {
		err := install()
		if err != nil {
			fmt.Println("Error installing GoLevel:", err)
		}
	} else {
		fmt.Println("Invalid command or project name")
	}
}

func createProject(projectName string) error {
	fmt.Println("Creating project:", projectName)

	// Clone the GoLevel repository
	cmd := exec.Command("git", "clone", "https://github.com/yourusername/golevel", projectName)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error cloning repository: %w", err)
	}

	// Change directory to the project
	err = os.Chdir(projectName)
	if err != nil {
		return fmt.Errorf("error changing directory: %w", err)
	}

	// Initialize Go module
	cmd = exec.Command("go", "mod", "init", projectName)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error initializing Go module: %w", err)
	}

	// Install dependencies
	cmd = exec.Command("go", "mod", "tidy")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error installing dependencies: %w", err)
	}

	fmt.Println("Project", projectName, "created successfully!")
	return nil
}

func install() error {
	if runtime.GOOS != "windows" {
		return fmt.Errorf("install script only supports Windows")
	}

	fmt.Println("Installing GoLevel...")

	// Compile the binary
	cmd := exec.Command("go", "build", "-o", "golevel", "install.go")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error building binary: %w", err)
	}

	// Get the path of the current executable
	executablePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error getting executable path: %w", err)
	}
	dirPath := filepath.Dir(executablePath)

	// Determine the PATH environment variable
	pathEnv := os.Getenv("PATH")
	if !strings.Contains(pathEnv, dirPath) {
		fmt.Println("Adding GoLevel to PATH...")

		// Add the directory to PATH
		newPath := fmt.Sprintf("%s;%s", dirPath, pathEnv)
		err = os.Setenv("PATH", newPath)
		if err != nil {
			return fmt.Errorf("error setting PATH environment variable: %w", err)
		}
	}

	fmt.Println("GoLevel installed successfully!")
	return nil
}
