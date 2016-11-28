package main 

import (
	"fmt"
	"errors"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var logger = shim.NewLogger("CLDChaincode")

//==============================================================================================================================
//	 Structure Definitions
//==============================================================================================================================
//	Chaincode - A blank struct for use with Shim (A HyperLedger included go file used for get/put state
//				and other HyperLedger functions)
//==============================================================================================================================
type  SimpleChaincode struct {
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

//==============================================================================================================================
//	TransactionEvent - Defines the structure for a event object. JSON on right tells it what JSON fields to map to
//			  that element when reading a JSON object into the struct e.g. JSON datetime -> Struct datetime.
//==============================================================================================================================
type TransactionEvent struct {
	SenderName            string `json:"senderName"`
	SenderCountry         string `json:"senderCountry"`
//	ReceiverName          string `json:"receiverName"`
//	ReceiverCountry       string `json:"receiverCountry"`
//	Amount		          string `json:"amount"`
//	DateTime	          string `json:"datetime"`
//	AccountNumber         string `json:"accountNumber"`
}


//==============================================================================================================================
//	Init Function - Called when the user deploys the chaincode
//==============================================================================================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke Init Method")
	  
    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting 1")
    }

    err := stub.PutState("hello_world", []byte(args[0]))
    if err != nil {
        return nil, err
    }
	
    return nil, nil
}

//==============================================================================================================================
//	Router Functions
//==============================================================================================================================
//	Invoke - Called on chaincode invoke. Takes a function name passed and calls that function.
//==============================================================================================================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {													//initialize the chaincode state, used as reset
		return t.Init(stub, "init", args)
	}else if function == "create_event" {
        return t.create_event(stub, args)
	}
	return nil, errors.New("Received unknown function invocation: " + function)
}

//=================================================================================================================================
//	Query - Called on chaincode query. Takes a function name passed and calls that function. Passes the
//  		initial arguments passed are passed on to the called function.
//=================================================================================================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query is running " + function)
	
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
	
	return nil, errors.New("Received unknown function invocation: " + function)
}

//=================================================================================================================================
//	 Create Function
//=================================================================================================================================
//	 Create Transaction Event - Creates the initial JSON for the Transaction Event and then saves it to the ledger.
//=================================================================================================================================
func (t *SimpleChaincode) create_event(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var tEvent TransactionEvent
	
	senderName     		:= "\"SenderName\":\""+args[0]+"\", "
	senderCountry       := "\"SenderCountry\":\""+args[1]+"\", "

	event_json := "{"+senderName+senderCountry+"}" 								// Concatenates the variables to create the total JSON object
	err := json.Unmarshal([]byte(event_json), &tEvent)							// Convert the JSON defined above into a TransactionEvent object for go

	if err != nil { 
		return nil, errors.New("Invalid JSON object") 
	}
	
	bytes, err := json.Marshal(tEvent)
	if err != nil { 
		fmt.Printf("SAVE_CHANGES: Error converting TransactionEvent record: %s", err); 
		return nil, errors.New("Error converting TransactionEvent record") 
	}

	err = stub.PutState("tr1", bytes)
	if err != nil { 
		fmt.Printf("SAVE_CHANGES: Error storing TransactionEvent record: %s", err); 
		return nil, errors.New("Error storing TransactionEvent record") 
	}

	return nil, nil 
}


//=================================================================================================================================
//	 Ping Function
//=================================================================================================================================
//	 Pings the peer to keep the connection alive
//=================================================================================================================================
func (t *SimpleChaincode) ping(stub shim.ChaincodeStubInterface) ([]byte, error) {
	return []byte("Hello, world!"), nil
}


