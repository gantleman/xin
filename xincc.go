package main
import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type XinContract struct {
}

func (s *XinContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *XinContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, _ := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "regist" {
		return shim.Success(nil)
	} else if function == "addrepair" {
		return shim.Success(nil)
	} else if function == "buy" {
		return  shim.Success(nil)
	} else if function == "getaddress" {
		return shim.Success(nil)
	} else if function == "geturl" {
		return shim.Success(nil)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *XinContract) regist(stub shim.ChaincodeStubInterface, args []string) {
}

func (s *XinContract) addrepair(stub shim.ChaincodeStubInterface, args []string){
}

func (s *XinContract) buy(stub shim.ChaincodeStubInterface, args []string) {
}

func (s *XinContract) geturl(stub shim.ChaincodeStubInterface, args []string) {
}

func main() {

	// Create a new Xin Contract
	err := shim.Start(new(XinContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}