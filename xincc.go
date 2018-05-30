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
		return shim.Success([]byte(s.getaddress(APIstub, args)))
	} else if function == "geturl" {
		return shim.Success([]byte(s.geturl(APIstub, args)))
	} else if function == "addcoin" {
		s.addcoin(APIstub, args)
		return shim.Success(nil)
	} else if function == "getcoin" {
		return shim.Success([]byte(s.getcoin(APIstub, args)))
	}

	return shim.Error("Invalid Smart Contract function name.")
}

//0usr, 1pwd
func (s *XinContract) getcoin(APIstub shim.ChaincodeStubInterface, args []string) string{
	if s.check(APIstub, args[0], args[1]) == 0 {
		return ""
	}

	var ckey string;
	ckey += "xin.coin."

	scoin, _ := APIstub.GetState(ckey + args[0])
	return string(scoin);
}

//0usr, 1pwd, 2value
func (s *XinContract) addcoin(APIstub shim.ChaincodeStubInterface, args []string) {
	if s.check(APIstub, args[0], args[1]) == 0 {
		return
	}

	var ckey string;
	ckey += "xin.coin."

	scoin, _ := APIstub.GetState(ckey + args[0])
	coint := atoi(string(scoin));
	acoint := atoi(string(args[2]));
	APIstub.PutState(ckey+args[0], []byte(strconv.Itoa(coint + acoint)))
}

///param usr pwd
func (s *XinContract) regist(APIstub shim.ChaincodeStubInterface, args []string) {
	var key string;
	key += "xin.acc."
	key += args[0]

	pwd, err := APIstub.GetState(args[0])
	if pwd != nil {
		fmt.Printf("regist error:" + string(pwd) + ":" + err.Error())
		return
	}

	APIstub.PutState(key, []byte(args[1]))

	fmt.Printf("regist:" + string(key) + ":" + args[1] + "\n")
}

func (s *XinContract) check(APIstub shim.ChaincodeStubInterface, usr string, pwd string) int {
	var key string;
	key += "xin.acc."
	key += usr

	rpwd, _ := APIstub.GetState(key)
	if rpwd == nil {
		fmt.Printf("usr error:" + key + "\n")
		return 0
	}

	if string(rpwd) != pwd {
		fmt.Printf("pwd error:" + string(rpwd) + "\n")
		return 0
	} else {
		return 1
	}
}

func atoi (s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
}

//0usr, 1pwd, 2car, 3fhash, 4url, 5price, 6stamp
func (s *XinContract) addrepair(APIstub shim.ChaincodeStubInterface, args []string){

	if s.check(APIstub, args[0], args[1]) == 0 {
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

	if s.check(APIstub, args[0], args[1]) == 0 {
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
	APIstub.PutState(ckey+args[0], []byte(strconv.Itoa(coint - price)))

	uscoin, _ := APIstub.GetState(ckey + string(usr))
	ucoint, _ := strconv.Atoi(string(uscoin));
	APIstub.PutState(ckey + string(usr), []byte(strconv.Itoa(ucoint + price)))
}

//0usr, 1pwd, 2car, 3fhash
func (s *XinContract) geturl(APIstub shim.ChaincodeStubInterface, args []string) (string){
	if s.check(APIstub, args[0], args[1]) == 0 {
		return ""
	}

	var key string;
	key += "xin.vin."
	key += args[2]
	key += "."
	key += args[3]
	key += "."

	index, _ := APIstub.GetState(key + "url")
	return string(index)
}

//0usr, 1pwd, 2car
func (s *XinContract) getaddress(APIstub shim.ChaincodeStubInterface, args []string) (string){
	if s.check(APIstub, args[0], args[1]) == 0 {
		return ""
	}

	var ikey string;
	ikey += "xin.index."
	ikey += args[2]

	index, _ := APIstub.GetState(ikey)
	return string(index)
}

//0usr, 1pwd, 2car, 3fhash
func (s *XinContract) isbuy(APIstub shim.ChaincodeStubInterface) sc.Response {
	function, args := APIstub.GetFunctionAndParameters()
	if s.check(APIstub, args[0], args[1]) == 0 {
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

	is, _ := APIstub.GetState(bkey)
	return shim.Success(is)
}

func main() {

	// Create a new Xin Contract
	err := shim.Start(new(XinContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}