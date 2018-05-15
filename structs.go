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


func (self TReport) printStruct(){
	log.Println("Location: ",self.Location)
	log.Println("TimeStamp: ",self.TimeStamp)
	log.Println("Value: ",self.Value)


}


const rCFailed=0x01