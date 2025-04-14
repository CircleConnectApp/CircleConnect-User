package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
) 
var DB *gorm.DB
 
func InitDB() {
	
    dbHost := "localhost"
    dbUser := "postgres"
    dbPassword := "yehia"
    dbName := "circleConnect"
    dbPort := "5432"
     
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