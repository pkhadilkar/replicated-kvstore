package main

// Single Server Key value store implementation

type Entry struct{
     Key string
     Value string
}

type ValueWrapper struct{
     Value string
}

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


