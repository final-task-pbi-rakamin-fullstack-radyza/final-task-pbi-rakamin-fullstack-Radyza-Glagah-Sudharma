package main

import (
    "project-final/router"
    "project-final/database"
    "log"
)

func main() {
    database.InitDB() 
    r := router.SetupRouter()  
    err := r.Run()  
    if err != nil {
        log.Fatal("Error starting server: ", err)
    }
}
