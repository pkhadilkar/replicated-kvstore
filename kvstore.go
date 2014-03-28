package main

import (
	"github.com/ant0ine/go-json-rest"
	"github.com/pkhadilkar/replicated-kvstore/server"
	"net/http"
)

func main() {
	handler := rest.ResourceHandler{
		EnableRelaxedContentType: true,
	}
	handler.SetRoutes(
		rest.Route{"POST", "/store", server.PostEntryHandler},
		rest.Route{"GET", "/store/:Key", server.GetEntryHandler},
		rest.Route{"DELETE", "/store/:Key", server.DeleteEntryHandler},
		rest.Route{"GET", "/store/incr/:Key", server.IncrEntryHandler}, //change this to UPDATE?
		rest.Route{"GET", "/store/decr/:Key", server.DecrEntryHandler},
	)
	// initialize kvStore
	server.Initialize()
	http.ListenAndServe(":63000", &handler)
}
