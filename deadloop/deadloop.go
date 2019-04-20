package main

import (
	"errors"
	"fmt"
  "time"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type DoNothing struct {
}

func main() {
	err := shim.Start(new(DoNothing))
	if err != nil {
		fmt.Printf("Error starting donothing: %s", err)
	}
}

func (t *DoNothing) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

func (t *DoNothing) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "donothing" {
          temp := 1
          for {
            temp = temp + 1
            temp = temp - 1
            time.Sleep(1 * time.Nanosecond)
          }
	  return nil, nil
	}
	return nil, errors.New("Received unknown function invocation")
}

func (t *DoNothing) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, errors.New("Received unknown function query")
}
