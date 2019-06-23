package main

import (
	"fmt"
	//"strconv"
	"strings"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct {
}


type user struct {
	Name string `json:"name"`
	StdID string `json:"stdID"`
	Tel string `json:"tel"`
	Status bool `json:"status"`
}

type Wallet struct {
	WalletName string `json:"name"`
	Money string `json:"money"`
	Owner string `json:"owner"`
}


func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	// justString := strings.Join(args,"")

	// args = strings.Split(justString, "|")

	fmt.Println("abac Invoke")
	function, args := stub.GetFunctionAndParameters()
	if  function == "createUser" {
		// Create User
		return t.createUser(stub, args)
	} else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	} else if function == "createWallet" {
		// Create Wallet
		return t.createWallet(stub, args)
	} 

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger  (DB)
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

//create user 
func (t *SimpleChaincode) createUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	justString := strings.Join(args,"")

	args = strings.Split(justString, "|")
	
	//     0            1            2             3
	//   "name"       "stdID"       "tel"       "status"
	if len(args) != 3{
		return shim.Error("Incorrect number of 15151515151arguments. Expecting 4")
	}
	// Name := strings.ToLower(args[0])
		 stdID := args[1]
	// Tel := args[2]
	// Status := true

	UserKey := "stdID|"+stdID //stdID|25888656
	// UserBytes, err := stub.GetState(UserKey)    //"StdID"+StdID)
	// if UserBytes == nil{
	// 	return shim.Error("Failed tp get StdID :" +err.Error())
	// } 
	//Json
	 User := &user{
		Name : strings.ToLower(args[0]),
		StdID : args[1],
		Tel : args[2],
		Status : true,
	}
	// Maeshal แปลง Array ---> Json 	
	UserJSONasBytes, err := json.Marshal(User)
	if err != nil {
		return shim.Error(err.Error())
	}
	// เอาลงDB
	stub.PutState(UserKey, UserJSONasBytes)
	return shim.Success(nil)
}

func (t *SimpleChaincode) createWallet(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	justString := strings.Join(args,"")

	args = strings.Split(justString, "|")
	
	if len(args) != 3{
		return shim.Error("Incorrect number of Wallet arguments. Expecting 3")
	}
	 WalletName := strings.ToLower(args[0])
	// Money := args[1]
	// Owner := args[2]
	

	WalletKey := "WalletName|"+WalletName //WalletName|IsusBig
	// WalletBytes, err := stub.GetState(WalletKey) 
	// if WalletBytes == nil{
	// 	return shim.Error("Failed tp get WalletName :" +err.Error())
	// } 
	//Json

	 Wallet := &Wallet{
		WalletName : strings.ToLower(args[0]),
		Money : args[1],
		Owner : args[2],
	}
	// Maeshal แปลง Array ---> Json 	
	WalletJSONasBytes, err := json.Marshal(Wallet)
	if err != nil {
		return shim.Error(err.Error())
	}
	// เอาลงDB
	stub.PutState(WalletKey,WalletJSONasBytes)
	return shim.Success(nil)
}


func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}