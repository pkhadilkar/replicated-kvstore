package main

// Single Server Key value store implementation

type Entry struct{
     Key string
     Value string
}

// Simple Implementation: v1 Use a global HashMap
var kvStore map[string]string = make(map[string]string)


func GetValue(key string) (string, bool) {
     value, ok := kvStore[key]
     return value, ok
}


func GetAllEntries() *[]*Entry {
     entries := make([]*Entry, len(kvStore))
     i := 0
     for key, value := range kvStore {
     	 entry := Entry{key, value}
     	 entries[i] = &entry
	 i += 1
     }
     return &entries
}



func PutValue(e *Entry){
     kvStore[e.Key] = e.Value
}

func DeleteEntry(key string){
     delete(kvStore, key)
}


