package database

import (
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
) 
var DB *gorm.DB
 
func InitDB() {
	
    dbHost := os.Getenv("DB_HOST")
    if dbHost == "" {
        dbHost = "localhost" // Fallback for local development
    }
    dbUser := os.Getenv("DB_USER")
    if dbUser == "" {
        dbUser = "postgres"
    }
    dbPassword := os.Getenv("DB_PASSWORD")
    if dbPassword == "" {
        dbPassword = "yehia"
    }
    dbName := os.Getenv("DB_NAME")
    if dbName == "" {
        dbName = "circleConnect"
    }
    dbPort := os.Getenv("DB_PORT")
    if dbPort == "" {
        dbPort = "5432"
    }
     
    connectionString := "host=" + dbHost + 
                        " user=" + dbUser + 
                        " password=" + dbPassword + 
                        " dbname=" + dbName + 
                        " port=" + dbPort + 
                        " sslmode=disable"
						
    var err error
    DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
     
    if err != nil {
        log.Fatal("Could not connect to the database: ", err)
    }
    
    log.Println("Successfully connected to the database")
}