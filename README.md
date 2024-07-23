# golevel
GoLevel - веб-фреймворк на Go


Клонируйте репозиторий:

    ```sh
    git clone https://github.com/fugilove/golevel.git
    cd golevel
    ```
 Установите зависимости:

    ```sh
    go mod tidy
    ```


## Создание моделей и миграций

1. Change directory to your project:
   cd %s

2. Install dependencies:
   go mod tidy

3. Run the server:
   go run main.go

4. To generate a new migration, use:
   go run cli/cli.go make:migration <migration_name>

5. To generate a new model, use:
   go run cli/cli.go make:model <model_name>

6. To run migrations, use:
   go run cli/cli.go migrate
go run cli/cli.go help
Available commands:
  make:migration <name>       - Create a new migration file.
  make:model <name>           - Create a new model file.
  make:controller <name>      - Create a new controller file.
  make:seeder <name>          - Create a new seeder file.
  migrate                     - Run all migrations.
  migrate:rollback            - Rollback the last migration.
  seed                        - Run all seeders.
  serve                       - Start the application server.
  config:cache                - Cache the configuration.
  config:clear                - Clear the configuration cache.
  help                        - Show this help message.

Enjoy building your project!
