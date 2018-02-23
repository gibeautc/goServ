package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"github.com/gorilla/mux"
	"log"
	"github.com/gorilla/mux"
	"time"
	"strconv"
)



	//report outside temp (Home)
    //get outside Temp (Home)
    //report Hsink Temp
    //get Hsink Temp
    //Report Node Status
    //Get Node Status
    //report 
    //Receive GM Message
    //Send GM Message
    //RX SatCom msg
    //Tx SatCom Msg




func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"Welcome to the API")
	
}

func IndexPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"Welcome to the API, you cannot POST to INDEX")

}

func GetOfficeTemp(w http.ResponseWriter, r *http.Request) {
	var tmp TReport
	tmp.Location="Office"
	tmp.Value=65.3
	json.NewEncoder(w).Encode(tmp)
	}
func ReportOfficeTemp(w http.ResponseWriter, r *http.Request) {
	log.Println(w,"Receiving Office Temp Report")
	vars:=mux.Vars(r)
	log.Println("Body Received")
	log.Println(r.Body)
	var tmp TReport
	var apiRes ApiResponse
	tmp.Location=vars["location"]
	tmp.Value,_=strconv.ParseFloat(vars["value"],64)
	tmp.TimeStamp=time.Now()
	apiRes.Success=true
	_=json.NewDecoder(r.Body).Decode(&apiRes)
	json.NewEncoder(w).Encode(apiRes)
}

func ReportSystemStatus(w http.ResponseWriter, r *http.Request) {
	log.Println(w,"Receiving Office Temp Report")
	vars:=mux.Vars(r)
	var tmp SystemStatus
	var apiRes ApiResponse
	apiRes.Success=true //unless we have a reason to set it false
	t,err:=strconv.ParseInt(vars["id"],10,16)
	if err!=nil{
		log.Println("Error Getting ID- Required")
		log.Println(err.Error())
		log.Println(vars["id"])
		apiRes.Message="ID ERROR"
		apiRes.Success=false
		_=json.NewDecoder(r.Body).Decode(&apiRes)
		json.NewEncoder(w).Encode(apiRes)
		return
	}
	tmp.Id=int16(t)

	tmp.Name=vars["name"]
	if tmp.Name==""{
		log.Println("Error Getting Name- Required")
		log.Println(err.Error())
		apiRes.Message="Name ERROR"
		apiRes.Success=false
		_=json.NewDecoder(r.Body).Decode(&apiRes)
		json.NewEncoder(w).Encode(apiRes)
		return
	}


	t,err=strconv.ParseInt(vars["cpu_percent"],10,16)
	if err!=nil{
		log.Println("Error Getting CpuPercent")
		log.Println(err.Error())
		apiRes.Message=apiRes.Message+"| ERROR on CpuPercent"

	}
	_=json.NewDecoder(r.Body).Decode(&apiRes)
	json.NewEncoder(w).Encode(apiRes)
	return
}

func GetSystemStatus(w http.ResponseWriter, r *http.Request){
	log.Println("Requesting Server System Status")
	var apiRes ApiResponse
	apiRes.Message="Write ME!"
	apiRes.Success=false
	_=json.NewDecoder(r.Body).Decode(&apiRes)
	json.NewEncoder(w).Encode(apiRes)
	return
}
