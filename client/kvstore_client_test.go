package client

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

const lineseparator = "====================================="

func TestKvstoreClient(t *testing.T) {
	time.Sleep(5 * time.Second) // time for Raft cluster to choose the leader
	fmt.Println("Put (\"Random person\", \"new phone\")")
	fmt.Println(Put("Random person", "new phone."))
	fmt.Println(lineseparator)

	fmt.Println("Get value for key \"Random peson\"")
	value, err := Get("Random person")
	if err != nil {
		t.Errorf("Error in get: ", err.Error())
	}
	fmt.Println(value)
	fmt.Println(lineseparator)

	fmt.Println("Put (\"bnm\", \"intelligble value\")")
	fmt.Println(Put("bnm", "intelligble value"))

	fmt.Println(lineseparator)

	fmt.Println("Get value for key \"bnm\"")
	value, err = Get("bnm")
	if err != nil {
		t.Errorf("Error in get: ", err.Error())
	}
	fmt.Println(value)
	fmt.Println(lineseparator)

	fmt.Println("Delete value for key \"bnm\"")
	err = Delete("bnm")
	if err != nil {
		t.Errorf("Error in delete: ", err.Error())
	}

	fmt.Println(lineseparator)

	fmt.Println("Get value for key \"bnm\"")
	value, err = Get("bnm")
	fmt.Println(value)
	fmt.Println(lineseparator)

	fmt.Println("Adding integer value (\"counter\", \"81\")")
	err = Put("counter", "81")
	if err != nil {
		t.Errorf("Error in put: ", err.Error())
	}
	fmt.Println(lineseparator)

	fmt.Println("Get value for key \"counter\"")
	value, err = Get("counter")
	if err != nil {
		t.Errorf("Error in get: ", err.Error())
	}
	fmt.Println(value)
	fmt.Println(lineseparator)

	fmt.Println("Decrement value for key \"counter\"")
	value, err = Decrement("counter")
	if err != nil {
		t.Errorf("Error in Decrement: ", err.Error())
	}
	fmt.Println("Decremented value for \"counter\" = ", value)
	fmt.Println(lineseparator)

	fmt.Println("Increment value for key \"counter\"")
	value, err = Increment("counter")
	if err != nil {
		t.Errorf("Error in increment: ", err.Error())
	}
	fmt.Println("Incremented value for \"counter\" = ", value)
	fmt.Println(lineseparator)
}

const clients = 10

var w sync.WaitGroup

// TestKvStoreStress runs a stress test to test performance
// and consistency of kvstore. It launches 1000 clients
// each of whom do 10 gets and 10 puts
func TestKvStoreStress(t *testing.T) {
	base := "client_id_"
	// one error message per client
	errChan := make(chan string, clients)
	w.Add(clients)
	for i := 0; i < clients; i += 1 {
		go launchClient(base+strconv.Itoa(i), errChan)
	}
	var buffer bytes.Buffer
	// check for error messages from one of the clients
	select {
	case msg, ok := <-errChan:
		if ok {
			for {
				buffer.WriteString(msg)
				msg, ok = <-errChan
				if !ok {
					break
				}
			}
			// not very smart. Writes all error messages to console
			t.Errorf(buffer.String(), errors.New(buffer.String()))
		}
	default:
		//no error detected
	}

	w.Wait()
}

func launchClient(id string, errChan chan<- string) {
	defer w.Done()
	for i := 0; i < 10; i += 1 {
		key := id + "_" + strconv.Itoa(i)
		value := key + "_value"

		if err := Put(key, value); err != nil {
			errChan <- "Error in client " + id + "\n" + err.Error()
			return
		}

		if storedVal, err := Get(key); err != nil {
			errChan <- "Error in client " + id + "\n" + err.Error()
			return
		} else if value != storedVal {
			errChan <- "Error in client " + id + "\nStore value does not match required value"
		}
	}
}
