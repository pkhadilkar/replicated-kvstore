/*
Package server contains a single server implementation of key value store. 
Types:
	* string (custom types can be encoded as their string representation)
	* Integer (64 bit integer)
Operations supported:
	   * get
	   * put
	   * increment (only for Integer values)
	   * decrement (only for Integer values)
*/

package server

import (
       "strconv"
       "errors"
       "sync"
)

// Entry type represents a key value store entry.
type Entry struct{
     Key string
     Value string
}

type ValueWrapper struct{
     Value string
}


// base for supported int type
const base = 10		
// number of bits in the supported interger type
const integerBits = 64	

// synchronize accesses to map using channels
// capacity 1 to ensure single map accessor
// Ensure that server start process initializes
// channel by writing a value

type kvStore struct {
     store map[string]ValueWrapper
     mutex sync.Mutex
}

// GetValue gets a value from map for a given key. 
// It returns false if key is not present and true otherwise.
func (s *kvStore) GetValue(key string) (ValueWrapper, bool) {
     s.lock()
     defer s.unlock()
     value, ok := s.store[key]
     return value, ok
}

// GetAllEntries returns all key value pairs in map
// as a slice of Entry objects.
// Deprecated: This method will be removed in future versions
// as it is not scalable. A single call to this method
// would block every other call on keystore
func (s *kvStore)GetAllEntries() *[]*Entry {
     entries := make([]*Entry, len(s.store))
     i := 0
     s.lock()
     defer s.unlock()
     for key, value := range s.store {
     	 entry := Entry{key, value.Value}
     	 entries[i] = &entry
	 i += 1
     }

     return &entries
}


// PutValue stores a given entry object in map
func (s *kvStore)PutValue(e *Entry){
     value := ValueWrapper{e.Value}
     s.lock()
     defer s.unlock()
     s.store[e.Key] = value
}
// DeleteEntry deletes entry for a given key from kvstore
func (s *kvStore) DeleteEntry(key string){
     s.lock()
     defer s.unlock()
     delete(s.store, key)
}

// lock method allows only one thread to operate on map
// at a time. Other concurrent threads block.
func (s *kvStore) lock(){
	s.mutex.Lock()
}

// unlock method releases lock
func (s *kvStore) unlock(){
     s.mutex.Unlock()
}

// getInt method returns int64 value for given key if
// it is present
func (s *kvStore) getInt(key string) (int64, error){
     s.lock()
     defer s.unlock()
     value, ok := s.store[key]
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

// IncrEntry increments integer value for a given key by 1
// if the value is present and it is of type integer
func (s *kvStore) IncrEntry(key string) (ValueWrapper, error){
     s.lock()
     defer s.unlock()
     value, ok := s.store[key]
     if !ok  {
     	return ValueWrapper{}, errors.New("Value not found")
     }
     i, err := strconv.ParseInt(value.Value, base, integerBits)
     if err != nil {
     	return ValueWrapper{}, err
     }     
     i = i + 1
     valueWrapper := ValueWrapper{strconv.FormatInt(i, base)}
     s.store[key] = valueWrapper
     return valueWrapper, err
}

// DecrEntry decrements integer value for a given key by 1
// if the value is present and it is of type integer
func (s *kvStore) DecrEntry(key string) (ValueWrapper, error){
     s.lock()
     defer s.unlock()
     value, ok := s.store[key]
     if !ok {
     	return ValueWrapper{}, errors.New("Value not found")
     }
     i, err := strconv.ParseInt(value.Value, base, integerBits)
     if err != nil {
     	return ValueWrapper{}, err
     }     
     i = i - 1
     valueWrapper := ValueWrapper{strconv.FormatInt(i, base)}
     s.store[key] = valueWrapper
     return valueWrapper, err
}