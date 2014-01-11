package main

import (
       "github.com/ant0ine/go-json-rest"
       "net/http"
       "net/url"
)

// Single Server Key value store implementation

type Entry struct{
     Key string
     Value string
}

// Simple Implementation: v1 Use a global HashMap
var kvStore map[string]string

func GetEntry(w *rest.ResponseWriter, r *rest.Request){
     key, err := url.QueryUnescape(r.PathParam("key"))
     if err != nil {
     	rest.Error(w, err.Error(), http.StatusInternalServerError)
     }
     value := kvStore[key]
     if value == "" {
     	rest.NotFound(w, r)
	return
     }
     w.WriteJson(&value)
}

func GetAllEntries(w *rest.ResponseWriter, r *rest.Request){
     entries := make([]*Entry, len(kvStore))
     i := 0
     for key, value := range kvStore {
     	 entry := Entry{key, value}
     	 entries[i] = &entry
	 i += 1
     }
     w.WriteJson(&entries)
}

func PostEntry(w *rest.ResponseWriter, r *rest.Request){
     entry := Entry{}
     err := r.DecodeJsonPayload(&entry)
     if err != nil {
     	rest.Error(w, err.Error(), http.StatusInternalServerError)
	return
     }
     
     if entry.Key == "" {
     	rest.Error(w, "Key should not be empty", 400)
	return
     }

     if entry.Value == "" {
     	rest.Error(w, "Value should not be empty", 400)
	return
     }
     kvStore[entry.Key] = entry.Value
     w.WriteJson(&entry)
}

func DeleteEntry(w *rest.ResponseWriter, r *rest.Request){
     key := r.PathParam("key")
     delete(kvStore, key)
}


func main(){
     kvStore = make(map[string]string)
          handler := rest.ResourceHandler {
     	     	EnableRelaxedContentType: true,
		}
     handler.SetRoutes(
		rest.Route{"GET", "/store", GetAllEntries},
		rest.Route{"POST", "/store", PostEntry},
		rest.Route{"GET", "/store/:key", GetEntry},
		rest.Route{"DELETE", "/store/:key", DeleteEntry},
		)
     http.ListenAndServe(":9090", &handler)
}