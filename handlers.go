package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"time"
	"github.com/shawntoffel/darksky"
)


    //RX SatCom msg
    //Tx SatCom Msg
	//Get current weather report (current hour,day?) based on location



func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"Welcome to the API")
	
}

func IndexPost(w http.ResponseWriter, r *http.Request) {
	var apiRes ApiResponse
	apiRes.Success=false
	apiRes.Message="ERROR: Cannot Post to index"
	_=json.NewDecoder(r.Body).Decode(&apiRes)
	json.NewEncoder(w).Encode(apiRes)

}

func GetTemp(w http.ResponseWriter, r *http.Request) {
	//fixme: need to look at what location is being requested
	var tmp TReport
	tmp.Location="office"
	tmp.Value=65.3
	json.NewEncoder(w).Encode(tmp)
	}
func ReportTemp(w http.ResponseWriter, r *http.Request) {
	var err error
	var apiRes ApiResponse
	var vars TReport
	apiRes.Success=true

	log.Println(w,"Receiving Temp Report")
	decoder:=json.NewDecoder(r.Body)

	err=decoder.Decode(&vars)
	if err!=nil{
		log.Println("Error Parsing JSON")
		log.Println(err.Error())
		apiRes.Success=false
		apiRes.Message="Failed to parse json"
	}

	log.Println(vars.Location)
	log.Println(vars.Value)
	vars.TimeStamp=time.Now()

	_=json.NewDecoder(r.Body).Decode(&apiRes)
	json.NewEncoder(w).Encode(apiRes)
}

func ReportSystemStatus(w http.ResponseWriter, r *http.Request) {
	var err error
	var vars SystemStatus
	var apiRes ApiResponse
	apiRes.Success=true //unless we have a reason to set it false

	log.Println(w,"Receiving System Status Report")

	decoder:=json.NewDecoder(r.Body)
	err=decoder.Decode(&vars)
	if err!=nil{
		log.Println("Error Parsing JSON")
		log.Println(err.Error())
		apiRes.Success=false
		apiRes.Message="Failed to parse json"
	}
	if vars.Id<1{
		log.Println("Missing/Malformed ID")
		apiRes.Success=false
		apiRes.Message="ERROR: Missing/malfored ID"
	}

	if vars.Name==""{
		log.Println("Missing/Malformed Name")
		apiRes.Success=false
		apiRes.Message="ERROR: Missing/malfored Name"
	}

	vars.printStruct()

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


func GmMsgReceived(w http.ResponseWriter,r *http.Request){
	log.Println("Requesting Server System Status")
	var apiRes ApiResponse
	apiRes.Message="Write ME!"
	apiRes.Success=false
	_=json.NewDecoder(r.Body).Decode(&apiRes)
	json.NewEncoder(w).Encode(apiRes)
	return
}


func GmMsgSend(w http.ResponseWriter,r *http.Request){
	log.Println("Requesting Server System Status")
	var apiRes ApiResponse
	apiRes.Message="Write ME!"
	apiRes.Success=false
	_=json.NewDecoder(r.Body).Decode(&apiRes)
	json.NewEncoder(w).Encode(apiRes)
	return
}

func SatWx(w http.ResponseWriter,r *http.Request){
	app:=new(App)
	app.db=ConnectToDB()
	log.Println("SatCom Box WX Request")
	decoder:=json.NewDecoder(r.Body)
	var t SatReport
	err:=decoder.Decode(&t)
	if err!=nil{
		log.Println(err.Error())
		return
	}
	log.Println("Lat: ",t.Lat)
	log.Println("Lon: ",t.Lon)
	client := darksky.New("304e4f1db901c61cf8cb2b6d9be6237a")
	request := darksky.ForecastRequest{}
	request.Latitude = darksky.Measurement(t.Lat)
	request.Longitude = darksky.Measurement(t.Lon)
	request.Options = darksky.ForecastRequestOptions{Exclude: "minutely"}
	forecast, err := client.Forecast(request)
	if err!=nil{
		log.Println(err.Error())
		return
	}
	for x:=0;x<len(forecast.Alerts);x++{
		log.Println(forecast.Alerts[x].Description)
	}
	report,err:=processForecast(forecast)
	if err!=nil{
		log.Println(err.Error())
		return
	}
	//this is where the response it written out
	json.NewEncoder(w).Encode(report)
	return
}
