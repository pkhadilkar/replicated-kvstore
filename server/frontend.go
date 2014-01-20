package server

import (
       "github.com/ant0ine/go-json-rest"
       "net/http"
       "net/url"
)

// Frontend module for single server key store

var s kvStore

// Initialize method intializes internal data structures
func Initialize() {
     s.store = make(map[string]ValueWrapper, 100000)
}

func GetEntryHandler(w *rest.ResponseWriter, r *rest.Request){
     key, err := url.QueryUnescape(r.PathParam("Key"))
     defer r.Body.Close()
     w.Header().Set("Content-type", "application/json")
     if err != nil {
     	rest.Error(w, err.Error(), http.StatusInternalServerError)
     }
     value, ok := s.GetValue(key)
     if !ok {
     	rest.NotFound(w, r)
	return
     }
     w.WriteJson(&value)
}

func GetAllEntriesHandler(w *rest.ResponseWriter, r *rest.Request){
     entries := s.GetAllEntries()
     defer r.Body.Close()
     w.Header().Set("Content-type", "application/json")
     w.WriteJson(entries)
}

func PostEntryHandler(w *rest.ResponseWriter, r *rest.Request){
     entry := Entry{}
     defer r.Body.Close()
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
     s.PutValue(&entry)
     w.WriteJson(&entry)
}

func DeleteEntryHandler(w *rest.ResponseWriter, r *rest.Request){
     key, err := url.QueryUnescape(r.PathParam("Key"))
     defer r.Body.Close()
     if err != nil {
     	rest.Error(w, err.Error(), http.StatusBadRequest)
     }
	 _, ok := s.GetValue(key)
     if !ok {
     	rest.NotFound(w, r)
		return
     }
     s.DeleteEntry(key)
}

func IncrEntryHandler(w *rest.ResponseWriter, r *rest.Request){
     key, err := url.QueryUnescape(r.PathParam("Key"))
     defer r.Body.Close()
     w.Header().Set("Content-type", "application/json")
     if(err != nil){
     	    rest.Error(w, err.Error(), http.StatusBadRequest)
     }
     value, incErr := s.IncrEntry(key)
     if incErr != nil {
     	rest.Error(w, incErr.Error(), http.StatusPreconditionFailed)
     }
     w.WriteJson(&value)
}

func DecrEntryHandler(w *rest.ResponseWriter, r *rest.Request){
     key, err := url.QueryUnescape(r.PathParam("Key"))
     defer r.Body.Close()
     w.Header().Set("Content-type", "application/json")
     if(err != nil){
     	    rest.Error(w, err.Error(), http.StatusBadRequest)
     }
     value, decErr := s.DecrEntry(key)
     if decErr != nil {
     	rest.Error(w, decErr.Error(), http.StatusPreconditionFailed)
     }
     w.WriteJson(&value)
}