package main

import (
       "github.com/ant0ine/go-json-rest"
       "net/http"
)

func main(){
     handler := rest.ResourceHandler {
     	     	EnableRelaxedContentType: true,
		}
     handler.SetRoutes(
		rest.Route{"GET", "/store", GetAllEntriesHandler},
		rest.Route{"POST", "/store", PostEntryHandler},
		rest.Route{"GET", "/store/:Key", GetEntryHandler},
		rest.Route{"DELETE", "/store/:Key", DeleteEntryHandler},
		rest.Route{"GET", "/store/incr/:Key", IncrEntryHandler},		//change this to UPDATE?
		rest.Route{"GET", "/store/decr/:Key", DecrEntryHandler},
		)
     http.ListenAndServe(":9090", &handler)
}