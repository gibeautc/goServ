package main

import (
	"time"
	"log"
	"database/sql"
)

type App struct{
	db *db
}

type Daemon struct{
	db *db
}

type db struct{
	db *sql.DB
}

type location struct{
	Id int
	Lat float32
	Lon	float32
	Name string
	LastUpdate time.Time
	Freq int
	Service int
}

type TReport struct{
	Location string `json:"location"`
	TimeStamp time.Time `json:"timestamp"`
	Value float64 `json:"value"`
}

type SatReport struct{
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

type ApiResponse struct{
	Success bool `json:"success"`
	Code int `json:"code"`
	Message string `json:"message"`
}

type SystemStatus struct{

	Id int16 `json:"id"`					//id of system  -required
	Name string `json:"name"`			//name of system -required
	CpuPercent int16 `json:"cpu_percent"`	//average cpu usage maybe since last update
	CpuTemp int16 `json:"cpu_temp"`		//current cpu temp
	DiskUsage int16 `json:"disk_usage"`	//current disk usage
	Voltage int16 `json:"voltage"`		//current voltage (if running off battery)
	Uptime int16 `json:"uptime"`			//uptime in hours
	Processes map[string]bool `json:"processes"`
	//this is a list of processes that system is watching, and their status
	TimeStamp time.Time `json:"timestamp"`
}


type HourWeather struct{
	Hour int `json:"hour"`
	Temp int `json:"temp"`
	Wind int `json:"wind"`
	Sky int `json:"sky"`
	Precip int `json:"precip"`
}

type DayWeather struct{
	Day int `json:"day"`
	High int `json:"high"`
	Low int `json:"low"`
	Wind int `json:"wind"`
	Sky int `json:"sky"`
	Precip int `json:"precip"`
	RiseTime int64 `json:"rise_time"`
	SetTime int64 `json:"set_time"`
}

type WxReport struct{
	Hourly []*HourWeather `json:"hourly"`
	Dayily []*DayWeather  `json:"daily"`
}

func (self SystemStatus) printStruct(){
	log.Println("ID: ",self.Id)
	log.Println("TimeStamp: ",self.TimeStamp)
	log.Println("Name: ",self.Name)
	log.Println("CPU Percent: ",self.CpuPercent)
	log.Println("CPU TEMP:",self.CpuTemp)
	log.Println("Disk Usage: ",self.DiskUsage)
	log.Println("Voltage: ",self.Voltage)
	log.Println("Uptime: ",self.Uptime)
	log.Println("Processes:")
	for k,v:=range self.Processes{
		log.Println("%s : %t",k,v)
	}

}

func (self TReport) printStruct(){
	log.Println("Location: ",self.Location)
	log.Println("TimeStamp: ",self.TimeStamp)
	log.Println("Value: ",self.Value)


}


const rCFailed=0x01