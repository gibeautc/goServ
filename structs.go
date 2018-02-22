package main

import "time"

type TReport struct{
	Location string `json:"location"`
	TimeStamp time.Time `json:"timestamp"`
	Value float32 `json:"value"`
}

