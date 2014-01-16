package server

import (
       "github.com/ant0ine/go-json-rest"
       "net/http"
       "net/url"
)

// Frontend module for single server key store

func Initialize(){
     token <- true
}

func GetEntryHandler(w *rest.ResponseWriter, r *rest.Request){
     key, err := url.QueryUnescape(r.PathParam("Key"))
     w.Header().Set("Content-type", "application/json")
     if err != nil {
     	rest.Error(w, err.Error(), http.StatusInternalServerError)
     }
     value, ok := GetValue(key)
     if !ok {
     	rest.NotFound(w, r)
	return
     }
     w.WriteJson(&value)
}

func GetAllEntriesHandler(w *rest.ResponseWriter, r *rest.Request){
     entries := GetAllEntries()
     w.Header().Set("Content-type", "application/json")
     w.WriteJson(entries)
}

func PostEntryHandler(w *rest.ResponseWriter, r *rest.Request){
     entry := Entry{}
     err := r.DecodeJsonPayload(&entry)
     w.Header().Set("Content-type", "application/json")
     if err != nil {
     	rest.Error(w, err.Error(), http.StatusBadRequest)
	return
     }
     
     if entry.Key == "" {
     	rest.Error(w, "Key should not be empty", http.StatusBadRequest)
	return
     }

     if entry.Value == "" {
     	rest.Error(w, "Value should not be empty", http.StatusBadRequest)
	return
     }
     PutValue(&entry)
     w.WriteJson(&entry)
}

func DeleteEntryHandler(w *rest.ResponseWriter, r *rest.Request){
     key, err := url.QueryUnescape(r.PathParam("Key"))
     if err != nil {
     	rest.Error(w, err.Error(), http.StatusBadRequest)
     }
	 _, ok := GetValue(key)
     if !ok {
     	rest.NotFound(w, r)
		return
     }
     DeleteEntry(key)
}

func IncrEntryHandler(w *rest.ResponseWriter, r *rest.Request){
     key, err := url.QueryUnescape(r.PathParam("Key"))
     w.Header().Set("Content-type", "application/json")
     if(err != nil){
     	    rest.Error(w, err.Error(), http.StatusBadRequest)
     }
     value, incErr := IncrEntry(key)
     if incErr != nil {
     	rest.Error(w, incErr.Error(), http.StatusPreconditionFailed)
     }
     w.WriteJson(&value)
}

func DecrEntryHandler(w *rest.ResponseWriter, r *rest.Request){
     key, err := url.QueryUnescape(r.PathParam("Key"))
     w.Header().Set("Content-type", "application/json")
     if(err != nil){
     	    rest.Error(w, err.Error(), http.StatusBadRequest)
     }
     value, decErr := DecrEntry(key)
     if decErr != nil {
     	rest.Error(w, decErr.Error(), http.StatusPreconditionFailed)
     }
     w.WriteJson(&value)
}