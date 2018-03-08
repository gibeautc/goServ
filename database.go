package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import (
	"log"
	"encoding/json"
)



func ConnectToDB() *db{
	d:=new(db)
	db, err := sql.Open("mysql", "root:aq12ws@/server")
	if err!=nil {
		log.Println("Failed to open Server Database Connection")
		return nil
	}

	err=db.Ping()
	if err !=nil{
		log.Println("Error Pinging Server Database")
		return nil
	}
	d.db=db
	return d

}


func (d db) AddSystemStatus(db *sql.DB,data *json.Decoder) {
	log.Println("Adding System Status to DB")
}


func (d db) GetCurrentSystemStatus(db *sql.DB,id int){
	log.Println("Getting Last System Status for ID: ",id)
}

func (d db) GetLocations() []location{
	var locs []location
	log.Println("Getting Locations from database")
	stmtOut, err := d.db.Prepare("SELECT * FROM locations")
	if err!=nil{
		log.Println(err.Error())
		return locs
	}
	defer stmtOut.Close()
	rows,err := stmtOut.Query()
	for rows.Next(){
		tmp:=new(location)
		rows.Scan(&tmp.Id,&tmp.Name,&tmp.Lat,&tmp.Lon,&tmp.LastUpdate,&tmp.Freq,&tmp.Service)
	}
	if err != nil {
		log.Println(err.Error())
	}


	return locs

}



//// Prepare statement for inserting data
//stmtIns, err := db.Prepare("INSERT INTO squareNum VALUES( ?, ? )") // ? = placeholder
//if err != nil {
//panic(err.Error()) // proper error handling instead of panic in your app
//}
//defer stmtIns.Close() // Close the statement when we leave main() / the program terminates
//
//// Prepare statement for reading data
//stmtOut, err := db.Prepare("SELECT squareNumber FROM squarenum WHERE number = ?")
//if err != nil {
//panic(err.Error()) // proper error handling instead of panic in your app
//}
//defer stmtOut.Close()
//
//// Insert square numbers for 0-24 in the database
//for i := 0; i < 25; i++ {
//_, err = stmtIns.Exec(i, (i * i)) // Insert tuples (i, i^2)
//if err != nil {
//panic(err.Error()) // proper error handling instead of panic in your app
//}
//}
//
//var squareNum int // we "scan" the result in here
//
//// Query the square-number of 13
//err = stmtOut.QueryRow(13).Scan(&squareNum) // WHERE number = 13
//if err != nil {
//panic(err.Error()) // proper error handling instead of panic in your app
//}
//fmt.Printf("The square number of 13 is: %d", squareNum)
//
//// Query another number.. 1 maybe?
//err = stmtOut.QueryRow(1).Scan(&squareNum) // WHERE number = 1
//if err != nil {
//panic(err.Error()) // proper error handling instead of panic in your app
//}
//fmt.Printf("The square number of 1 is: %d", squareNum)
//}