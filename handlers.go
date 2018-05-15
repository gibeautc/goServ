package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
)


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



func GmMsg(w http.ResponseWriter,r *http.Request){
	log.Println("Requesting Server System Status")
	var apiRes ApiResponse
	apiRes.Message="Write ME!"
	apiRes.Success=false
	_=json.NewDecoder(r.Body).Decode(&apiRes)
	json.NewEncoder(w).Encode(apiRes)
	return
}

func SatMsg(w http.ResponseWriter,r *http.Request){
	/*
	This will be the message from the satBox, so it will be
	lat		4
	lon		4
	bat		1
	status	1

	status bits
	0    	Set=moving
	1
	2
	3
	4
	5
	6
	7


	Minimum of 10 bytes, then anything after 10, would be an ascii encoded message

	*/
	//data:=make([]byte,0)
	app:=new(App)
	app.db=ConnectToDB()
	log.Println("SatCom Box Update")

	if r.ContentLength<10{
		log.Println("Message from SatCom too short")
		return
	}




	//log.Println("Lat: ",t.Lat)
	//log.Println("Lon: ",t.Lon)
	//client := darksky.New("304e4f1db901c61cf8cb2b6d9be6237a")
	//request := darksky.ForecastRequest{}
	//request.Latitude = darksky.Measurement(t.Lat)
	//request.Longitude = darksky.Measurement(t.Lon)
	//request.Options = darksky.ForecastRequestOptions{Exclude: "minutely"}
	//forecast, err := client.Forecast(request)
	//if err!=nil{
	//	log.Println(err.Error())
	//	return
	//}
	//for x:=0;x<len(forecast.Alerts);x++{
	//	log.Println(forecast.Alerts[x].Description)
	//}
	//report,err:=processForecast(forecast)
	//if err!=nil{
	//	log.Println(err.Error())
	//	return
	//}
	////this is where the response it written out
	//json.NewEncoder(w).Encode(report)
	//return



	//python stuff from last time
	//data=request.form['data']
	//data=data.decode('hex')
	//momsn=request.form['momsn']
	//transmit_time=request.form['transmit_time']
	//iridium_lat=request.form['iridium_latitude']
	//iridium_lon=request.form['iridium_longitude']
	//ir_cep=request.form['iridium_cep']

}
