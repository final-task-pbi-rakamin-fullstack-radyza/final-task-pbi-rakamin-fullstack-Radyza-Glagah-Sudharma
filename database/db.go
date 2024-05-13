package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "log"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
    //Database dapat disesuaikan, untuk kali ini saya menggunakan local database postgresql
    dsn := "host=localhost user=admin dbname=rakamin sslmode=disable password=rahasia"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }
    DB = db
    log.Println("Database connection successfully established")
    return DB
}

func GetDB() *gorm.DB {
    return DB
}
