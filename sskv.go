package main

import (
       "fmt"
)

// Single Server Key value store implementation

// Simple Implementation: v1 Use a global HashMap
var kvStore map[string]string

func Get(key string) string {
     return kvStore[key]
}


func Put(key string, value string) {
     kvStore[key] = value
}


func Delete(key string) {
     delete(kvStore, key)
}

func main(){
     kvStore = make(map[string]string)
     Put("key1", "important value1")
     fmt.Println(Get("key1"))
}