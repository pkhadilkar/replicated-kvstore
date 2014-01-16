package server

import (
       "strconv"
       "errors"
)

// Single Server Key value store implementation
// Types
type Entry struct{
     Key string
     Value string
}

type ValueWrapper struct{
     Value string
}

// constants
const base = 10		// base for supported int type
const integerBits = 64	// number of bits in the supported interger type

// Simple Implementation: v1 Use a global HashMap
var kvStore map[string]ValueWrapper = make(map[string]ValueWrapper, 100000)

// synchronize accesses to map using channels
// capacity 1 to ensure single map accessor
// Ensure that server start process initializes
// channel by writing a value
var token chan bool = make(chan bool, 1)

func GetValue(key string) (ValueWrapper, bool) {
     lock()
     defer unlock()
     value, ok := kvStore[key]
     return value, ok
}


func GetAllEntries() *[]*Entry {
     entries := make([]*Entry, len(kvStore))
     i := 0
     lock()
     defer unlock()
     for key, value := range kvStore {
     	 entry := Entry{key, value.Value}
     	 entries[i] = &entry
	 i += 1
     }

     return &entries
}



func PutValue(e *Entry){
     value := ValueWrapper{e.Value}
     lock()
     defer unlock()
     kvStore[e.Key] = value
}

func DeleteEntry(key string){
     lock()
     defer unlock()
     delete(kvStore, key)
}

func lock(){
	<- token
}

func unlock(){
     token <- true
}

func getInt(key string) (int64, error){
     lock()
     defer unlock()
     value, ok := kvStore[key]
     if !ok{
     	return 0, errors.New("Key was not found in the map")
     }
     // parse the value to int
     i, err := strconv.ParseInt(value.Value, base, integerBits)
     if err != nil {
     	return 0, err
     }
     return i, err
}


func IncrEntry(key string) (ValueWrapper, error){
     lock()
     defer unlock()
     value, ok := kvStore[key]
     if !ok  {
     	return ValueWrapper{}, errors.New("Value not found")
     }
     i, err := strconv.ParseInt(value.Value, base, integerBits)
     if err != nil {
     	return ValueWrapper{}, err
     }     
     i = i + 1
     valueWrapper := ValueWrapper{strconv.FormatInt(i, base)}
     kvStore[key] = valueWrapper
     return valueWrapper, err
}

func DecrEntry(key string) (ValueWrapper, error){
     lock()
     defer unlock()
     value, ok := kvStore[key]
     if !ok {
     	return ValueWrapper{}, errors.New("Value not found")
     }
     i, err := strconv.ParseInt(value.Value, base, integerBits)
     if err != nil {
     	return ValueWrapper{}, err
     }     
     i = i - 1
     valueWrapper := ValueWrapper{strconv.FormatInt(i, base)}
     kvStore[key] = valueWrapper
     return valueWrapper, err
}