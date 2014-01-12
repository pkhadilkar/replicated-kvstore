Single Server Key Value Store (SSKV)
===================================

Single Server key value server in Go language. Current interface provides flat key value space with keys and values as strings. The server provides REST interface to the key value store. Server uses [go-json-rest] (https://github.com/ant0ine/go-json-rest) by Antoine Imbert.

Install
-------
To use the server, create an executable using "go install". Running this executable will create a server that listens on port 9090. 

Examples
--------
Following commands can be used to used test the functionality

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
```
