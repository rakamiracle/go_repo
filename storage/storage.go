package storage

import (
    "errors"
    "user-management/models"
)

var users = []models.User{
    {ID: 1, Name: "Alice", Email: "alice@example.com", Password: "password123"},
    {ID: 2, Name: "Bob", Email: "bob@example.com", Password: "123456"},
}

func GetUsers() []models.User {
    return users
}

func GetUserByID(id int) (*models.User, error) {
    for _, user := range users {
        if user.ID == id {
            return &user, nil
        }
    }
    return nil, errors.New("user not found")
}

func AddUser(user models.User) {
    users = append(users, user)
}

func UpdateUser(id int, updatedUser models.User) error {
    for i, user := range users {
        if user.ID == id {
            users[i] = updatedUser
            return nil
        }
    }
    return errors.New("user not found")
}

func DeleteUser(id int) error {
    for i, user := range users {
        if user.ID == id {
            users = append(users[:i], users[i+1:]...)
            return nil
        }
    }
    return errors.New("user not found")
}
