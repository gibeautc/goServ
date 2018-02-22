package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	//"github.com/gorilla/mux"
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

func GetOfficeTemp(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w,"Request of Office Temp")
	var tmp TReport
	tmp.Location="Office"
	tmp.Value=65.3
	json.NewEncoder(w).Encode(tmp)
	}
func ReportOfficeTemp(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w,"Receiving Office Temp Report")
	//params :=mux.Vars(r)
	var tmp TReport
	_=json.NewDecoder(r.Body).Decode(&tmp)
	json.NewEncoder(w).Encode(tmp)
}
