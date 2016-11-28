package main 

import (
	"fmt"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//==============================================================================================================================
//	 Structure Definitions
//==============================================================================================================================
//	Chaincode - A blank struct for use with Shim (A HyperLedger included go file used for get/put state
//				and other HyperLedger functions)
//==============================================================================================================================
type  SimpleChaincode struct {
}

//==============================================================================================================================
//	TransactionEvent - Defines the structure for a event object. JSON on right tells it what JSON fields to map to
//			  that element when reading a JSON object into the struct e.g. JSON datetime -> Struct datetime.
//==============================================================================================================================
//type TransactionEvent struct {
//	SenderName            string `json:"senderName"`
//	SenderCountry         string `json:"senderCountry"`
////	ReceiverName          string `json:"receiverName"`
////	ReceiverCountry       string `json:"receiverCountry"`
////	Amount		          string `json:"amount"`
////	DateTime	          string `json:"datetime"`
////	AccountNumber         string `json:"accountNumber"`
//}


//==============================================================================================================================
//	Init Function - Called when the user deploys the chaincode
//==============================================================================================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting 1")
    }

    err := stub.PutState("hello_world", []byte(args[0]))
    if err != nil {
        return nil, err
    }
    fmt.Println("invoke Init Method")
	
    return nil, nil
}

//==============================================================================================================================
//	Router Functions
//==============================================================================================================================
//	Invoke - Called on chaincode invoke. Takes a function name passed and calls that function.
//==============================================================================================================================
//func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
//	fmt.Println("invoke is running " + function)
//
//	// Handle different functions
//	if function == "init" {													//initialize the chaincode state, used as reset
//		return t.Init(stub, "init", args)
//	}else if function == "create_event" {
//        return t.create_event(stub, args)
//	}
//}

//=================================================================================================================================
//	 Create Function
//=================================================================================================================================
//	 Create Transaction Event - Creates the initial JSON for the Transaction Event and then saves it to the ledger.
//=================================================================================================================================
//func (t *SimpleChaincode) create_event(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
//	var tEvent TransactionEvent
//	
//	senderName     		:= "\"SenderName\":\""+args[0]+"\", "
//	senderCountry       := "\"SenderCountry\":\""+args[1]+"\", "
//
//	event_json := "{"+senderName+senderCountry+"}" 								// Concatenates the variables to create the total JSON object
//	err = json.Unmarshal([]byte(event_json), &tEvent)							// Convert the JSON defined above into a TransactionEvent object for go
//
//	if err != nil { return nil, errors.New("Invalid JSON object") }
//	
//	bytes, err := json.Marshal(tEvent)
//	if err != nil { fmt.Printf("SAVE_CHANGES: Error converting TransactionEvent record: %s", err); return false, errors.New("Error converting TransactionEvent record") }
//
//	err = stub.PutState("temp", bytes)
//	if err != nil { fmt.Printf("SAVE_CHANGES: Error storing TransactionEvent record: %s", err); return false, errors.New("Error storing TransactionEvent record") }
//
//	return true, nil 
//}


//=================================================================================================================================
//	 Ping Function
//=================================================================================================================================
//	 Pings the peer to keep the connection alive
//=================================================================================================================================
func (t *SimpleChaincode) ping(stub shim.ChaincodeStubInterface) ([]byte, error) {
	return []byte("Hello, world!"), nil
}


