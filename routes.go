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
        "GM Message",
        "POST",//I think it will be a post coming from gm api
        "/sat/gm",
        GmMsg,
    },
    //all /sat/ functions are for SatCom Box
    Route{
        "Sat Msg RX",
        "POST",
        "/sat/sat",
        SatMsg,
    },
}
