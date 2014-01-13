package main

import (
       "github.com/pkhadilkar/sskv/server"
       "github.com/ant0ine/go-json-rest"
       "net/http"
)

func main(){
     handler := rest.ResourceHandler {
     	     	EnableRelaxedContentType: true,
		}
     handler.SetRoutes(
		rest.Route{"GET", "/store", server.GetAllEntriesHandler},
		rest.Route{"POST", "/store", server.PostEntryHandler},
		rest.Route{"GET", "/store/:Key", server.GetEntryHandler},
		rest.Route{"DELETE", "/store/:Key", server.DeleteEntryHandler},
		rest.Route{"GET", "/store/incr/:Key", server.IncrEntryHandler},		//change this to UPDATE?
		rest.Route{"GET", "/store/decr/:Key", server.DecrEntryHandler},
		)
     http.ListenAndServe(":9090", &handler)
}