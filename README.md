Single Server Key Value Store
====================================

Single Server key value server in Go language. Current interface provides flat key value space with keys and values as strings. The server provides REST interface to the key value store. Server uses [go-json-rest] (https://github.com/ant0ine/go-json-rest) by Antoine Imbert.

Install
-------
To use the server, create an executable using "go install". Running this executable will create a server that listens on port 9090. 
```
$ go get github.com/pkhadilkar/kvstore
$ go install github.com/pkhadilkar/kvstore	#this should create executable named kvstore (.exe on Windows)
$ ./kvstore	#kvstore.exe on Windows
```
Requests can be submitted to server using JSON as shown in Examples. Note that this *only* starts the server. To test the server, please refer to examples section.

Types
-----
The server supports two types string and integer (64 bits) as value types. Values of type integer have increment and decrement operation defined on them.

Examples / Test
--------
To automatically launch the sever and test sample functionality use
```
$ python test.py
```

Following commands can be used to used test individual features. Server should be launced manually as mentioned in the install before this.

```
$ curl -d '{"Key": "Pushkar", "Value": "+91-9975627439"}' http://127.0.0.1:9090/store
$ curl -d '{"Key": "Swapnil", "Value": "+91-9975946292"}' http://127.0.0.1:9090/store
$ curl -d '{"Key": "Random person", "Value": "+91-8679847479"}' http://127.0.0.1:9090/store
$ curl http://127.0.0.1:9090/store	  	   #displays all entries in key value store
$ curl http://127.0.0.1:9090/store/Pushkar
$ curl http://127.0.0.1:9090/store/Swapnil
$ curl http://127.0.0.1:9090/store/Random+person
$ curl -X DELETE http://127.0.0.1:9090/store/Random+person
$ curl http://127.0.0.1:9090/store
$ curl -d "{"Key": "counter", "Value": "15"}' http://127.0.0.1:9090/store
$ curl http://127.0.0.1:9090/store/counter    #original value
$ curl http://127.0.0.1:9090/store/incr/counter 	#increment the value
$ curl http://127.0.0.1:9090/store/counter		#display incremented value
```
