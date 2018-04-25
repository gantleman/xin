pragma solidity ^0.4.0;

contract xintoken {
    struct rate
    {
        uint  v_rate;//1~5
        string comment;
    }
    
    struct repair 
    {
        string url;
        uint price;
        uint stamp;
        uint balance;
        address seller;
        
        mapping(address => uint) buy;
        mapping(address => rate) fallow;
    }
    
    mapping(address => mapping(address => repair)) public info;
    mapping(address => uint) public coin;
    mapping(address => string) public index;
    
    
    mapping(address => mapping(address=>bool)) public mkeys;
    address public admin;
    
    function xintoken(address padmin) public {
         admin = padmin;
    }
    
    function addcoin(address buyer, uint value) public payable{
        if(msg.sender != admin)
            return;
        coin[buyer] += value;
    }
    
   function getcount(address car) public view returns(uint r){
       r = bytes(index[car]).length;
   }
    
    function toString(address x) internal returns (string) {
        bytes memory b = new bytes(20);
        for (uint i = 0; i < 20; i++)
            b[i] = byte(uint8(uint(x) / (2**(8*(19 - i)))));
        return string(b);
    }
    
    function bytes32ToString (bytes32 data) internal returns (string) {
        bytes memory bytesString = new bytes(32);
        for (uint j=0; j<32; j++) {
            byte char = byte(bytes32(uint(data) * 2 ** (8 * j)));
            if (char != 0) {
                bytesString[j] = char;
            }
        }
        return string(bytesString);
    }
    
    function strConcat(string _a, string _b) internal returns (string r){
        bytes memory _ba = bytes(_a);
        bytes memory _bb = bytes(_b);
        string memory abcde = new string(_ba.length + _bb.length);
        bytes memory babcde = bytes(abcde);
        uint k = 0;
        for (uint i = 0; i < _ba.length; i++) babcde[k++] = _ba[i];
        for (i = 0; i < _bb.length; i++) babcde[k++] = _bb[i];
        r = string(babcde);
    }
    
    
    ///添加维保记录
    function add_repair(address car, address fhash, string url, uint price, uint stamp) public payable returns (int r) {
        //assert(bytes(info[car][fhash].url).length != 0);
        //assert(bytes(url).length == 0);
        r = 1;
        info[car][fhash].url = url;
        info[car][fhash].price = price;
        info[car][fhash].seller = msg.sender;
        info[car][fhash].stamp = stamp;
        r++;
       if(mkeys[car][fhash] == false){
            string memory rets;
           
           // rets = strConcat(rets, bytes32ToString(bytes32(stamp)));
            
            rets = strConcat(rets, ",");
            //rets = strConcat(rets, toString(msg.sender));
            rets = strConcat(rets, ",");
            //rets = strConcat(rets, toString(fhash));
            rets = strConcat(rets, ",");
            //rets = strConcat(rets, bytes32ToString(bytes32(price)));
            rets = strConcat(rets, ";");
            index[car] = strConcat(index[car], rets);
            mkeys[car][fhash] = true;
            r++;
        }
    }

    ///添加评分
    function add_rate(address car, address fhash, uint irate, string comment) public payable {
        assert(bytes(info[car][fhash].url).length == 0);
        
        info[car][fhash].fallow[msg.sender].v_rate = irate;
        info[car][fhash].fallow[msg.sender].comment = comment;
    }

    ///购买维保记录
    function buy(address car, address fhash) public payable {
        assert(bytes(info[car][fhash].url).length == 0);
        assert(info[car][fhash].price > coin[msg.sender]);

            
        info[car][fhash].buy[msg.sender] += info[car][fhash].price;
        info[car][fhash].balance += info[car][fhash].price * 9 /10;
        
        coin[admin] += info[car][fhash].price/10;
        coin[msg.sender] = coin[msg.sender] - info[car][fhash].price;
    }
    
    ///获取地址
    function getaddress(address car) public view returns (string rets){
        return index[car];
    }
    
    ///4s店检查是否购买过维保记录
    function isbuy(address car, address fhash, address buyer) public view returns (uint){
        if(info[car][fhash].seller != msg.sender)
            return;
        if(info[car][fhash].buy[buyer] / info[car][fhash].price >= 1)
            return 1;
        else
            return 0; 
    }
    
    ///获取url
    function geturl(address car, address fhash) public view returns (string){
        if(info[car][fhash].buy[msg.sender] / info[car][fhash].price < 1)
            return '';
            
        return info[car][fhash].url; 
    }
    
    //提现接口
    function collect() public payable{
        assert(msg.sender != admin);
        admin.transfer(address(this).balance);
    }
    
    function selfcollect(address car, address seller, uint value) public payable returns (uint){
        assert(msg.sender != admin);
        info[car][seller].balance -= value;
        return info[car][seller].balance;
    }
    
    function getcoin(address seller) public view returns (uint){
        if(msg.sender != admin)
            return;
            
        return coin[seller];
    }
    
    function getcoin() public view returns (uint){
        return coin[msg.sender];
    }
}