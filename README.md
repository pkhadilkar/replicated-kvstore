Single Server Key Value Store
====================================

Single Server key value server in Go language. Current interface provides flat key value space with keys and values as strings. The server provides REST interface to the key value store. Server uses [go-json-rest] (https://github.com/ant0ine/go-json-rest) by Antoine Imbert.

Install
-------
To use the server, create an executable using "go install". Running this executable will create a server that listens on port 9090. 
```
$ go get github.com/pkhadilkar/kvstore
$ go install github.com/pkhadilkar/kvstore	#this should create executable named kvstore (.exe on Windows)
$ /kvstore	#kvstore.exe on Windows
```
Requests can be submitted to server using JSON as shown in Examples. Note that this **only** starts the server. To test the server, please refer to [examples](#examples-test) section.

Types
-----
The server supports 

+ String
+ Integer (64 bit)

Operations *get* and *put* are supported on both types. Additionally Integer type supports

+ Increment
+ Decrement

operations which increment or decrement the integer value for the key by 1.

Examples / Test
--------
To automatically launch the sever and test sample functionality use
```
$ cd $GOPATH/src/github.com/pkhadilkar/kvstore
$ python test.py
```
test.py works only on Linux or Windows platform. To test on Mac OS X family, manually launch kvstore server and use tests in [API](#api).

Following commands can be used to used test individual features. Server should be launced manually as mentioned in the install before this.

```
$ curl -d '{"Key": "Pushkar", "Value": "+91-9975627439"}' http://localhost:9090/store
$ curl -d '{"Key": "Swapnil", "Value": "+91-9975946292"}' http://localhost:9090/store
$ curl -d '{"Key": "Random person", "Value": "+91-8679847479"}' http://localhost:9090/store
$ curl http://localhost:9090/store	  	   #displays all entries in key value store
$ curl http://localhost:9090/store/Pushkar
$ curl http://localhost:9090/store/Swapnil
$ curl http://localhost:9090/store/Random+person
$ curl -X DELETE http://localhost:9090/store/Random+person
$ curl http://localhost:9090/store
$ curl -d "{"Key": "counter", "Value": "15"}' http://localhost:9090/store
$ curl http://localhost:9090/store/counter    #original value
$ curl http://localhost:9090/store/incr/counter 	#increment the value
$ curl http://localhost:9090/store/counter		#display incremented value
```

API 
-----
A simple Go API for accessing the key value store is provided in *client* package . To test the API
```
$ cd $GOPATH/src/github.com/pkhadilkar/kvstore/client	#do this step if you are not already in client directory
$ go test
```
This assumes that server is already running.