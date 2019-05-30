package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	var Mike, John,Amy,Bob,Jack string    // Entities
	var Mikeval, Johnval,Amyval,Bobval,Jackval int // Asset holdings
	var err error
	
	Mike = "Mike"
	Mikeval, err = strconv.Atoi("100")
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	John ="John"
	Johnval, err = strconv.Atoi("200")
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	Amy ="Amy"
	Amyval, err = strconv.Atoi("300")
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	Bob = "Bob"
	Bobval, err = strconv.Atoi("400")
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	Jack = "Jack"
	Jackval, err = strconv.Atoi("500")
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}


	fmt.Printf("Mikeval = %d, Johnval = %d, Amyval = %d, Bobval = %d, Jackval = %d\n", Mikeval, Johnval,Amyval,Bobval,Jackval)

	// Write the state to the ledger
	err = stub.PutState(Mike, []byte(strconv.Itoa(Mikeval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(John, []byte(strconv.Itoa(Johnval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(Amy, []byte(strconv.Itoa(Amyval)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(Bob, []byte(strconv.Itoa(Bobval)))
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(Jack, []byte(strconv.Itoa(Jackval)))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		// Make payment of X units from A to B
		return t.invoke(stub, args)
	}  else if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	} else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	}
	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}


// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var X, Y int          // Transaction value
	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	A = args[0]
	B = args[1]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Avalbytes == nil {
		return shim.Error("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Bvalbytes == nil {
		return shim.Error("Entity not found")
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}
	Y, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}

	Aval = Aval - X
	Bval = Bval + Y
	fmt.Printf("%s = %d, %s = %d\n", A ,Aval,B, Bval)

	
	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
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

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}