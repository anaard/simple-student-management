# Golang Student Management System REST API

This Golang API provides basic JWT Authentication and CRUD operations to manage a student system.

## Setup

* Change the provided .env file in the project's root directory to include all your variables.
* Run the following command to create and start the database and development containers (requires docker and docker compose):

  ```
  docker compose up
  ```

## Tecnologies Used

* Language: Golang
* Database: MySQL

### Libraries Used

* Server: net/http
* Router: gorilla/mux
* JWT: dgrijalva/go-jwt
* Database ORM: gorm
* Input Validation: go-playgroud/validator

## API Documentation

### Authentication

#### POST `/register`

###### Body

```json
{
    "name":"example_name",
    "username": "example_username",
    "password": "example_password",
    "email":"example_email@gmail.com"
}
```

##### Output

```json
{
    "User": {
        "ID": "id_number",
        "CreatedAt": "date_creation",
        "UpdatedAt": "date_update",
        "DeletedAt": null,
        "name": "example_name",
        "email": "example_email@gmail.com",
        "username": "example_username",
        "password": "example_password"
    },
    "Token": "example_token"
}
```

#### POST `/login`

###### Body

```json
{
    "id": "email_or_username",
    "password": "example_password"
}
```

##### Output

```json
{
    "User": {
        "ID": "id_number",
        "CreatedAt": "date_creation",
        "UpdatedAt": "date_update",
        "DeletedAt": null,
        "name": "example_name",
        "email": "example_email@gmail.com",
        "username": "example_username",
        "password": "example_password"
    },
    "Token": "example_token"
}

```

### Data Manipulation

To execute any of the following requests, you must include the received token with your request. In Postman, navigate to the "Authorization" tab, choose "Bearer Token" as the Type, and paste the token into the provided field.

#### POST `/student`

###### Body

```json
{
    "name": "example_name",
    "total_faults": 0,
    "grade_average": 70
}
```

##### Output

```json
{
    "ID": "example_id",
    "CreatedAt": "creation_time",
    "UpdatedAt": "uptade_time",
    "DeletedAt": null,
    "name": "example_name",
    "class_id": 0,
    "total_faults": 0,
    "grade_average": 70
}
```

#### GET `/student/all`

##### Output

```json
[
    {
        "ID": "example_id",
        "CreatedAt": "creation_time",
        "UpdatedAt": "update_time",
        "DeletedAt": null,
        "name": "example_name",
        "class_id": 0,
        "total_faults": 0,
        "grade_average": 70
    }
]
```

#### GET `/student/<id>`

##### Output

```json
{
    "ID": "example_id",
    "CreatedAt": "creation_time",
    "UpdatedAt": "update_time",
    "DeletedAt": null,
    "name": "example_name",
    "class_id": 0,
    "total_faults": 0,
    "grade_average": 70
}
```

#### PUT `/student/<id>`

###### Body

```json
{
    "name": "corrected_name",
    "total_faults": 0,
    "grade_average": 70
}
```

##### Output

```json
{
    "ID": "example_id",
    "CreatedAt": "creation_time",
    "UpdatedAt": "update_time",
    "DeletedAt": null,
    "name": "corrected_name",
    "class_id": 0,
    "total_faults": 0,
    "grade_average": 70
}
```

#### DELETE `/student/<id>`

##### Output (The student is deleted from the database and the class)

```json
{
    "ID": "example_id",
    "CreatedAt": "2024-02-17T16:14:46Z",
    "UpdatedAt": "2024-02-17T16:14:46Z",
    "DeletedAt": null,
    "name": "example_name",
    "class": 0,
    "total_faults": 0,
    "grade_average": 70
}
```

#### POST `/class` (The class is always created empty)

##### Output

```json
{
    "ID": "example_id",
    "CreatedAt": "creation_time",
    "UpdatedAt": "update_time",
    "DeletedAt": null,
    "TotalStudentNumber": 0,
    "Students": null
}
```

#### GET `/class/all`

##### Output

```json
[
    {
        "ID": "example_id",
        "CreatedAt": "creation_time",
        "UpdatedAt": "update_time",
        "DeletedAt": null,
        "TotalStudentNumber": 0,
        "Students": null
    }
]
```

#### GET `/class/<id>`

##### Output

```json
{
    "ID": "example_id",
    "CreatedAt": "creation_time",
    "UpdatedAt": "update_time",
    "DeletedAt": null,
    "TotalStudentNumber": 0,
    "Students": null
}
```

#### DELETE `/class/<id>`

##### Output

```json
{
    "ID": "example_id",
    "CreatedAt": "creation_time",
    "UpdatedAt": "update_time",
    "DeletedAt": null,
    "TotalStudentNumber": "student_number",
    "Students": null
}
```

#### POST `/<classId>/<studentId>`

##### Output (The class_id will be equal to the Id provided in the URL and the student will be added to the Students attribute in the Class struct)

```json
{
    "ID": "example_id",
    "CreatedAt": "creation_time",
    "UpdatedAt": "update_time",
    "DeletedAt": null,
    "name": "example_name",
    "class_id": "changed_id",
    "total_faults": 0,
    "grade_average": 70
}
```

#### DELETE `/<classId>/<studentId>`

##### Output (The class_id will be equal to 0 and the student will be removed from the Students attribute in the Class struct)

```json
{
    "ID": "example_id",
    "CreatedAt": "creation_time",
    "UpdatedAt": "update_time",
    "DeletedAt": null,
    "name": "example_name",
    "class_id": 0,
    "total_faults": 0,
    "grade_average": 70
}
```
