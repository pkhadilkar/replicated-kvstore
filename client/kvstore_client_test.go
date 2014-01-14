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

    fmt.Println("Adding integer value (\"counter\", \"81\")")
    err = Put("counter", "81")
    if err != nil{
       t.Errorf("Error in put: ", err.Error())
    }
    fmt.Println(lineseparator)

    fmt.Println("Get value for key \"counter\"")
    value, err = Get("counter")
    if err != nil{
       t.Errorf("Error in get: ", err.Error())
    }
    fmt.Println(value)
    fmt.Println(lineseparator)

    fmt.Println("Decrement value for key \"counter\"")
    value, err = Decrement("counter")
    if err != nil{
       t.Errorf("Error in Decrement: ", err.Error())
    }
    fmt.Println("Decremented value for \"counter\" = ", value)
    fmt.Println(lineseparator)

    fmt.Println("Increment value for key \"counter\"")
    value, err = Increment("counter")
    if err != nil{
       t.Errorf("Error in increment: ", err.Error())
    }
    fmt.Println("Incremented value for \"counter\" = ", value)
    fmt.Println(lineseparator)
}