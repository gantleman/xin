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
        uint balance;
        
        mapping(address => uint) buy;
        mapping(address => rate) fallow;
    }
    
    struct sell4s
    {
        uint T;
        address[] mkeys;
        mapping(address => repair) detailed;
    }
    
    mapping(address => sell4s) public info;
    
    address admin;
    
    function xintoken(address padmin) public {
         admin = padmin;
    }
    
 ///添加维保记录
    function add_repair(address car, string url, uint price) public payable {
        info[car].detailed[msg.sender].url = url;
        info[car].detailed[msg.sender].price = price;
        
        for(uint i =0; i< info[car].mkeys.length; i++)
        {
            if(info[car].mkeys[i] == msg.sender)
                return;
        }
        info[car].mkeys[info[car].mkeys.length] = msg.sender;
    }

    ///添加评分
    function add_rate(address car, address sell, uint irate, string comment) public payable {
        info[car].detailed[sell].fallow[msg.sender].v_rate = irate;
        info[car].detailed[sell].fallow[msg.sender].comment = comment;
    }

    ///购买维保记录
    function buy(address car, address sell) public payable returns (string url){
        if(bytes(info[car].detailed[sell].url).length == 0)
            return '';

        info[car].detailed[sell].buy[msg.sender] += msg.value;
        info[car].detailed[sell].balance += msg.value * 9 /10;
        
        return info[car].detailed[sell].url;
    }
    
     function toString(address x) internal returns (string) {
        bytes memory b = new bytes(20);
        for (uint i = 0; i < 20; i++)
            b[i] = byte(uint8(uint(x) / (2**(8*(19 - i)))));
        return string(b);
    }
    
    function strConcat(string _a, string _b) internal returns (string){
        bytes memory _ba = bytes(_a);
        bytes memory _bb = bytes(_b);
        string memory abcde = new string(_ba.length + _bb.length);
        bytes memory babcde = bytes(abcde);
        uint k = 0;
        for (uint i = 0; i < _ba.length; i++) babcde[k++] = _ba[i];
        for (i = 0; i < _bb.length; i++) babcde[k++] = _bb[i];
        return string(babcde);
    }

    ///获取4s车商地址
    function getaddress(address car) public view returns (string){
        string memory rets;
        for(uint i =0; i< info[car].mkeys.length; i++)
        {
            rets = strConcat(rets, toString(info[car].mkeys[i]));
            rets = strConcat(rets, ";");
        }
    }
    
    ///检查是否购买过维保记录
    function isbuy(address car, address buyer) public view returns (uint){
        return info[car].detailed[msg.sender].buy[buyer] / info[car].detailed[msg.sender].price >= 1 ? 1:0;
    }
    
    //提现接口
    function collect() public payable{
        if(msg.sender != admin)
        return;
        
        admin.transfer(address(this).balance);
    }
    
    function selfcollect(address car) public payable{
        if(info[car].detailed[msg.sender].balance == 0)
            return;
        if(address(this).balance == 0)
            return;
            
        if(address(this).balance >= info[car].detailed[msg.sender].balance)
            admin.transfer(info[car].detailed[msg.sender].balance);
        else
            admin.transfer(address(this).balance);
    }
}