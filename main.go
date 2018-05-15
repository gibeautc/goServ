package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	//Get Port number as argument or config file?
	log.Println("Starting SatComServer....")
	go daemon()
	router:=NewRouter()
	err:=sendGMMsg("Server Startup")
	if err!=nil{
		log.Println(err.Error())
	}
	err=sendSatMsg([]byte("Server Startup"))
	if err!=nil{
		log.Println(err.Error())
	}

	log.Fatal(http.ListenAndServe(":9090", router))
}


func daemon(){
	dae:=new(Daemon)
	dae.db=ConnectToDB()
	log.Println("Worker Daemon Starting...")
	for{
		dae.checkUpdate()
		time.Sleep(1*time.Second)
	}
}


func (d *Daemon) checkUpdate(){

}