package database

import (
	"fmt"

	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
    dsn := "root:root@tcp(db:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
    
    for {
        db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err != nil {
            fmt.Println("MySQL not yet ready, waiting...")
            time.Sleep(1 * time.Second)
        } else {
            fmt.Println("Connection Opened to Database!")
            DB = db
            break  
        }
    }
}