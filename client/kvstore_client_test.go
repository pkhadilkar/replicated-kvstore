package client

import (
       "testing"
       "fmt"
)

const lineseparator = "====================================="

func TestKvstoreClient(t *testing.T) {

    fmt.Println("Put (\"Random person\", \"new phone\")")
    fmt.Println(Put("Random person", "new phone."))
    fmt.Println(lineseparator)
    
    fmt.Println("Get value for key \"Random peson\"")
    value, err := Get("Random person")
    if err != nil{
       t.Errorf("Error in get: ", err.Error())
    }
    fmt.Println(value)
    fmt.Println(lineseparator)    

    fmt.Println("Put (\"bnm\", \"intelligble value\")")
    fmt.Println(Put("bnm", "intelligble value"))

    fmt.Println(lineseparator)
    
    fmt.Println("Get value for key \"bnm\"")
    value, err = Get("bnm")
    if err != nil{
       t.Errorf("Error in get: ", err.Error())
    }
    fmt.Println(value)
    fmt.Println(lineseparator)
    
    fmt.Println("Delete value for key \"bnm\"")
    err = Delete("bnm")
    if err != nil{
       t.Errorf("Error in delete: ", err.Error())
    }
    
    fmt.Println(lineseparator)
    
    fmt.Println("Get value for key \"bnm\"")
    value, err = Get("bnm")
    fmt.Println(value)
    fmt.Println(lineseparator)

}