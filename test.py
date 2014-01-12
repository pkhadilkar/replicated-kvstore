#!/bin/python
# Test script for single server key value store
import subprocess
import platform
import distutils.spawn
import json
import urllib, urllib2
import httplib

line_separator = "==============================="
serverBase = "127.0.0.1:9090"
protocol = 'http://'
appRoot = '/store'

def get_json(key, value) :
    x = {"Key": key, "Value" : value}
    return json.dumps(x)

def post_request_print(key, value):
    """ Post a POST request and pring response """
    print 'Put ("'+ key +', "'+ value + '")'
    req = urllib2.Request(protocol + serverBase + appRoot)
    req.add_header('Content-type', 'application/json')
    data = get_json(key, value)
    response = urllib2.urlopen(req, data)
    print response.read()



print "Launching server ...."

executable_path = ""
executable_name = ""

if platform.system() == "Windows":
    executable_name = "kvstore.exe"
elif platform.system() == "Linux":
    executable_name = "kvstore"
else:
    print "Unsupported platform:", platform.system()
    exit(1)


executable_path = distutils.spawn.find_executable(executable_name)
if executable_path == None :
    print "Please ensure that", executable_name, "is on system path variable"
    exit(1)

proc = subprocess.Popen([executable_path], stdout=open('server_output.log', 'w'))

print "Running tests ...."
print
print

# use urllib.quote_plus function to encode the query
post_request_print("Pushkar", "95860959999")


print line_separator
post_request_print("Swapnil", "7569847694")


print line_separator
post_request_print("Random person", "57896798")

print line_separator

print "Display all key value pairs"
response = urllib.urlopen(protocol + serverBase + appRoot)
print response.read()

print line_separator

print 'Display value stored for ("Random person")'
response = urllib.urlopen(protocol + serverBase + appRoot + "/" + urllib.quote_plus("Random person"))
print response.read()

print line_separator

print 'Delete data associated with key "Random person"'
conn = httplib.HTTPConnection(serverBase)
req = conn.request('DELETE', appRoot + '/' + urllib.quote_plus("Random person"))
response = conn.getresponse()
print response.read()

print line_separator


print "Display all key value pairs"
response = urllib.urlopen(protocol + serverBase + appRoot)
print response.read()

print line_separator

print 'Create an entry with integer value ("counter", "15")'
post_request_print("counter", "15")

print line_separator

print 'Increment value associated with "counter"'
response = urllib.urlopen(protocol + serverBase + appRoot + "/incr/counter")
print response.read()

print line_separator
print 'Decrement value associated with "counter"'
response = urllib.urlopen(protocol + serverBase + appRoot + "/decr/counter")
print response.read()

print line_separator
print
print
print "All tests completed successfully"

proc.kill()
