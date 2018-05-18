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
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "regist" {
		s.regist(APIstub, args)
		return shim.Success(nil)
	} else if function == "addrepair" {
		s.addrepair(APIstub, args)
		return shim.Success(nil)
	} else if function == "buy" {
		s.buy(APIstub, args)
		return  shim.Success(nil)
	} else if function == "getaddress" {
		s.getaddress(APIstub, args)
		return shim.Success(nil)
	} else if function == "geturl" {
		s.geturl(APIstub, args)
		return shim.Success(nil)
	}

	return shim.Error("Invalid Smart Contract function name.")
}
///param usr pwd
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

func (s *XinContract) check(APIstub shim.ChaincodeStubInterface, usr string, pwd string) int {
	var key string;
	key += "xin.acc."
	key += args[0]

	rpwd, err := APIstub.GetState([]byte(usr))
	if err == nil {
		return 0
	}

	if rpwd != pwd {
		return 0
	}else
	{
		return 1
	}
}

func atoi (s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
}

//0usr, 1pwd, 2car, 3fhash, 4url, 5price, 6stamp
func (s *XinContract) addrepair(APIstub shim.ChaincodeStubInterface, args []string){

	if check(args[0], args[1]) == 0
	{
		return
	}

	var key string;
	key += "xin.vin."
	key += args[2]
	key += "."
	key += args[3]
	key += "."

	APIstub.PutState(key + "url", []byte(args[4]))
	APIstub.PutState(key + "price", []byte(args[5]))
	APIstub.PutState(key + "stamp", []byte(args[6]))
	APIstub.PutState(key + "usr", []byte(args[0]))
	
	var bi []byte
	var ikey,index string;
	ikey += "xin.index."
	ikey += args[2]

	bi, _ = APIstub.GetState(ikey)
	index = string(bi)
	index += args[6]
	index += ","
	index += args[0]
	index += ","
	index += args[3]
	index += ","
	index += args[5]
	index += ";"
	APIstub.PutState(ikey, []byte(index))
}

//0usr, 1pwd, 2car, 3fhash
func (s *XinContract) buy(APIstub shim.ChaincodeStubInterface, args []string) {

	if check(args[0], args[1]) == 0
	{
		return
	}

	var key,bkey string;
	key += "xin.vin."
	key += args[2]
	key += "."
	key += args[3]
	key += "."

	bkey = key + "buy."
	bkey += args[0]

	APIstub.PutState(bkey, []byte(strconv.Itoa(1)))

	sprice, _ := APIstub.GetState(key + "price")
	usr, _ := APIstub.GetState(key + "usr")

	var ckey string;
	ckey += "xin.coin."

	scoin, _ := APIstub.GetState(ckey + args[0])
	coint := atoi(string(scoin));
	price := atoi(string(sprice));
	APIstub.PutState(key, []byte(strconv.Itoa(coint - price)))

	uscoin, _ := APIstub.GetState(ckey + string(usr))
	ucoint,_ := strconv.Atoi(string(uscoin));
	APIstub.PutState(key, []byte(strconv.Itoa(u_coint + price)))
}

//0usr, 1pwd, 2car,
func (s *XinContract) geturl(APIstub shim.ChaincodeStubInterface, args []string) (string){
	if check(args[0], args[1]) == 0
	{
		return
	}

	var ikey string;
	ikey += "xin.index."
	ikey += args[2]

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