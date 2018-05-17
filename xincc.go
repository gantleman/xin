package main
import (
	"fmt"
	"strconv"

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

func (s *XinContract) regist(APIstub shim.ChaincodeStubInterface, args []string) {
	var key string;
	key += "xin.acc."
	key += args[0]

	_, err := APIstub.GetState(args[0])
	if err == nil {
		return
	}

	APIstub.PutState(key, []byte(args[1]))
}

func atoi (s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
}

//usr pwd address car, address fhash, string url, uint price, uint stamp
func (s *XinContract) addrepair(APIstub shim.ChaincodeStubInterface, args []string){
	var key string;
	key += "xin.vin."
	key += args[0]
	key += "."
	key += args[1]
	key += "."

	APIstub.PutState(key + "url", []byte(args[2]))
	APIstub.PutState(key + "price", []byte(args[3]))
	APIstub.PutState(key + "stamp", []byte(args[4]))
	APIstub.PutState(key + "usr", []byte(args[5]))
	
	var bi []byte
	var ikey,index string;
	ikey += "xin.index."
	ikey += args[0]

	bi, _ = APIstub.GetState(ikey)
	index = string(bi)
	index += args[4]
	index += ","
	index += args[5]
	index += ","
	index += args[1]
	index += ","
	index += args[3]
	index += ";"
	APIstub.PutState(ikey, []byte(index))
}

func (s *XinContract) buy(APIstub shim.ChaincodeStubInterface, args []string) {

	var key,bkey string;
	key += "xin.vin."
	key += args[0]
	key += "."
	key += args[1]
	key += "."

	bkey = key + "buy."
	bkey += args[3]

	APIstub.PutState(bkey, []byte(strconv.Itoa(1)))

	var ckey string;
	ckey += "xin.coin."

	sprice, _ := APIstub.GetState(key + "price")
	usr, _ := APIstub.GetState(key + "usr")

	ucoin, _ := APIstub.GetState(ckey + string(usr))
	scoin, _ := APIstub.GetState(ckey + args[3])

	coint := atoi(string(scoin));
	price := atoi(string(sprice));
	APIstub.PutState(key, []byte(strconv.Itoa(coint - price)))

	u_coint,_ := strconv.Atoi(string(ucoin));
	APIstub.PutState(key, []byte(strconv.Itoa(u_coint + price)))
}

func (s *XinContract) geturl(APIstub shim.ChaincodeStubInterface, args []string) (string){
	var ikey string;
	ikey += "xin.index."
	ikey += args[0]

	index, _ := APIstub.GetState(ikey)
	return string(index)
}

func main() {

	// Create a new Xin Contract
	err := shim.Start(new(XinContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}