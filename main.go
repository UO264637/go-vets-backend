package main

import (
    "project/go-vets-backend/db"
    "project/go-vets-backend/handlers"
)

func main() {
    db.ConnectDB()
    handlers.Handlers()
}
 /*
{
  "name": "Vicente",
  "surname":"García Díaz",
  "email":"vi@gmail.com",
  "password":"123456",
  "birthDate":"2023-01-14T15:30:01Z"
}
 */