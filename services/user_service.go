package services

import (
    "errors"
    "user-management/models"
    "user-management/storage"
)

func ValidateAndCreateUser(user models.User) error {
    if user.Name == "" || user.Email == "" {
        return errors.New("name and email are required")
    }
    storage.AddUser(user)
    return nil
}

func ValidateAndUpdateUser(id int, updatedUser models.User) error {
    if updatedUser.Name == "" || updatedUser.Email == "" {
        return errors.New("name and email are required")
    }
    return storage.UpdateUser(id, updatedUser)
}
