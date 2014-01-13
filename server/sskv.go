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


func GetValue(key string) (ValueWrapper, bool) {
     value, ok := kvStore[key]
     return value, ok
}


func GetAllEntries() *[]*Entry {
     entries := make([]*Entry, len(kvStore))
     i := 0
     for key, value := range kvStore {
     	 entry := Entry{key, value.Value}
     	 entries[i] = &entry
	 i += 1
     }
     return &entries
}



func PutValue(e *Entry){
     value := ValueWrapper{e.Value}
     kvStore[e.Key] = value
}

func DeleteEntry(key string){
     delete(kvStore, key)
}

func getInt(key string) (int64, error){
     value, ok := GetValue(key)
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
     i, err := getInt(key)
     if err != nil {
     	return ValueWrapper{}, err
     }
     i = i + 1
     valueWrapper := ValueWrapper{strconv.FormatInt(i, base)}
     kvStore[key] = valueWrapper
     return valueWrapper, err
}

func DecrEntry(key string) (ValueWrapper, error){
     i, err := getInt(key)
     if err != nil {
     	return ValueWrapper{}, err
     }
     i = i - 1
     valueWrapper := ValueWrapper{strconv.FormatInt(i, base)}
     kvStore[key] = valueWrapper
     return valueWrapper, err
}