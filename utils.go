package main

import (
	"github.com/shawntoffel/darksky"
	"time"
)

func processForecast(forecast darksky.ForecastResponse) (WxReport,error){
	var ret WxReport
	for x:=0;x<len(forecast.Hourly.Data);x++{
		d:=forecast.Hourly.Data[x]
		h:=new(HourWeather)
		h.Hour=time.Unix(int64(d.Time),0).Hour()
		h.Temp=int(d.ApparentTemperature)
		h.Sky=int(d.CloudCover*100)
		h.Precip=int(d.PrecipIntensity*100)
		h.Wind=int(d.WindSpeed)
		ret.Hourly=append(ret.Hourly,h)
		if len(ret.Hourly)>19{
			break
		}
	}
	for x:=0;x<len(forecast.Daily.Data);x++{
		d:=forecast.Daily.Data[x]
		h:=new(DayWeather)
		h.Day=time.Unix(int64(d.Time),0).Day()
		h.High=int(d.ApparentTemperatureHigh)
		h.Low=int(d.ApparentTemperatureLow)
		h.Sky=int(d.CloudCover*100)
		h.Precip=int(d.PrecipIntensity*10000)
		h.Wind=int(d.WindSpeed)
		//h.RiseTime=time.Unix(int64(d.SunriseTime),0)
		//h.SetTime=time.Unix(int64(d.SunsetTime),0)
		h.RiseTime=int64(d.SunriseTime)
		h.SetTime=int64(d.SunsetTime)

		ret.Dayily=append(ret.Dayily,h)
		if len(ret.Dayily)>6{
			break
		}
	}

	return ret,nil
}