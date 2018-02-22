package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "log"



func ConnectToApiDB() *sql.DB{
	
	
db, err := sql.Open("mysql", "root:aq12ws@/api")
if err!=nil{
	log.Println("Failed to open Database Connection")
	return nil
}
return db

}


func ConnectToWeatherDB() *sql.DB{


	db, err := sql.Open("mysql", "root:aq12ws@/weather")
	if err!=nil{
		log.Println("Failed to open Database Connection")
		return nil
	}
	return db

}

func ConnectToDarkSkyDB() *sql.DB{


	db, err := sql.Open("mysql", "root:aq12ws@/darksky")
	if err!=nil{
		log.Println("Failed to open Database Connection")
		return nil
	}
	return db

}
