# A task management application (To-Do App)

## The following features:

### Key Features

#### 1. User Authentication

    1. Registration
    2. Login & Logout
    3. Route protection (only logged in users can access the dashboard)

#### 2. Task Managemen

    1. CRUD task (Create, Read, Update, Delete)
    2. Task status (unfinished / completed)
    3. Filter tasks by status

#### 3. Third Party API Integration

    1. Fetch daily motivational quotes from public API (eg: Add a little bit of body
       text) and display on the dashboard

#### 4. Dashboard

    1. Display a list of tasks
    2. Display motivational quotes
    3. Display information about logged in users

### Technology

#### Frontend:

* Language: Typescript
* Framework: React JS
* Tool: Vite
* UI: Tailwind css
* Http request: Axios
* [README](https://github.com/tri-hariyadi/test_upscale/blob/master/todo-app-client/README.md)

#### Backend:

* Language: Go
* Framework: native Go
* Database: PostgreSQL
* RESTful API
* Integration with public API [api-ninjas.com](https://api-ninjas.com/api/quotes) [GetQuote](https://github.com/tri-hariyadi/test_upscale/blob/master/todo-app-server/apps/todo/service/todo_service_impl.go)
* [README](https://github.com/tri-hariyadi/test_upscale/blob/master/todo-app-server/README.md)

#### Extra

* Deployed to the platform publik frontend is deployed to vercel end backend is deployed to railway
* Unit Test to Ensure Functions Work
* Using CI/CD Github Action to auto deploy

### How to Run

#### 1. Make sure `docker` and `docker-compose` is installed in your machine

#### 2. In the todo-app-server make sure to COPY `config.example.yaml` to `config.yaml` and in the client make sure API url is correct in `todo-app-client/src/lib/api/api.ts`

#### 3. You can change the variable value, but make sure the variable matches the one in the docker-compos.yml environment, especially in the postgresql database environment.

#### 4. Run client server application using docker-compose

```shell
docker-compose up -d --build
```

#### 5. If there is a problem with the application then it is necessary to run data migration

first install golang-migrate
```shell
go install -tags 'database1,database2' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

When installing Golang Migrate, an executable file named migrate is automatically created in the $GOPATH/bin/ folder.
This migrate file is the Golang Migrate application that we will use to create the Database Migration. Next run the
migration in path `common/database/migrations`

```shell
migrate -database "databasconnections" -path folder up
Example: migrate -database "postgres://root:secret@localhost:5432/todo_app?sslmode=disable" -path common/database/migrations
````

### If you want to run client and server applications separately, you can see the guide here:

#### client: [Run client application separately](https://github.com/tri-hariyadi/test_upscale/blob/master/todo-app-client/README.md)

#### server: [Run server application separately](https://github.com/tri-hariyadi/test_upscale/blob/master/todo-app-server/README.md)


