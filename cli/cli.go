package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run cli/cli.go <command> [args...]")
		return
	}

	command := os.Args[1]
	projectName := "golevel" // Замените это на реальное имя проекта или получите через аргумент командной строки

	switch command {
	case "make:migration":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run cli/cli.go make:migration <migration_name>")
			return
		}
		createMigration(projectName, os.Args[2])
	case "make:model":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run cli/cli.go make:model <model_name>")
			return
		}
		createModel(projectName, os.Args[2])
	case "make:controller":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run cli/cli.go make:controller <controller_name>")
			return
		}
		createController(projectName, os.Args[2])
	case "make:seeder":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run cli/cli.go make:seeder <seeder_name>")
			return
		}
		createSeeder(projectName, os.Args[2])
	case "migrate":
		runMigrations(projectName)
	case "migrate:rollback":
		rollbackMigration(projectName)
	case "seed":
		runSeeders(projectName)
	case "serve":
		startServer(projectName)
	case "config:cache":
		cacheConfig(projectName)
	case "config:clear":
		clearConfigCache(projectName)
	case "help":
		showHelp()
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}
}

func createMigration(projectName, name string) {
	migrationsDir := filepath.Join(projectName, "database", "migrations")

	if err := os.MkdirAll(migrationsDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating migrations directory %s: %v\n", migrationsDir, err)
		return
	}

	timestamp := time.Now().Format("20060102150405")
	fileName := fmt.Sprintf("%s_%s.go", timestamp, strings.ToLower(name))
	filePath := filepath.Join(migrationsDir, fileName)

	content := fmt.Sprintf(`package migrations

import (
	"database/sql"
	"log"
)

func Up(db *sql.DB) error {
	// Write your migration logic here
	log.Printf("Running migration: %s", %q)
	return nil
}

func Down(db *sql.DB) error {
	// Write your rollback logic here
	log.Printf("Rolling back migration: %s", %q)
	return nil
}
`, name, name)

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error creating migration file %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("Migration %s created successfully!\n", fileName)
}

func createModel(projectName, name string) {
	modelsDir := filepath.Join(projectName, "app", "models")
	fileName := fmt.Sprintf("%s.go", strings.ToLower(name))
	filePath := filepath.Join(modelsDir, fileName)

	if err := os.MkdirAll(modelsDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating models directory %s: %v\n", modelsDir, err)
		return
	}

	content := fmt.Sprintf(`package models

import (
	"time"
)

type %s struct {
	ID        int64     `+"`json:\"id\"`"+`
	Username  string    `+"`json:\"username\"`"+`
	Email     string    `+"`json:\"email\"`"+`
	Password  string    `+"`json:\"-\"`"+`
	CreatedAt time.Time `+"`json:\"created_at\"`"+`
}
`, strings.Title(name))

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error creating model file %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("Model %s created successfully!\n", name)
}

func createController(projectName, name string) {
	controllersDir := filepath.Join(projectName, "app", "controllers")
	fileName := fmt.Sprintf("%s_controller.go", strings.ToLower(name))
	filePath := filepath.Join(controllersDir, fileName)

	if err := os.MkdirAll(controllersDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating controllers directory %s: %v\n", controllersDir, err)
		return
	}

	content := fmt.Sprintf(`package controllers

import (
	"net/http"
	"html/template"
)

type %sController struct{}

func (c %sController) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("resources/views/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
`, name, name)

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error creating controller file %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("Controller %s created successfully!\n", name)
}

func createSeeder(projectName, name string) {
	seederDir := filepath.Join(projectName, "database", "seeders")
	fileName := fmt.Sprintf("%s_seeder.go", strings.ToLower(name))
	filePath := filepath.Join(seederDir, fileName)

	if err := os.MkdirAll(seederDir, os.ModePerm); err != nil {
		fmt.Printf("Error creating seeders directory %s: %v\n", seederDir, err)
		return
	}

	content := fmt.Sprintf(`package seeders

import (
	"database/sql"
	"log"
)

func Seed(db *sql.DB) {
	// Write your seeder logic here
	log.Printf("Running seeder: %s", %q)
}
`, name, name)

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error creating seeder file %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("Seeder %s created successfully!\n", name)
}

func runMigrations(projectName string) {
	fmt.Println("Running migrations for project:", projectName)
	// Реализуйте логику выполнения миграций здесь
}

func rollbackMigration(projectName string) {
	fmt.Println("Rolling back last migration for project:", projectName)
	// Реализуйте логику отката миграций здесь
}

func runSeeders(projectName string) {
	fmt.Println("Running seeders for project:", projectName)
	// Реализуйте логику выполнения сидеров здесь
}

func startServer(projectName string) {
	fmt.Println("Starting server for project:", projectName)
	// Реализуйте логику запуска сервера здесь
}

func cacheConfig(projectName string) {
	fmt.Println("Caching configuration for project:", projectName)
	// Реализуйте логику кэширования конфигурации```go
}

func clearConfigCache(projectName string) {
	fmt.Println("Clearing configuration cache for project:", projectName)
	// Реализуйте логику очистки кэша конфигурации здесь
}

func showHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  make:migration <name>       - Create a new migration file.")
	fmt.Println("  make:model <name>           - Create a new model file.")
	fmt.Println("  make:controller <name>      - Create a new controller file.")
	fmt.Println("  make:seeder <name>          - Create a new seeder file.")
	fmt.Println("  migrate                     - Run all migrations.")
	fmt.Println("  migrate:rollback            - Rollback the last migration.")
	fmt.Println("  seed                        - Run all seeders.")
	fmt.Println("  serve                       - Start the application server.")
	fmt.Println("  config:cache                - Cache the configuration.")
	fmt.Println("  config:clear                - Clear the configuration cache.")
	fmt.Println("  help                        - Show this help message.")
}
