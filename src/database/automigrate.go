package database

import models "github.com/champseo2017/go_project_learn_2024/src/models"

func Automigrate() {
    DB.AutoMigrate(models.User{}, models.Cart{}, models.Product{})
}