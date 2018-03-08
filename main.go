package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	//Get Port number as argument or config file?
	log.Println("Starting....")
	go daemon()
	router:=NewRouter()
    log.Fatal(http.ListenAndServe(":5000", router))
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