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
	TranID           	  string `json:"tranID"`
	SenderName            string `json:"senderName"`
	SenderCountry         string `json:"senderCountry"`
	ReceiverName          string `json:"receiverName"`
	ReceiverCountry       string `json:"receiverCountry"`
	Amount		          string `json:"amount"`
//	DateTime	          string `json:"datetime"`
//	AccountNumber         string `json:"accountNumber"`
}

//==============================================================================================================================
//	TranID Holder - Defines the structure that holds all the tranIDs for TransactionEvents that have been created.
//				    Used as an index when querying all transactions.
//==============================================================================================================================
type TRAN_Holder struct {
	TranIDs []string `json:"tranID"`
}


//==============================================================================================================================
//	Init Function - Called when the user deploys the chaincode
//==============================================================================================================================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke Init Method")
	
	var tranHld TRAN_Holder
	  
    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting 1")
    }

    bytes, err := json.Marshal(tranHld)

    if err != nil { 
    	return nil, errors.New("Error creating TRAN_Holder record") 
    }

	err = stub.PutState("tranIDs", bytes)
	
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
	}else if function == "ping" {
        return t.ping(stub)
    }
	
	return nil, errors.New("Received unknown function invocation: " + function)
}

//=================================================================================================================================
//	Query - Called on chaincode query. Takes a function name passed and calls that function. Passes the
//  		initial arguments passed are passed on to the called function.
//=================================================================================================================================
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("Query is running " + function)
	
	var err error
	
	if function == "get_event_details" {
		if len(args) != 1 { 
			fmt.Printf("Incorrect number of arguments passed"); 
			return nil, errors.New("QUERY: Incorrect number of arguments passed") 
		}
		
		v, err := t.retrieve_tranEvent(stub, args[0])
		if err != nil { 
			fmt.Printf("QUERY: Error retrieving tranEvent: %s", err); 
			return nil, errors.New("QUERY: Error retrieving tranEvent "+err.Error()) 
		}
		
		return t.get_vehicle_details(stub, v)
	}

	return nil, errors.New("Received unknown function invocation " + function)
}

//==============================================================================================================================
//	 retrieve_tranEvent - Gets the state of the data at tranID in the ledger then converts it from the stored
//					JSON into the TransactionEvent struct for use in the contract. Returns the TransactionEvent struct.
//					Returns empty TransactionEvent if it errors.
//==============================================================================================================================
func (t *SimpleChaincode) retrieve_tranEvent(stub shim.ChaincodeStubInterface, tranEventID string) (Vehicle, error) {

	var tranEvent TransactionEvent

	bytes, err := stub.GetState(tranEventID);

	if err != nil {	
		fmt.Printf("retrieve_tranEvent: Failed to retrieving TransactionEvent: %s", err); 
		return tranEvent, errors.New("retrieve_tranEvent: Error retrieving TransactionEvent with tranEventID = " + tranEventID) 
	}

	err = json.Unmarshal(bytes, &tranEvent);

    if err != nil {	
    	fmt.Printf("retrieve_tranEvent: Corrupt Event record "+string(bytes)+": %s", err); 
    	return tranEvent, errors.New("retrieve_tranEvent: Corrupt Event record"+string(bytes))	
    }

	return tranEvent, nil
}

//=================================================================================================================================
//	 Create Function
//=================================================================================================================================
//	 Create Transaction Event - Creates the initial JSON for the Transaction Event and then saves it to the ledger.
//=================================================================================================================================
func (t *SimpleChaincode) create_event(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var tEvent TransactionEvent
	
	tranID     			:= "\"TranID\":\""+args[0]+"\", "
	senderName     		:= "\"SenderName\":\""+args[1]+"\", "
	senderCountry       := "\"SenderCountry\":\""+args[2]+"\", "
	receiverName     	:= "\"ReceiverName\":\""+args[3]+"\", "
	receiverCountry     := "\"ReceiverCountry\":\""+args[4]+"\", "
	amount     			:= "\"Amount\":\""+args[5]+"\", "

    // Concatenates the variables to create the total JSON object
	event_json := "{"+tranID+senderName+senderCountry+receiverName+receiverCountry+amount+"}" 		
	// Convert the JSON defined above into a TransactionEvent object for go
	err := json.Unmarshal([]byte(event_json), &tEvent)										
	if err != nil { 
		return nil, errors.New("Invalid JSON object") 
	}

	bytes, err := json.Marshal(tEvent)
	if err != nil { 
		return nil, errors.New("Error converting transaction event") 
	}

	// Save new tran event record
	err = stub.PutState(tEvent.TranID, bytes)
	if err != nil { 
		fmt.Printf("create_event: Error storing transaction event: %s", err); 
		return nil, errors.New("Error storing transaction event") 
	}

	// Update tranIDs with newly created ID and store it in chain.
	bytes, err = stub.GetState("tranIDs")
	if err != nil { 
		return nil, errors.New("Unable to get tranIDs") 
	}

	var tranHld TRAN_Holder
	err = json.Unmarshal(bytes, &tranHld)
	if err != nil {	
		return nil, errors.New("Corrupt TRAN_Holder record") 
	}

	tranHld.TranIDs = append(tranHld.TranIDs, args[0])
	bytes, err = json.Marshal(tranHld)

	err = stub.PutState("tranIDs", bytes)
	if err != nil { 
		fmt.Printf("create_event: Error storing TranIDs: %s", err); 
		return nil, errors.New("Error storing TranIDs") 
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


