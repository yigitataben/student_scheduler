# Student Scheduler

## Introduction

Student Scheduler is a web application designed to help students manage their lecture schedules efficiently. It allows
users to sign up, create lecture plans, manage their plans, and interact with their schedule seamlessly.

## Table of Contents

- [Installation](#installation)
- [Configuration](#configuration)
- [Database Migration](#database-migration)
- [Running the Application](#running-the-application)
- [API Endpoints](#api-endpoints)
    - [User Endpoints](#user-endpoints)
    - [Lecture Endpoints](#lecture-endpoints)
    - [Plan Endpoints](#plan-endpoints)
- [Project Structure](#project-structure)
- [License](#license)

## Installation

1. Clone the repository:

```bash
    git clone https://github.com/yigitataben/student_scheduler.git
```

2. Navigate to the project directory:

```bash
    cd student_scheduler
```

## Configuration

3. Set up the environment variables by creating a .env file in the root directory. Include the following variables:

```makefile
    DB = your_mysql_connection_string
    PORT = desired_port_number
```

## Database Migration

4. Install dependencies:

```bash
    go mod tidy
```

5. Run database migration to sync the database schema:

```bash
    go run initializers/main.go
```

## Running the Application

6. Build and run the application:

```bash
    go build -o main .
    ./main
```

7. Access the application at http://localhost:PORT in your browser.


## API Endpoints

### User Endpoints

    POST /signup: Sign up a new user.
    GET /users: Retrieve all users.
    GET /users/:id: Retrieve a user by ID.
    PUT /users/:id: Update a user's information.
    DELETE /users/:id: Delete a user by ID.

### Lecture Endpoints

    POST /lectures: Create new lectures.
    GET /lectures: Retrieve all lectures.
    GET /lectures/:id: Retrieve a lecture by ID.
    DELETE /lectures/:id: Delete a lecture by ID.

### Plan Endpoints

    POST /plans: Create a new plan.
    GET /plans: Retrieve all plans.
    GET /plans/:id: Retrieve a plan by ID.
    PUT /plans/:id: Update a plan by ID.
    DELETE /plans/:id: Delete a plan by ID.

## Project Structure

The project structure follows the MVC (Model-View-Controller) pattern:

- *controllers/*: Contains the controller logic for handling HTTP requests.
- *initializers/*: Contains initialization logic such as database connection and schema migration.
- *models/*: Defines the data models used in the application.
- *repositories/*: Contains the repository logic for interacting with the database.
- *routes/*: Defines the HTTP routes and maps them to controller methods.
- *services/*: Contains the business logic for handling application-specific tasks.
- *main.go*: Entry point of the application.

## License

[MIT](https://choosealicense.com/licenses/mit/)

