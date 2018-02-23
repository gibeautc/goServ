package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import (
	"log"
	"encoding/json"
)



func ConnectToApiDB() *sql.DB{

	db, err := sql.Open("mysql", "root:aq12ws@/api")
	if err!=nil {
		log.Println("Failed to open Api Database Connection")
		return nil
	}

	err=db.Ping()
	if err !=nil{
		log.Println("Error Pinging Api Database")
		return nil
	}
return db

}


func ConnectToWeatherDB() *sql.DB{


	db, err := sql.Open("mysql", "root:aq12ws@/weather")
	if err!=nil{
		log.Println("Failed to open Weather Database Connection")
		return nil
	}
	err=db.Ping()
	if err !=nil{
		log.Println("Error Pinging Weather Database")
		return nil
	}
	return db

}

func ConnectToDarkSkyDB() *sql.DB{


	db, err := sql.Open("mysql", "root:aq12ws@/darksky")
	if err!=nil{
		log.Println("Failed to open DarkSky Database Connection")
		return nil
	}
	err=db.Ping()
	if err !=nil{
		log.Println("Error Pinging DarkSky Database")
		return nil
	}
	return db

}



func AddSystemStatus(db *sql.DB,data *json.Decoder) {
	log.Println("Adding System Status to DB")
}


func GetCurrentSystemStatus(db *sql.DB,id int){
	log.Println("Getting Last System Status for ID: ",id)
}