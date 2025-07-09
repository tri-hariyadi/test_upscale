# Todo App Server

### Technology used:

* Language: Go
* Framework: native Go
* Database: PostgreSQL

### Project Structure:

```
├── apps                                    # Used to store features of application for example management todo, user etc.
│   ├── todo                                # one of the features in the application
│   │   ├── handler                         # Presentation layer (HTTP handler / controller)
│   │   │   ├── todo_handler.go             # Interface for handler.
│   │   │   └── todo_handler_impl.go        # handler implementation, for example HTTP handler.
│   │   │
│   │   ├── repository                      # Data access layer (DB interaction)
│   │   │   ├── todo_repository.go          # Repo interface for todo.
│   │   │   └── todo_repository_impl.go     # Implementation of DB operations (Create, Update, etc.).
│   │   │
│   │   └── service                         # Business layer / domain logic
│   │       ├── todo_service.go             # Interface service.
│   │       └── todo_service_impl.go        # Logic such as validation, data processing, etc.
│   │    
│   ├── ...
│   └── ... other features
│            
├── common                                  # **Contains global components for the entire application such as configuration, database, router, and server.**
│   ├── config/config.go                    # Parsing dan akses konfigurasi dari config.yaml.
│   ├── database
│   │   ├── migrations                      # Save the SQL migration file (using tools golang-migrate)
│   │   └── database.go                     # Setup DB coonection (PostgreSQL)
│   │
│   ├── injector
│   │   ├── http_server.go                  # Dependency injection for HTTP handlers and services.
│   │   └── wire_gen.go
│   │
│   ├── logger/logger.go                    # Setup logger global using logrus.
│   ├── router/router.go                    # Setup all application routing (register handlers to paths).
│   └── server
│       ├── http_server.go                  # Configure and start HTTP server.
│       └── shutdown.go                     # Handling graceful server shutdown.
│
├── helpers    
│   ├── auth                                # Authentication and token related utilities.
│   ├── exceptions                          # Custom error handling and error wrapping.
│   └── resp_handler                        # Utility to respond to requests consistently (standard response format).
│
├── middleware                              # Contains HTTP middleware used on routers.
│   ├── cors.go                             # Handles CORS configuration.
│   ├── logging.go                          # Logging every incoming request and outgoing response.
│   ├── recovery.go                         # Panic handling to prevent server crashes.
│   └── token_verify.go                     # Verify JWT or other authorization tokens.
│
├── model                                   # **Contains data structure definitions for entities and the web.**
│   ├── web                                 # 
│   │   ├── response_wrapper.go             # DTO for todo responses.
│   │   ├── todo_request_response.go        # DTO for todo request and response.
│   │   └── user_request_response.go        # DTO for users request and response.
│   │
│   ├── todo.go                             # Model for todo
│   └── user.go                             # Model for user
│
├── config.yaml                             # The main configuration file is like: database name, database port, app port, api secret etc.
├── go.mod                                  # Contains information about Go project modules and dependencies.
├── go.sum                                  # Contains information about Go project modules and dependencies.
└── main.go                                 # Application entry point. Here is done: Load config, Setup logger, Init DB, Inject dependencies, start server
```

### How to Run

1. Make sure go version 1.24.1 is installed in your machine
2. Run go version to make sure go is installed
3. Run `go mod tidy`
4. Run `migrate -database "databasconnections" -path folder up`
    1. for
       example `migrate -database "postgres://root:secret@localhost:5432/todo_app?sslmode=disable" -path common/database/migrations up`
5. For auto generate Dependency Injection code you need to install google wire
    1. Run `go install github.com/google/wire/cmd/wire@latest` Automatically there will be a binary file in
       `$GOPATH/bin/wire`
    2. So that the wire command line application can be accessed, don't forget to put it in the `$PATH` of your operating
       system.
6. Run `cd common/injector` then Run `wire`, if success file containing the dependencies required by the application will be generated in the `wire_gen.go` file, if there is no file `wire_gen.go` my be google wire is not installed correctly
7. Run `go run main.go` and server will be started

If you use `docker` and `docker-compose`, you can run a single package project and its server application with a single command without the hassle of running the client and server applications individually. Follow the guide here.
