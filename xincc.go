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
	var key string;
	key += "xin.acc."
	key += args[0]

	pwd, err := APIstub.GetState(args[0])
	if err == nil {
		return
	}

	APIstub.PutState(key, args[1])
}
//usr pwd address car, address fhash, string url, uint price, uint stamp
//分别放入不同的变量内并拼接字符串
//这里要对账户密码做校验
func (s *XinContract) addrepair(stub shim.ChaincodeStubInterface, args []string){
	var key string;
	key += "xin.vin."
	key += args[0]
	key += "."
	key += args[1]
	key += "."

	APIstub.PutState(key + "url", strconv.Atoi(args[2]))
	APIstub.PutState(key + "price", strconv.Atoi(args[3]))
	APIstub.PutState(key + "stamp", strconv.Atoi(args[4]))
	APIstub.PutState(key + "usr", strconv.Atoi(args[5]))

	var ikey string;
	ikey += "xin.index."
	ikey += args[0]

	index, err := APIstub.GetState(ikey)
	index += strconv.Atoi(args[4])
	index += ","
	index += strconv.Atoi(args[5])
	index += ","
	index += strconv.Atoi(args[1])
	index += ","
	index += strconv.Atoi(args[3])
	index += ";"
}

func (s *XinContract) buy(stub shim.ChaincodeStubInterface, args []string) {

	var key string;
	key += "xin.vin."
	key += args[0]
	key += "."
	key += args[1]
	key += "."

	bkey = key + "buy."
	bkey += args[3]

	APIstub.PutState(bkey, strconv.Itoa(1))

	var ckey string;
	ckey += "xin.coin."

	sprice, err := APIstub.GetState(key + "price")
	usr, err := APIstub.GetState(key + "usr")

	ucoin, err := APIstub.GetState(ckey + usr)
	scoin, err := APIstub.GetState(ckey + args[3])

	coint := strconv.Atoi(scoin);
	price := strconv.Atoi(sprice);
	APIstub.PutState(key, strconv.Itoa(coin - price))

	u_coint := strconv.Atoi(ucoin);
	APIstub.PutState(key, strconv.Itoa(u_coint + price))
}

func (s *XinContract) geturl(stub shim.ChaincodeStubInterface, args []string) {
 APIstub.GetState(args[0])
}

func main() {

	// Create a new Xin Contract
	err := shim.Start(new(XinContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}