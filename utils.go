package main

import (
	"github.com/shawntoffel/darksky"
	"time"
	"net/http"
	"bytes"
	//"io/ioutil"
	"log"
	"errors"
	"math"
	"io/ioutil"
)

func sendGMMsg(msg string) error {
	//curl -d '{"text" : "Your message here", "bot_id" : "b3e83fd81cfbe44a7ea8a22030"}' https://api.groupme.com/v3/bots/post
	url := "https://api.groupme.com/v3/bots/post"
	log.Println("URL:>", url)
	var jsonStr = []byte(`{"text":"`+msg+`","bot_id":"b3e83fd81cfbe44a7ea8a22030"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.Status!="202 Accepted"{
		return errors.New("Not Excepted:"+resp.Status)

	}
	//log.Println("response Status:", resp.Status)
	//log.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	return nil
}


func sendSatMsg(msg []byte) error {
	log.Println("Sending Message to Sat, Length:",len(msg))
	//curl -d '{"text" : "Your message here", "bot_id" : "b3e83fd81cfbe44a7ea8a22030"}' https://api.groupme.com/v3/bots/post
	url := "https://core.rock7.com/rockblock/MT"
	log.Println("URL:>", url)
	var jsonStr = []byte(`{"data":"`+"11111111"+`","imei":"300234064380130","username":"gibeautc@oregonstate.edu","password":"myvice12"}`)
	//var jsonStr = []byte(`{"data":"`+hex.Dump(msg)+`","imei":"300234064380130","username":"gibeautc@oregonstate.edu","password":"myvice12"}`)
	log.Println("Trying to send:",string(jsonStr))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//if resp.Status!="202 Accepted"{
	//	return errors.New("Not Excepted:"+resp.Status)
	//
	//}
	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	bodyStr:=string(body)
	log.Println("Body:",bodyStr)
	//OK,12345678    will be unique number
	//FAILED,15,Textual description of failure
	//10-Invalid login credentials
	//11-No RockBLOCK with this IMEI found on your account
	//12-RockBLOCK has no line rental
	//13-Your account has insufficient credit
	//14-Could not decode hex data
	//15-Data too long
	//16-No data
	return nil
}



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

func fakeCheckIn(){
	lat:=45.123
	lon:=123.123
	Status:=0
	bat:=12.4
	msg:="Test Message"
	var tmp [4]byte
	var buf []byte

	//lat
	n := math.Float32bits(float32(lat))
	tmp[0] = byte(n >> 24)
	tmp[1] = byte(n >> 16)
	tmp[2] = byte(n >> 8)
	tmp[3] = byte(n)
	for x:=0;x<len(tmp);x++{
		buf=append(buf,tmp[x])
	}

	//lon
	n = math.Float32bits(float32(lon))
	tmp[0] = byte(n >> 24)
	tmp[1] = byte(n >> 16)
	tmp[2] = byte(n >> 8)
	tmp[3] = byte(n)
	for x:=0;x<len(tmp);x++{
		buf=append(buf,tmp[x])
	}

	buf=append(buf,byte(Status))
	buf=append(buf,byte(bat*100))
	msgByte:=[]byte(msg)
	for x:=0;x<len(msgByte);x++{
		buf=append(buf,msgByte[x])
	}

}