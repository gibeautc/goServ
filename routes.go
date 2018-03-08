package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler

        handler = route.HandlerFunc
        handler = Logger(handler, route.Name)
        
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }

    return router
}

var routes = Routes{
    Route{
        "IndexPost",
        "POST",
        "/",
        IndexPost,
    },
    Route{
        "Index",
        "GET",
        "/",
        Index,
    },
    Route{
        "GetOfficeTemp",
        "GET",
        "/getTemp",
        GetTemp,
    },
    Route{
        "PostOfficeTemp",
        "POST",
        "/reportTemp/",
        ReportTemp,
    },
    Route{
        "PostSystemStatus",
        "POST",
        "/systemstatus/",
        ReportSystemStatus,
    },
    Route{
        "GetSystemStatus",
        "GET",
        "/systemstatus/",
        GetSystemStatus,
    },
    Route{
        "GM Call Back",
        "POST",//I think it will be a post coming from gm api
        "/gmcallback/",
        GmMsgReceived,
    },
    Route{
        "Send GM Message",
        "POST",//I think it will be a post coming from gm api
        "/gmSend/",
        GmMsgSend,
    },
    //all /sat/ functions are for SatCom Box
    Route{
        "Sat WX request",
        "POST",
        "/sat/wx",
        SatWx,
    },
}
