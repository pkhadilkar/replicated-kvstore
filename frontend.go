package main

import (
       "github.com/ant0ine/go-json-rest"
       "net/http"
       "net/url"
)

// Frontend module for single server key store

func GetEntryHandler(w *rest.ResponseWriter, r *rest.Request){
     key, err := url.QueryUnescape(r.PathParam("Key"))
     if err != nil {
     	rest.Error(w, err.Error(), http.StatusInternalServerError)
     }
     value, ok := GetValue(key)
     if !ok {
     	rest.NotFound(w, r)
	return
     }
     w.Header().Set("Content-type", "application/json")
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
     PutValue(&entry)
     w.Header().Set("Content-type", "application/json")
     w.WriteJson(&entry)
}

func DeleteEntryHandler(w *rest.ResponseWriter, r *rest.Request){
     key, err := url.QueryUnescape(r.PathParam("Key"))
	 if err != nil {
     	rest.Error(w, err.Error(), http.StatusInternalServerError)
     }
	 _, ok := GetValue(key)
     if !ok {
     	rest.NotFound(w, r)
		return
     }
     DeleteEntry(key)
}
