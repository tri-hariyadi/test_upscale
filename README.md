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

#### Backend:

* Language: Go
* Framework: native Go
* Database: PostgreSQL
* RESTful API

### How to Run

#### 1. Make sure `docker` and `docker-compose` is installed in your machine

#### 2. In the todo-app-server make sure to COPY `config.example.yaml` to `config.yaml`

#### 3. You can change the variable value, but make sure the variable matches the one in the docker-compos.yml environment, especially in the postgresql database environment.

#### 4. Run client server application using docker-compose

```shell
docker-compose up -d --build
```

### If you want to run client and server applications separately, you can see the guide here:

#### client

#### server


