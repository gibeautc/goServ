package main

import "time"

type TReport struct{
	Location string `json:"location"`
	TimeStamp time.Time `json:"timestamp"`
	Value float64 `json:"value"`
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
}



const rCFailed=0x01